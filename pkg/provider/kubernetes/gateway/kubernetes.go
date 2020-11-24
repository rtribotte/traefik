package gateway

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/mitchellh/hashstructure"
	ptypes "github.com/traefik/paerser/types"
	"github.com/traefik/traefik/v2/pkg/config/dynamic"
	"github.com/traefik/traefik/v2/pkg/job"
	"github.com/traefik/traefik/v2/pkg/log"
	"github.com/traefik/traefik/v2/pkg/provider"
	"github.com/traefik/traefik/v2/pkg/safe"
	"github.com/traefik/traefik/v2/pkg/tls"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/service-apis/apis/v1alpha1"
)

const (
	providerName = "kubernetesgateway"
)

// Provider holds configurations of the provider.
type Provider struct {
	Endpoint          string          `description:"Kubernetes server endpoint (required for external cluster client)." json:"endpoint,omitempty" toml:"endpoint,omitempty" yaml:"endpoint,omitempty"`
	Token             string          `description:"Kubernetes bearer token (not needed for in-cluster client)." json:"token,omitempty" toml:"token,omitempty" yaml:"token,omitempty"`
	CertAuthFilePath  string          `description:"Kubernetes certificate authority file path (not needed for in-cluster client)." json:"certAuthFilePath,omitempty" toml:"certAuthFilePath,omitempty" yaml:"certAuthFilePath,omitempty"`
	Namespaces        []string        `description:"Kubernetes namespaces." json:"namespaces,omitempty" toml:"namespaces,omitempty" yaml:"namespaces,omitempty" export:"true"`
	LabelSelector     string          `description:"Kubernetes label selector to select specific GatewayClasses." json:"labelSelector,omitempty" toml:"labelSelector,omitempty" yaml:"labelSelector,omitempty" export:"true"`
	ThrottleDuration  ptypes.Duration `description:"Kubernetes refresh throttle duration" json:"throttleDuration,omitempty" toml:"throttleDuration,omitempty" yaml:"throttleDuration,omitempty" export:"true"`
	lastConfiguration safe.Safe
	EntryPoints       map[string]Entrypoint `json:"-" toml:"-" yaml:"-" label:"-" file:"-"`
}

// Entrypoint defines the available entry points.
type Entrypoint struct {
	Address        string
	HasHTTPTLSConf bool
}

func (p *Provider) newK8sClient(ctx context.Context) (*clientWrapper, error) {
	// Label selector validation
	_, err := labels.Parse(p.LabelSelector)
	if err != nil {
		return nil, fmt.Errorf("invalid label selector: %q", p.LabelSelector)
	}
	log.FromContext(ctx).Infof("label selector is: %q", p.LabelSelector)

	withEndpoint := ""
	if p.Endpoint != "" {
		withEndpoint = fmt.Sprintf(" with endpoint %s", p.Endpoint)
	}

	var client *clientWrapper
	switch {
	case os.Getenv("KUBERNETES_SERVICE_HOST") != "" && os.Getenv("KUBERNETES_SERVICE_PORT") != "":
		log.FromContext(ctx).Infof("Creating in-cluster Provider client%s", withEndpoint)
		client, err = newInClusterClient(p.Endpoint)
	case os.Getenv("KUBECONFIG") != "":
		log.FromContext(ctx).Infof("Creating cluster-external Provider client from KUBECONFIG %s", os.Getenv("KUBECONFIG"))
		client, err = newExternalClusterClientFromFile(os.Getenv("KUBECONFIG"))
	default:
		log.FromContext(ctx).Infof("Creating cluster-external Provider client%s", withEndpoint)
		client, err = newExternalClusterClient(p.Endpoint, p.Token, p.CertAuthFilePath)
	}

	if err != nil {
		return nil, err
	}
	client.labelSelector = p.LabelSelector

	return client, nil
}

// Init the provider.
func (p *Provider) Init() error {
	return nil
}

