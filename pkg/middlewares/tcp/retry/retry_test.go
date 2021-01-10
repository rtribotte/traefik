package tcpretry

import (
	"context"
	"io/ioutil"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	ptypes "github.com/traefik/paerser/types"
	"github.com/traefik/traefik/v2/pkg/config/dynamic"
	"github.com/traefik/traefik/v2/pkg/tcp"
)

func TestRetry(t *testing.T) {
	testCases := []struct {
		desc                  string
		config                dynamic.TCPRetry
		wantRetryAttempts     int
		expectedResponse      string
		amountFaultyEndpoints int
	}{
		{
			desc:                  "no retry on success",
			config:                dynamic.TCPRetry{Attempts: 1},
			wantRetryAttempts:     0,
			expectedResponse:      "OK",
			amountFaultyEndpoints: 0,
		},
		{
			desc:                  "no retry on success with backoff",
			config:                dynamic.TCPRetry{Attempts: 1, InitialInterval: ptypes.Duration(time.Microsecond * 50)},
			wantRetryAttempts:     0,
			expectedResponse:      "OK",
			amountFaultyEndpoints: 0,
		},
		{
			desc:                  "no retry when max connection attempts is one",
			config:                dynamic.TCPRetry{Attempts: 1},
			wantRetryAttempts:     0,
			amountFaultyEndpoints: 1,
		},
		{
			desc:                  "no retry when max connection attempts is one with backoff",
			config:                dynamic.TCPRetry{Attempts: 1, InitialInterval: ptypes.Duration(time.Microsecond * 50)},
			wantRetryAttempts:     0,
			amountFaultyEndpoints: 1,
		},
		{
			desc:                  "one retry when one server is faulty",
			config:                dynamic.TCPRetry{Attempts: 2},
			wantRetryAttempts:     1,
			expectedResponse:      "OK",
			amountFaultyEndpoints: 1,
		},
		{
			desc:                  "one retry when one server is faulty with backoff",
			config:                dynamic.TCPRetry{Attempts: 2, InitialInterval: ptypes.Duration(time.Microsecond * 50)},
			wantRetryAttempts:     1,
			expectedResponse:      "OK",
			amountFaultyEndpoints: 1,
		},
		{
			desc:                  "two retries when two servers are faulty",
			config:                dynamic.TCPRetry{Attempts: 3},
			wantRetryAttempts:     2,
			expectedResponse:      "OK",
			amountFaultyEndpoints: 2,
		},
		{
			desc:                  "two retries when two servers are faulty with backoff",
			config:                dynamic.TCPRetry{Attempts: 3, InitialInterval: ptypes.Duration(time.Microsecond * 50)},
			wantRetryAttempts:     2,
			expectedResponse:      "OK",
			amountFaultyEndpoints: 2,
		},
		{
			desc:                  "max attempts exhausted closes connection",
			config:                dynamic.TCPRetry{Attempts: 3},
			wantRetryAttempts:     2,
			amountFaultyEndpoints: 3,
		},
		{
			desc:                  "max attempts exhausted closes connection with backoff",
			config:                dynamic.TCPRetry{Attempts: 3, InitialInterval: ptypes.Duration(time.Microsecond * 50)},
			wantRetryAttempts:     2,
			amountFaultyEndpoints: 3,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			loadBalancer := tcp.NewWRRLoadBalancer()

			for i := 0; i < test.amountFaultyEndpoints; i++ {
				loadBalancer.AddServer(faultyServer{})
			}

			server, client := net.Pipe()
			validServer := tcp.HandlerFunc(func(conn tcp.WriteCloser) {
				write, err := conn.Write([]byte("OK"))
				require.NoError(t, err)
				assert.Equal(t, 2, write)

				err = conn.Close()
				require.NoError(t, err)
			})

			loadBalancer.AddServer(validServer)

			retryListener := &countingRetryListener{}
			retry, err := New(context.Background(), loadBalancer, test.config, retryListener, "traefikTest")
			require.NoError(t, err)

			go func() {
				retry.ServeTCP(&contextWriteCloser{client})
			}()

			read, err := ioutil.ReadAll(server)
			require.NoError(t, err)

			assert.Equal(t, test.expectedResponse, string(read))

			assert.Equal(t, test.wantRetryAttempts, retryListener.timesCalled)
		})
	}
}

type faultyServer struct {
}

func (f faultyServer) ServeTCP(conn tcp.WriteCloser) {
	conn.Close()
}

type contextWriteCloser struct {
	net.Conn
}

func (c contextWriteCloser) CloseWrite() error {
	panic("implement me")
}

func (c contextWriteCloser) Context() context.Context {
	return context.Background()
}

func (c contextWriteCloser) WithContext(ctx context.Context) tcp.WriteCloser {
	panic("implement me")
}

func TestRetryEmptyServerList(t *testing.T) {
	loadBalancer := tcp.NewWRRLoadBalancer()

	server, client := net.Pipe()

	retryListener := &countingRetryListener{}
	retry, err := New(context.Background(), loadBalancer, dynamic.TCPRetry{Attempts: 3}, retryListener, "traefikTest")
	require.NoError(t, err)

	go func() {
		retry.ServeTCP(&contextWriteCloser{client})
	}()

	read, err := ioutil.ReadAll(server)
	require.NoError(t, err)

	assert.Equal(t, "", string(read))

	assert.Equal(t, 0, retryListener.timesCalled)

}

func TestRetryListeners(t *testing.T) {
	retryListeners := Listeners{&countingRetryListener{}, &countingRetryListener{}}

	retryListeners.Retried(1)
	retryListeners.Retried(1)

	for _, retryListener := range retryListeners {
		listener := retryListener.(*countingRetryListener)
		if listener.timesCalled != 2 {
			t.Errorf("retry listener was called %d time(s), want %d time(s)", listener.timesCalled, 2)
		}
	}
}

// countingRetryListener is a Listener implementation to count the times the Retried fn is called.
type countingRetryListener struct {
	timesCalled int
}

func (l *countingRetryListener) Retried(attempt int) {
	l.timesCalled++
}
