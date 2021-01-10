package tcp

import (
	"context"
	"crypto/tls"
)

// TLSHandler handles TLS connections.
type TLSHandler struct {
	Next   Handler
	Config *tls.Config
}

// ServeTCP terminates the TLS connection.
func (t *TLSHandler) ServeTCP(conn WriteCloser) {
	newConn := tls.Server(conn, t.Config)

	t.Next.ServeTCP(&tlsConnWrapper{
		Conn: newConn,
		ctx:  conn.Context(),
	})
}

type tlsConnWrapper struct {
	*tls.Conn
	ctx context.Context
}

func (w *tlsConnWrapper) Context() context.Context {
	return w.ctx
}