// Provide allows the k8s provider to provide configurations to traefik
// using the given configuration channel.
func (p *Provider) Provide(configurationChan chan<- dynamic.Message, pool *safe.Pool) error {
	ctxLog := log.With(context.Background(), log.Str(log.ProviderName, providerName))
	logger := log.FromContext(ctxLog)

	k8sClient, err := p.newK8sClient(ctxLog)
	if err != nil {
		return err
	}

	pool.GoCtx(func(ctxPool context.Context) {
		operation := func() error {
			eventsChan, err := k8sClient.WatchAll(p.Namespaces, ctxPool.Done())
			if err != nil {
				logger.Errorf("Error watching kubernetes events: %v", err)
				timer := time.NewTimer(1 * time.Second)
				select {
				case <-timer.C:
					return err
				case <-ctxPool.Done():
					return nil
				}
			}

			throttleDuration := time.Duration(p.ThrottleDuration)
			throttledChan := throttleEvents(ctxLog, throttleDuration, pool, eventsChan)
			if throttledChan != nil {
				eventsChan = throttledChan
			}

			for {
				select {
				case <-ctxPool.Done():
					return nil
				case event := <-eventsChan:
					// Note that event is the *first* event that came in during this throttling interval -- if we're hitting our throttle, we may have dropped events.
					// This is fine, because we don't treat different event types differently.
					// But if we do in the future, we'll need to track more information about the dropped events.
					conf := p.loadConfigurationFromGateway(ctxLog, k8sClient)

					confHash, err := hashstructure.Hash(conf, nil)
					switch {
					case err != nil:
						logger.Error("Unable to hash the configuration")
					case p.lastConfiguration.Get() == confHash:
						logger.Debugf("Skipping Kubernetes event kind %T", event)
					default:
						p.lastConfiguration.Set(confHash)
						configurationChan <- dynamic.Message{
							ProviderName:  providerName,
							Configuration: conf,
						}
					}

					// If we're throttling,
					// we sleep here for the throttle duration to enforce that we don't refresh faster than our throttle.
					// time.Sleep returns immediately if p.ThrottleDuration is 0 (no throttle).
					time.Sleep(throttleDuration)
				}
			}
		}

		notify := func(err error, time time.Duration) {
			logger.Errorf("Provider connection error: %v; retrying in %s", err, time)
		}
		err := backoff.RetryNotify(safe.OperationWithRecover(operation), backoff.WithContext(job.NewBackOff(backoff.NewExponentialBackOff()), ctxPool), notify)
		if err != nil {
			logger.Errorf("Cannot connect to Provider: %v", err)
		}
	})

	return nil
}

