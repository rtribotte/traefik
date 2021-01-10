package dynamic

import (
	ptypes "github.com/traefik/paerser/types"
)

// +k8s:deepcopy-gen=true

// Middleware holds the Middleware configuration.
type TCPMiddleware struct {
	Chain       *TCPChain       `json:"chain,omitempty" toml:"chain,omitempty" yaml:"chain,omitempty" export:"true"`
	IPWhiteList *TCPIPWhiteList `json:"ipWhiteList,omitempty" toml:"ipWhiteList,omitempty" yaml:"ipWhiteList,omitempty" export:"true"`
	Retry       *TCPRetry       `json:"retry,omitempty" toml:"retry,omitempty" yaml:"retry,omitempty" export:"true"`
}

// +k8s:deepcopy-gen=true

// TCPChain holds a chain of middlewares.
type TCPChain struct {
	Middlewares []string `json:"middlewares,omitempty" toml:"middlewares,omitempty" yaml:"middlewares,omitempty" export:"true"`
}

// +k8s:deepcopy-gen=true

// TCPIPWhiteList holds the TCP ip white list configuration.
type TCPIPWhiteList struct {
	SourceRange []string `json:"sourceRange,omitempty" toml:"sourceRange,omitempty" yaml:"sourceRange,omitempty"`
}

// +k8s:deepcopy-gen=true

// Retry holds the TCP retry configuration.
type TCPRetry struct {
	Attempts        int             `json:"attempts,omitempty" toml:"attempts,omitempty" yaml:"attempts,omitempty" export:"true"`
	InitialInterval ptypes.Duration `json:"initialInterval,omitempty" toml:"initialInterval,omitempty" yaml:"initialInterval,omitempty" export:"true"`
}
