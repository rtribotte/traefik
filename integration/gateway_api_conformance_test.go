package integration

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/go-check/check"
	"github.com/rs/zerolog/log"
	"github.com/traefik/traefik/v3/integration/try"
	checker "github.com/vdemeester/shakers"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
	"sigs.k8s.io/gateway-api/conformance/tests"
	"sigs.k8s.io/gateway-api/conformance/utils/flags"
	"sigs.k8s.io/gateway-api/conformance/utils/suite"
)

// GateSuite tests suite.
type GateSuite struct{ BaseSuite }

func (s *GateSuite) SetUpSuite(c *check.C) {
	s.createComposeProject(c, "k8s")
	s.composeUp(c)

	abs, err := filepath.Abs("./fixtures/k8s/config.skip/kubeconfig.yaml")
	c.Assert(err, checker.IsNil)

	err = try.Do(60*time.Second, func() error {
		_, err := os.Stat(abs)
		return err
	})
	c.Assert(err, checker.IsNil)

	err = os.Setenv("KUBECONFIG", abs)
	c.Assert(err, checker.IsNil)
}

func (s *GateSuite) TearDownSuite(c *check.C) {
	s.composeDown(c)

	generatedFiles := []string{
		"./fixtures/k8s/config.skip/kubeconfig.yaml",
		"./fixtures/k8s/config.skip/k3s.log",
		"./fixtures/k8s/coredns.yaml",
		"./fixtures/k8s/rolebindings.yaml",
		"./fixtures/k8s/traefik.yaml",
		"./fixtures/k8s/ccm.yaml",
	}

	for _, filename := range generatedFiles {
		if err := os.Remove(filename); err != nil {
			log.Warn().Err(err).Send()
		}
	}
}

func (g *GateSuite) TestIngressConfiguration(c *check.C) {
	cmd, display := g.traefikCmd(withConfigFile("fixtures/k8s_gateway_conformance.toml"))
	defer display(c)

	err := cmd.Start()
	c.Assert(err, checker.IsNil)
	defer g.killCmd(cmd)

	cfg, err := config.GetConfig()
	if err != nil {
		c.Fatalf("Error loading Kubernetes config: %v", err)
	}
	client, err := client.New(cfg, client.Options{})
	if err != nil {
		c.Fatalf("Error initializing Kubernetes client: %v", err)
	}
	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		c.Fatalf("Error initializing Kubernetes REST client: %v", err)
	}

	gatev1.AddToScheme(client.Scheme())

	//supportedFeatures := suite.ParseSupportedFeatures(*flag.String("supported-features", "Gateway,HTTPRoute", "Supported features included in conformance tests suites"))
	exemptFeatures := suite.ParseSupportedFeatures(*flags.ExemptFeatures)
	skipTests := suite.ParseSkipTests(*flags.SkipTests)
	namespaceLabels := suite.ParseKeyValuePairs(*flags.NamespaceLabels)
	namespaceAnnotations := suite.ParseKeyValuePairs(*flags.NamespaceAnnotations)

	c.Logf("Running conformance tests with %s GatewayClass\n cleanup: %t\n debug: %t\n enable all features: %t \n supported features: [%v]\n exempt features: [%v]",
		*flags.GatewayClassName, *flags.CleanupBaseResources, *flags.ShowDebug, *flags.EnableAllSupportedFeatures, *flags.SupportedFeatures, *flags.ExemptFeatures)

	cSuite := suite.New(suite.Options{
		Client:     client,
		RestConfig: cfg,
		// This clientset is needed in addition to the client only because
		// controller-runtime client doesn't support non CRUD sub-resources yet (https://github.com/kubernetes-sigs/controller-runtime/issues/452).
		Clientset:            clientset,
		GatewayClassName:     *flags.GatewayClassName,
		Debug:                *flags.ShowDebug,
		CleanupBaseResources: *flags.CleanupBaseResources,
		SupportedFeatures: sets.New[suite.SupportedFeature]().
			Insert(suite.GatewayExtendedFeatures.UnsortedList()...).
			Insert(suite.ReferenceGrantCoreFeatures.UnsortedList()...).
			Insert(suite.HTTPRouteCoreFeatures.UnsortedList()...).
			Insert(suite.HTTPRouteExtendedFeatures.UnsortedList()...).
			Insert(suite.HTTPRouteExperimentalFeatures.UnsortedList()...).
			Insert(suite.TLSRouteCoreFeatures.UnsortedList()...).
			Insert(suite.MeshCoreFeatures.UnsortedList()...),
		ExemptFeatures:             exemptFeatures,
		EnableAllSupportedFeatures: *flags.EnableAllSupportedFeatures,
		NamespaceLabels:            namespaceLabels,
		NamespaceAnnotations:       namespaceAnnotations,
		SkipTests:                  skipTests,
		RunTest:                    *flags.RunTest,
	})
	t := testing.T{}
	cSuite.Setup(&t)

	if len(tests.ConformanceTests) == 0 {
		c.Fatal("Empty")
	}

	cSuite.Run(&t, tests.ConformanceTests)
}