// nolint:funlen // TODO: refactor
// TODO Handle errors and update resources statuses (gatewayClass, gateway).
func (p *Provider) loadConfigurationFromGateway(ctx context.Context, client Client) *dynamic.Configuration {
	conf := &dynamic.Configuration{
		UDP: &dynamic.UDPConfiguration{
			Routers:  map[string]*dynamic.UDPRouter{},
			Services: map[string]*dynamic.UDPService{},
		},
		TCP: &dynamic.TCPConfiguration{
			Routers:  map[string]*dynamic.TCPRouter{},
			Services: map[string]*dynamic.TCPService{},
		},
		HTTP: &dynamic.HTTPConfiguration{
			Routers:     map[string]*dynamic.Router{},
			Middlewares: map[string]*dynamic.Middleware{},
			Services:    map[string]*dynamic.Service{},
		},
		TLS: &dynamic.TLSConfiguration{},
	}

	logger := log.FromContext(log.With(ctx))

	allowedGateway := map[string]struct{}{}

	gatewayClasses, err := client.GetGatewayClasses()
	if err != nil {
		logger.Errorf("Cannot found GatewayClasses : %v", err)
		return conf
	}

	for _, gatewayClass := range gatewayClasses {
		if gatewayClass.Spec.Controller == "traefik.io/gateway-controller" {
			allowedGateway[gatewayClass.Name] = struct{}{}

			err := client.UpdateGatewayClassStatus(gatewayClass, metav1.Condition{
				Type:               "InvalidParameters",
				Status:             metav1.ConditionFalse,
				LastTransitionTime: metav1.Now(),
				Reason:             "Handled",
				Message:            "Handled by Traefik controller",
			})
			if err != nil {
				logger.Errorf("Failed to update InvalidParameters condition: %v", err)
			}
		}
	}

	// TODO check if we can only use the default filtering mechanism
	gateways, err := client.GetGateways()
	if err != nil {
		logger.Errorf("Cannot found Gateways : %v", err)
	}

	for _, gateway := range gateways {
		ctxLog := log.With(ctx, log.Str("Gateway", gateway.Namespace+"/"+gateway.Name))
		logger := log.FromContext(ctxLog)

		// As Status.Addresses is not implemented yet, initialize an empty array to follow the API expectations
		gatewayStatus := v1alpha1.GatewayStatus{
			Addresses: []v1alpha1.GatewayAddress{},
		}

		if _, ok := allowedGateway[gateway.Spec.GatewayClassName]; !ok {
			gatewayStatus.Conditions = append(gatewayStatus.Conditions, metav1.Condition{
				Type:               string(v1alpha1.GatewayConditionScheduled),
				Status:             metav1.ConditionFalse,
				LastTransitionTime: metav1.Now(),
				Reason:             string(v1alpha1.GatewayReasonNoSuchGatewayClass),
				Message:            "Cannot found matching GatewayClass",
			})

			err := client.UpdateGatewayStatus(gateway, gatewayStatus)
			if err != nil {
				logger.Debugf("An error occurred while updating gateway status: %v", err)
			}

			continue
		}

		tlsConfigs := make(map[string]*tls.CertAndStores)
		routers := make(map[string]*dynamic.Router)
		services := make(map[string]*dynamic.Service)
		listenerInError := false

		// GatewayReasonListenersNotValid is used when one or more
		// Listeners have an invalid or unsupported configuration
		// and cannot be configured on the Gateway.
		listenerStatuses := make([]v1alpha1.ListenerStatus, len(gateway.Spec.Listeners))
		for i, listener := range gateway.Spec.Listeners {
			listenerStatuses[i] = v1alpha1.ListenerStatus{
				Port:       listener.Port,
				Conditions: []metav1.Condition{},
			}

			ep, err := p.entryPointName(listener.Port, listener.Protocol)
			if err != nil {
				msg := fmt.Sprintf("Cannot found entryPoint for Gateway : %v", err)
				// update "Detached" status with "PortUnavailable" reason
				listenerStatuses[i].Conditions = append(listenerStatuses[i].Conditions, metav1.Condition{
					Type:               string(v1alpha1.ListenerConditionDetached),
					Status:             metav1.ConditionTrue,
					LastTransitionTime: metav1.Now(),
					Reason:             string(v1alpha1.ListenerReasonPortUnavailable),
					Message:            msg,
				})

				logger.Errorf(msg)
				listenerInError = true
				continue
			}

			if listener.Protocol == v1alpha1.HTTPSProtocolType || listener.Protocol == v1alpha1.TLSProtocolType {
				if listener.TLS == nil {
					msg := fmt.Sprintf("No TLS configuration for Gateway Listener port %d and protocol %q", listener.Port, listener.Protocol)
					// update "Detached" status with "UnsupportedProtocol" reason
					listenerStatuses[i].Conditions = append(listenerStatuses[i].Conditions, metav1.Condition{
						Type:               string(v1alpha1.ListenerConditionDetached),
						Status:             metav1.ConditionTrue,
						LastTransitionTime: metav1.Now(),
						Reason:             string(v1alpha1.ListenerReasonUnsupportedProtocol),
						Message:            msg,
					})

					logger.Errorf(msg)
					listenerInError = true
					continue
				}

				if listener.TLS.CertificateRef.Kind != "Secret" || listener.TLS.CertificateRef.Group != "core" {
					msg := fmt.Sprintf("Unsupported TLS CertificateRef group/kind : %v/%v", listener.TLS.CertificateRef.Group, listener.TLS.CertificateRef.Kind)
					// update "ResolvedRefs" status true with "InvalidCertificateRef" reason
					listenerStatuses[i].Conditions = append(listenerStatuses[i].Conditions, metav1.Condition{
						Type:               string(v1alpha1.ListenerConditionResolvedRefs),
						Status:             metav1.ConditionTrue,
						LastTransitionTime: metav1.Now(),
						Reason:             string(v1alpha1.ListenerReasonInvalidCertificateRef),
						Message:            msg,
					})

					logger.Errorf(msg)
					listenerInError = true
					continue
				}

				configKey := gateway.Namespace + "/" + listener.TLS.CertificateRef.Name
				if _, tlsExists := tlsConfigs[configKey]; !tlsExists {
					tlsConf, err := getTLS(client, listener.TLS.CertificateRef.Name, gateway.Namespace)
					if err != nil {
						msg := fmt.Sprintf("Error while retrieving certificate: %v", err)
						// update "ResolvedRefs" status true with "InvalidCertificateRef" reason
						listenerStatuses[i].Conditions = append(listenerStatuses[i].Conditions, metav1.Condition{
							Type:               string(v1alpha1.ListenerConditionResolvedRefs),
							Status:             metav1.ConditionFalse,
							LastTransitionTime: metav1.Now(),
							Reason:             string(v1alpha1.ListenerReasonInvalidCertificateRef),
							Message:            msg,
						})

						logger.Errorf(msg)
						listenerInError = true
						continue
					}

					tlsConfigs[configKey] = tlsConf
				}
			}

			// Supported Protocol
			if listener.Protocol != v1alpha1.HTTPProtocolType && listener.Protocol != v1alpha1.HTTPSProtocolType {
				msg := fmt.Sprintf("Unsupported listener protocol %q", listener.Protocol)
				// update "Detached" status true with "UnsupportedProtocol" reason
				listenerStatuses[i].Conditions = append(listenerStatuses[i].Conditions, metav1.Condition{
					Type:               string(v1alpha1.ListenerConditionDetached),
					Status:             metav1.ConditionTrue,
					LastTransitionTime: metav1.Now(),
					Reason:             string(v1alpha1.ListenerReasonUnsupportedProtocol),
					Message:            msg,
				})

				logger.Errorf(msg)
				listenerInError = true
				continue
			}

			// Supported Route types
			if listener.Routes.Kind != "HTTPRoute" {
				msg := fmt.Sprintf("Unsupported Route Kind %q", listener.Routes.Kind)
				// update "ResolvedRefs" status true with "InvalidRoutesRef" reason
				listenerStatuses[i].Conditions = append(listenerStatuses[i].Conditions, metav1.Condition{
					Type:               string(v1alpha1.ListenerConditionResolvedRefs),
					Status:             metav1.ConditionFalse,
					LastTransitionTime: metav1.Now(),
					Reason:             string(v1alpha1.ListenerReasonInvalidRoutesRef),
					Message:            msg,
				})

				logger.Errorf(msg)
				listenerInError = true
				continue
			}

			// TODO: support RouteNamespaces
			httpRoutes, err := client.GetHTTPRoutes(gateway.Namespace, labels.SelectorFromSet(listener.Routes.Selector.MatchLabels))
			if err != nil {
				msg := fmt.Sprintf("Cannot fetch HTTPRoutes for namespace %q and matchLabels %v", gateway.Namespace, listener.Routes.Selector.MatchLabels)
				// update "ResolvedRefs" status true with "InvalidRoutesRef" reason
				listenerStatuses[i].Conditions = append(listenerStatuses[i].Conditions, metav1.Condition{
					Type:               string(v1alpha1.ListenerConditionResolvedRefs),
					Status:             metav1.ConditionFalse,
					LastTransitionTime: metav1.Now(),
					Reason:             string(v1alpha1.ListenerReasonInvalidRoutesRef),
					Message:            msg,
				})

				logger.Errorf(msg)
				listenerInError = true
				continue
			}

			for _, httpRoute := range httpRoutes {
				// Should never happen
				if httpRoute == nil {
					continue
				}

				hostRule := hostRule(httpRoute.Spec)

				for _, routeRule := range httpRoute.Spec.Rules {
					router := dynamic.Router{
						Rule:        extractRule(routeRule, hostRule),
						EntryPoints: []string{ep},
					}

					if listener.TLS != nil {
						// TODO support let's encrypt
						router.TLS = &dynamic.RouterTLSConfig{}
					}

					routerKey, err := makeRouterKey(router.Rule, makeID(httpRoute.Namespace, httpRoute.Name))
					if err != nil {
						msg := fmt.Sprintf("Skipping HTTPRoute %s: cannot make router's key with rule %s: %v", httpRoute.Name, router.Rule, err)
						// update "ResolvedRefs" status true with "DroppedRoutes" reason
						listenerStatuses[i].Conditions = append(listenerStatuses[i].Conditions, metav1.Condition{
							Type:               string(v1alpha1.ListenerConditionResolvedRefs),
							Status:             metav1.ConditionFalse,
							LastTransitionTime: metav1.Now(),
							Reason:             string(v1alpha1.ListenerReasonDegradedRoutes),
							Message:            msg,
						})

						logger.Errorf(msg)
						// TODO update the RouteStatus condition / deduplicate conditions on listener
						continue
					}

					if routeRule.ForwardTo != nil {
						wrrService, subServices, err := loadServices(client, gateway.Namespace, routeRule.ForwardTo)
						if err != nil {
							msg := fmt.Sprintf("Cannot load service from HTTPRoute %s/%s : %v", gateway.Namespace, httpRoute.Name, err)
							// update "ResolvedRefs" status true with "DroppedRoutes" reason
							listenerStatuses[i].Conditions = append(listenerStatuses[i].Conditions, metav1.Condition{
								Type:               string(v1alpha1.ListenerConditionResolvedRefs),
								Status:             metav1.ConditionFalse,
								LastTransitionTime: metav1.Now(),
								Reason:             string(v1alpha1.ListenerReasonDegradedRoutes),
								Message:            msg,
							})

							logger.Errorf(msg)
							// TODO update the RouteStatus condition / deduplicate conditions on listener
							continue
						}

						for svcName, svc := range subServices {
							services[svcName] = svc
						}

						serviceName := provider.Normalize(routerKey + "-wrr")
						services[serviceName] = wrrService

						router.Service = serviceName
					}

					if router.Service != "" {
						routerKey = provider.Normalize(routerKey)
						routers[routerKey] = &router
					}
				}
			}

			if len(listenerStatuses[i].Conditions) == 0 {
				// GatewayConditionReady "Ready", GatewayConditionReason "ListenerReady"
				listenerStatuses[i].Conditions = append(listenerStatuses[i].Conditions, metav1.Condition{
					Type:               string(v1alpha1.ListenerConditionReady),
					Status:             metav1.ConditionTrue,
					LastTransitionTime: metav1.Now(),
					Reason:             "ListenerReady",
					Message:            "No error found",
				})
			} else {
				listenerInError = true
			}
		}

		gatewayStatus.Listeners = listenerStatuses

		if listenerInError {
			// GatewayConditionReady "Ready", GatewayConditionReason "ListenersNotValid"
			gatewayStatus.Conditions = append(gatewayStatus.Conditions, metav1.Condition{
				Type:               string(v1alpha1.GatewayConditionReady),
				Status:             metav1.ConditionFalse,
				LastTransitionTime: metav1.Now(),
				Reason:             string(v1alpha1.GatewayReasonListenersNotValid),
				Message:            "All Listeners must be valid",
			})

			err := client.UpdateGatewayStatus(gateway, gatewayStatus)
			if err != nil {
				logger.Debugf("An error occurred while updating gateway status: %v", err)
			}

			continue
		}

		// update "Scheduled" status with "ResourcesAvailable" reason
		gatewayStatus.Conditions = append(gatewayStatus.Conditions, metav1.Condition{
			Type:               string(v1alpha1.GatewayConditionScheduled),
			Status:             metav1.ConditionTrue,
			LastTransitionTime: metav1.Now(),
			Reason:             "ResourcesAvailable",
			Message:            "Found matching HTTPRoutes",
		})

		// update "Ready" status with "ListenersValid" reason
		gatewayStatus.Conditions = append(gatewayStatus.Conditions, metav1.Condition{
			Type:               string(v1alpha1.GatewayConditionReady),
			Status:             metav1.ConditionTrue,
			LastTransitionTime: metav1.Now(),
			Reason:             "ListenersValid",
			Message:            "All listeners are valid",
		})

		err := client.UpdateGatewayStatus(gateway, gatewayStatus)
		if err != nil {
			logger.Debugf("An error occurred while updating gateway status: %v", err)
		}

		if len(routers) > 0 {
			for k, router := range routers {
				// TODO manage duplicate key
				conf.HTTP.Routers[k] = router
			}
		}

		if len(services) > 0 {
			for k, service := range services {
				// TODO manage duplicate key
				conf.HTTP.Services[k] = service
			}
		}

		if len(tlsConfigs) > 0 {
			if conf.TLS == nil {
				conf.TLS = &dynamic.TLSConfiguration{}
			}
			conf.TLS.Certificates = append(conf.TLS.Certificates, getTLSConfig(tlsConfigs)...)
		}
	}

	return conf
}

