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

package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/CloudCoreo/cli/client/content"
	"golang.org/x/net/context/ctxhttp"
)

type clientOptions struct {
	interceptor Interceptor
}

// Option type
type Option func(*clientOptions)

// Interceptor is a generic request interceptor, useful for
// modifying or canceling the request.
type Interceptor func(*http.Request) error

// WithInterceptor returns a ClientOption for adding an interceptor
// to a Client.
func WithInterceptor(ci Interceptor) Option {
	return func(opts *clientOptions) {
		opts.interceptor = ci
	}
}

// Client struct
type Client struct {
	client   http.Client
	endpoint string
	opts     clientOptions
	auth     Auth
}

// MakeClient make client
func MakeClient(apiKey, secretKey, endpoint string) (*Client, error) {

	if apiKey == "None" || secretKey == "None" || apiKey == "" || secretKey == "" {
		return nil, NewError(content.ErrorMissingAPIOrSecretKey)
	}

	a := Auth{APIKey: apiKey, SecretKey: secretKey}
	i := Interceptor(a.SignRequest)
	c := newClient(endpoint, WithInterceptor(i))

	return c, nil
}

// New creates a new Client for a given endpoint that can be configured with
// multiple ClientOption
func newClient(endpoint string, opts ...Option) *Client {
	client := &Client{
		endpoint: endpoint,
	}

	for _, opt := range opts {
		opt(&client.opts)
	}

	return client
}

// Do performs an HTTP request with a given context - the response will be decoded
// into obj.
func (c *Client) Do(ctx context.Context, method, path string, body io.Reader, obj interface{}) error {
	req, err := c.buildRequest(method, path, body)
	if err != nil {
		return err
	}

	resp, err := ctxhttp.Do(ctx, &c.client, req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		message := new(bytes.Buffer)
		message.ReadFrom(resp.Body)
		msg := fmt.Sprintf("%s", message.String())
		return NewError(msg)
	}

	// Read all of resp.Body regardless of status code so we don't leak connections.
	// The extra io.Copy is to ensure everything has been read, since a json.Decoder doesn't
	// have that guarantee.
	if obj != nil {
		err = json.NewDecoder(resp.Body).Decode(obj)
	}

	io.Copy(ioutil.Discard, resp.Body)

	return err
}

func (c *Client) buildRequest(method, urlPath string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, urlPath, body)
	if err != nil {
		return nil, err
	}
	if method == "POST" || method == "PUT" {
		req.Header.Set("Content-Type", "application/json")
	}

	if c.opts.interceptor != nil {
		if err := c.opts.interceptor(req); err != nil {
			return nil, err
		}
	}

	return req, nil
}
