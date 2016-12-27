package coreo

import (
	"context"
)

// Option allows specifying various settings configurable by
// the coreo client user for overriding the defaults
type Option func(*options)

type options struct {
	host      string
	apiKey    string
	secretKey string
}

// Host specifies the host address of the Coreo API server.
func Host(host string) Option {
	return func(opts *options) {
		opts.host = host
	}
}

//APIKey specifies the apiKey of the Coreo API server request.
func APIKey(apiKey string) Option {
	return func(opts *options) {
		opts.apiKey = apiKey
	}
}

// SecretKey specifies the secretKey of the Coreo API server request.
func SecretKey(secretKey string) Option {
	return func(opts *options) {
		opts.secretKey = secretKey
	}
}

// NewContext creates a versioned context.
func NewContext() context.Context {
	return context.Background()
}