func hostRule(httpRouteSpec v1alpha1.HTTPRouteSpec) string {
	if len(httpRouteSpec.Hostnames) > 0 {
		hostRule := ""
		for i, hostname := range httpRouteSpec.Hostnames {
			if i > 0 && len(hostname) > 0 {
				hostRule += "`, `"
			}
			hostRule += string(hostname)
		}

		if hostRule != "" {
			return "Host(`" + hostRule + "`)"
		}
	}

	return ""
}

func extractRule(routeRule v1alpha1.HTTPRouteRule, hostRule string) string {
	var rule string
	var matchesRules []string

	for _, match := range routeRule.Matches {
		var matchRules []string
		// TODO handle other path types
		if match.Path.Type == v1alpha1.PathMatchExact {
			matchRules = append(matchRules, "Path(`"+match.Path.Value+"`)")
		}

		if match.Path.Type == v1alpha1.PathMatchPrefix {
			matchRules = append(matchRules, "PathPrefix(`"+match.Path.Value+"`)")
		}

		// TODO handle other headers types
		if match.Headers != nil && match.Headers.Type == v1alpha1.HeaderMatchExact {
			var headerRules []string

			for headerName, headerValue := range match.Headers.Values {
				headerRules = append(headerRules, "Headers(`"+headerName+"`,`"+headerValue+"`)")
			}
			// to have a consistent order
			sort.Strings(headerRules)
			matchRules = append(matchRules, headerRules...)
		}

		matchesRules = append(matchesRules, strings.Join(matchRules, " && "))
	}

	// If no matches are specified, the default is a prefix
	// path match on "/", which has the effect of matching every
	// HTTP request.
	if len(routeRule.Matches) == 0 {
		matchesRules = append(matchesRules, "PathPrefix(`/`)")
	}

	if hostRule != "" {
		if len(matchesRules) == 0 {
			return hostRule
		}
		rule += hostRule + " && "
	}

	if len(matchesRules) == 1 {
		return rule + matchesRules[0]
	}

	if len(rule) == 0 {
		return strings.Join(matchesRules, " || ")
	}

	return rule + "(" + strings.Join(matchesRules, " || ") + ")"
}

