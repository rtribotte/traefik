package tcp

import (
	"context"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/pires/go-proxyproto"
	"github.com/traefik/traefik/v2/pkg/config/dynamic"
	"github.com/traefik/traefik/v2/pkg/log"
)

// Proxy forwards a TCP request to a TCP service.
type Proxy struct {
	address          string
	target           *net.TCPAddr
	terminationDelay time.Duration
	proxyProtocol    *dynamic.ProxyProtocol
	refreshTarget    bool
}

// NewProxy creates a new Proxy.
func NewProxy(address string, terminationDelay time.Duration, proxyProtocol *dynamic.ProxyProtocol) (*Proxy, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return nil, err
	}

	if proxyProtocol != nil && (proxyProtocol.Version < 1 || proxyProtocol.Version > 2) {
		return nil, fmt.Errorf("unknown proxyProtocol version: %d", proxyProtocol.Version)
	}

	// enable the refresh of the target only if the address in an IP
	refreshTarget := false
	if host, _, err := net.SplitHostPort(address); err == nil && net.ParseIP(host) == nil {
		refreshTarget = true
	}

	return &Proxy{
		address:          address,
		target:           tcpAddr,
		refreshTarget:    refreshTarget,
		terminationDelay: terminationDelay,
		proxyProtocol:    proxyProtocol,
	}, nil
}

// ServeTCP forwards the connection to a service.
func (p *Proxy) ServeTCP(conn WriteCloser) {
	log.Debugf("Handling connection from %s", conn.RemoteAddr())

	// needed because of e.g. server.trackedConnection
	defer conn.Close()

	if p.refreshTarget {
		tcpAddr, err := net.ResolveTCPAddr("tcp", p.address)
		if err != nil {
			log.Errorf("Error resolving tcp address: %v", err)
			return
		}
		p.target = tcpAddr
	}

	connBackend, err := net.DialTCP("tcp", nil, p.target)
	if err != nil {
		log.Errorf("Error while connection to backend: %v", err)
		return
	}

	closeWriter := &connWrapper{
		TCPConn: *connBackend,
		ctx:     conn.Context(),
	}

	// maybe not needed, but just in case
	defer closeWriter.Close()
	errChan := make(chan error)

	if p.proxyProtocol != nil && p.proxyProtocol.Version > 0 && p.proxyProtocol.Version < 3 {
		header := proxyproto.HeaderProxyFromAddrs(byte(p.proxyProtocol.Version), conn.RemoteAddr(), conn.LocalAddr())
		if _, err := header.WriteTo(closeWriter); err != nil {
			log.FromContext(conn.Context()).Errorf("Error while writing proxy protocol headers to backend connection: %v", err)
			return
		}
	}

	go p.connCopy(conn, closeWriter, errChan)
	go p.connCopy(closeWriter, conn, errChan)

	err = <-errChan
	if err != nil {
		log.FromContext(conn.Context()).Errorf("Error during connection: %v", err)
	}

	<-errChan
}

func (p Proxy) connCopy(dst, src WriteCloser, errCh chan error) {
	_, err := io.Copy(dst, src)
	errCh <- err

	errClose := dst.CloseWrite()
	if errClose != nil {
		log.FromContext(dst.Context()).Debugf("Error while terminating connection: %v", errClose)
		return
	}

	if p.terminationDelay >= 0 {
		err := dst.SetReadDeadline(time.Now().Add(p.terminationDelay))
		if err != nil {
			log.FromContext(dst.Context()).Debugf("Error while setting deadline: %v", err)
		}
	}
}

type connWrapper struct {
	net.TCPConn
	ctx context.Context
}

func (w *connWrapper) Context() context.Context {
	return w.ctx
}
