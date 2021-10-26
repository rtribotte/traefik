package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/lucas-clemente/quic-go/http3"
	"github.com/traefik/traefik/v2/pkg/config/static"
	"github.com/traefik/traefik/v2/pkg/log"
	traefiktls "github.com/traefik/traefik/v2/pkg/tls"
)

type http3server struct {
	*http3.Server

	http3conn net.PacketConn
}

func newHTTP3Server(ctx context.Context, configuration *static.EntryPoint, httpsServer *httpServer, tlsConfigLookup traefiktls.ConfigLookup) (*http3server, error) {
	if configuration.HTTP3 == nil {
		return nil, nil
	}

	conn, err := net.ListenPacket("udp", configuration.GetAddress())
	if err != nil {
		return nil, fmt.Errorf("error while starting http3 listener: %w", err)
	}

	h3 := &http3server{
		http3conn: conn,
	}

	h3.Server = &http3.Server{
		Server: &http.Server{
			Addr:         configuration.GetAddress(),
			Handler:      httpsServer.Server.(*http.Server).Handler,
			ErrorLog:     httpServerLogger,
			ReadTimeout:  time.Duration(configuration.Transport.RespondingTimeouts.ReadTimeout),
			WriteTimeout: time.Duration(configuration.Transport.RespondingTimeouts.WriteTimeout),
			IdleTimeout:  time.Duration(configuration.Transport.RespondingTimeouts.IdleTimeout),
			TLSConfig:    &tls.Config{GetConfigForClient: tlsConfigLookup},
		},
	}

	previousHandler := httpsServer.Server.(*http.Server).Handler

	setQuicHeaders := getQuicHeadersSetter(configuration)

	httpsServer.Server.(*http.Server).Handler = http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		err := setQuicHeaders(rw.Header())
		if err != nil {
			log.FromContext(ctx).Errorf("failed to set HTTP3 headers: %v", err)
		}

		previousHandler.ServeHTTP(rw, req)
	})

	return h3, nil
}

// TODO: rewrite if at some point `port` become an exported field of http3.Server.
func getQuicHeadersSetter(configuration *static.EntryPoint) func(header http.Header) error {
	advertisedAddress := configuration.GetAddress()
	if configuration.HTTP3.AdvertisedPort != 0 {
		advertisedAddress = fmt.Sprintf(`:%d`, configuration.HTTP3.AdvertisedPort)
	}

	// if `QuickConfig` of h3.server happens to be configured,
	// it should also be configured identically in the headerServer
	headerServer := &http3.Server{
		Server: &http.Server{
			Addr: advertisedAddress,
		},
	}

	// set quic headers with the "header" http3 server instance
	return headerServer.SetQuicHeaders
}

func (e *http3server) Start() error {
	return e.Serve(e.http3conn)
}