func (p *Provider) entryPointName(port v1alpha1.PortNumber, protocol v1alpha1.ProtocolType) (string, error) {
	entryPointName := ""
	portStr := strconv.FormatInt(int64(port), 10)
	for name, entryPoint := range p.EntryPoints {
		if strings.HasSuffix(entryPoint.Address, ":"+portStr) {
			// if the protocol is HTTP the entryPoint must have no TLS conf
			if protocol != "HTTP" || !entryPoint.HasHTTPTLSConf {
				entryPointName = name
				break
			}
		}
	}

	if entryPointName == "" {
		return "", fmt.Errorf("no matching entryPoint for port %d and protocol %q", port, protocol)
	}

	return entryPointName, nil
}

func makeRouterKey(rule, name string) (string, error) {
	h := sha256.New()
	if _, err := h.Write([]byte(rule)); err != nil {
		return "", err
	}

	key := fmt.Sprintf("%s-%.10x", name, h.Sum(nil))

	return key, nil
}

func makeID(namespace, name string) string {
	if namespace == "" {
		return name
	}

	return namespace + "-" + name
}

func getTLS(k8sClient Client, secretName, namespace string) (*tls.CertAndStores, error) {
	secret, exists, err := k8sClient.GetSecret(namespace, secretName)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch secret %s/%s: %w", namespace, secretName, err)
	}
	if !exists {
		return nil, fmt.Errorf("secret %s/%s does not exist", namespace, secretName)
	}

	cert, key, err := getCertificateBlocks(secret, namespace, secretName)
	if err != nil {
		return nil, err
	}

	return &tls.CertAndStores{
		Certificate: tls.Certificate{
			CertFile: tls.FileOrContent(cert),
			KeyFile:  tls.FileOrContent(key),
		},
	}, nil
}

