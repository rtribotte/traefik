package tls

import (
	"crypto/tls"
	"sync"
)

// TLSConfigStore store for dynamic tls.Config.
type TLSConfigStore struct {
	// tls.Config keyed by host
	hostTLSConfigs map[string]*tls.Config
	lock           sync.RWMutex
}

func NewTLSConfigStore() *TLSConfigStore {
	return &TLSConfigStore{
		hostTLSConfigs: make(map[string]*tls.Config),
	}
}

// AddConfig defines the tls.Config for the given sniHost.
func (s *TLSConfigStore) AddConfig(sniHost string, config *tls.Config) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.hostTLSConfigs == nil {
		s.hostTLSConfigs = map[string]*tls.Config{}
	}
	s.hostTLSConfigs[sniHost] = config
}

// GetConfigs returns a collection of the tls.Config keyed by host.
func (s *TLSConfigStore) GetConfigs() map[string]*tls.Config {
	s.lock.RLock()
	defer s.lock.RUnlock()

	result := map[string]*tls.Config{}
	for sniHost, config := range s.hostTLSConfigs {
		result[sniHost] = config.Clone()
	}

	return result
}

// GetConfig returns the tls.Config for the given host.
func (s *TLSConfigStore) GetConfig(sniHost string) *tls.Config {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if tlsConfig, ok := s.hostTLSConfigs[sniHost]; ok {
		return tlsConfig
	}

	return nil
}
