// Copyright © 2016 Paul Allen <paul@cloudcoreo.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