func getTLSConfig(tlsConfigs map[string]*tls.CertAndStores) []*tls.CertAndStores {
	var secretNames []string
	for secretName := range tlsConfigs {
		secretNames = append(secretNames, secretName)
	}
	sort.Strings(secretNames)

	var configs []*tls.CertAndStores
	for _, secretName := range secretNames {
		configs = append(configs, tlsConfigs[secretName])
	}

	return configs
}

func getCertificateBlocks(secret *corev1.Secret, namespace, secretName string) (string, string, error) {
	var missingEntries []string

	tlsCrtData, tlsCrtExists := secret.Data["tls.crt"]
	if !tlsCrtExists {
		missingEntries = append(missingEntries, "tls.crt")
	}

	tlsKeyData, tlsKeyExists := secret.Data["tls.key"]
	if !tlsKeyExists {
		missingEntries = append(missingEntries, "tls.key")
	}

	if len(missingEntries) > 0 {
		return "", "", fmt.Errorf("secret %s/%s is missing the following TLS data entries: %s",
			namespace, secretName, strings.Join(missingEntries, ", "))
	}

	cert := string(tlsCrtData)
	if cert == "" {
		missingEntries = append(missingEntries, "tls.crt")
	}

	key := string(tlsKeyData)
	if key == "" {
		missingEntries = append(missingEntries, "tls.key")
	}

	if len(missingEntries) > 0 {
		return "", "", fmt.Errorf("secret %s/%s contains the following empty TLS data entries: %s",
			namespace, secretName, strings.Join(missingEntries, ", "))
	}

	return cert, key, nil
}

// loadServices is generating a WRR service, even when there is only one target.
func loadServices(client Client, namespace string, targets []v1alpha1.HTTPRouteForwardTo) (*dynamic.Service, map[string]*dynamic.Service, error) {
	services := map[string]*dynamic.Service{}

	wrrSvc := &dynamic.Service{
		Weighted: &dynamic.WeightedRoundRobin{
			Services: []dynamic.WRRService{},
		},
	}

	for _, forwardTo := range targets {
		svc := dynamic.Service{
			LoadBalancer: &dynamic.ServersLoadBalancer{
				PassHostHeader: func(v bool) *bool { return &v }(true),
			},
		}

		// TODO Handle BackendRefs

		if forwardTo.ServiceName != nil {
			service, exists, err := client.GetService(namespace, *forwardTo.ServiceName)
			if err != nil {
				return nil, nil, err
			}

			if !exists {
				return nil, nil, errors.New("service not found")
			}

			if len(service.Spec.Ports) > 1 && forwardTo.Port == 0 {
				// If the port is unspecified and the backend is a Service
				// object consisting of multiple port definitions, the route
				// must be dropped from the Gateway. The controller should
				// raise the "ResolvedRefs" condition on the Gateway with the
				// "DroppedRoutes" reason.  The gateway status for this route
				// should be updated with a condition that describes the error
				// more specifically.
				log.WithoutContext().Errorf("A multiple ports Kubernetes Service cannot be used if unspecified forwardTo.Port")
				continue
			}

			var portName string
			var portSpec corev1.ServicePort
			var match bool

			for _, p := range service.Spec.Ports {
				if forwardTo.Port == 0 || p.TargetPort.IntVal == int32(forwardTo.Port) {
					portName = p.Name
					portSpec = p
					match = true
					break
				}
			}

			if !match {
				return nil, nil, errors.New("service port not found")
			}

			endpoints, endpointsExists, endpointsErr := client.GetEndpoints(namespace, *forwardTo.ServiceName)
			if endpointsErr != nil {
				return nil, nil, endpointsErr
			}

			if !endpointsExists {
				return nil, nil, errors.New("endpoints not found")
			}

			if len(endpoints.Subsets) == 0 {
				return nil, nil, errors.New("subset not found")
			}

			var port int32
			for _, subset := range endpoints.Subsets {
				for _, p := range subset.Ports {
					if portName == p.Name {
						port = p.Port
						break
					}
				}

				if port == 0 {
					return nil, nil, errors.New("cannot define a port")
				}

				protocol := getProtocol(portSpec, portName)

				for _, addr := range subset.Addresses {
					svc.LoadBalancer.Servers = append(svc.LoadBalancer.Servers, dynamic.Server{
						URL: fmt.Sprintf("%s://%s:%d", protocol, addr.IP, port),
					})
				}
			}

			serviceName := provider.Normalize(makeID(service.Namespace, service.Name) + "-" + strconv.FormatInt(int64(port), 10))
			services[serviceName] = &svc

			weight := int(forwardTo.Weight)
			wrrSvc.Weighted.Services = append(wrrSvc.Weighted.Services, dynamic.WRRService{Name: serviceName, Weight: &weight})
		}
	}

	if len(services) == 0 {
		return nil, nil, errors.New("no service has been created")
	}

	return wrrSvc, services, nil
}

func getProtocol(portSpec corev1.ServicePort, portName string) string {
	protocol := "http"
	if portSpec.Port == 443 || strings.HasPrefix(portName, "https") {
		protocol = "https"
	}

	return protocol
}

func throttleEvents(ctx context.Context, throttleDuration time.Duration, pool *safe.Pool, eventsChan <-chan interface{}) chan interface{} {
	if throttleDuration == 0 {
		return nil
	}
	// Create a buffered channel to hold the pending event (if we're delaying processing the event due to throttling)
	eventsChanBuffered := make(chan interface{}, 1)

	// Run a goroutine that reads events from eventChan and does a non-blocking write to pendingEvent.
	// This guarantees that writing to eventChan will never block,
	// and that pendingEvent will have something in it if there's been an event since we read from that channel.
	pool.GoCtx(func(ctxPool context.Context) {
		for {
			select {
			case <-ctxPool.Done():
				return
			case nextEvent := <-eventsChan:
				select {
				case eventsChanBuffered <- nextEvent:
				default:
					// We already have an event in eventsChanBuffered, so we'll do a refresh as soon as our throttle allows us to.
					// It's fine to drop the event and keep whatever's in the buffer -- we don't do different things for different events
					log.FromContext(ctx).Debugf("Dropping event kind %T due to throttling", nextEvent)
				}
			}
		}
	})

	return eventsChanBuffered
}
