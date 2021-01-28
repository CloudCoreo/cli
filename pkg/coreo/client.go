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
	"github.com/CloudCoreo/cli/client"
)

// Client struct
type Client struct {
	opts options
}

// NewClient creates a new client.
func NewClient(opts ...Option) *Client {
	var c Client
	return c.Option(opts...)
}

// Option configures the Coreo client with the provided options
func (c *Client) Option(opts ...Option) *Client {
	for _, opt := range opts {
		opt(&c.opts)
	}
	return c
}

//MakeClient make client method
func (c *Client) MakeClient() (*client.Client, error) {
	return client.MakeClient(c.opts.refreshToken, c.opts.host)
}

//ListCloudAccounts Get list of cloud accounts
func (c *Client) ListCloudAccounts() ([]*client.CloudAccount, error) {
	ctx := NewContext()
	clt, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	cloudAccounts, err := clt.GetCloudAccounts(ctx)
	if err != nil {
		return nil, err
	}

	return cloudAccounts, nil
}

//ShowCloudAccountByID show cloud account by ID
func (c *Client) ShowCloudAccountByID(accountNumber, provider string) (*client.CloudAccount, error) {
	ctx := NewContext()
	clt, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	cloudAccount, err := clt.GetCloudAccountByID(ctx, accountNumber, provider)
	if err != nil {
		return nil, err
	}

	return cloudAccount, nil
}

//CreateCloudAccount Create cloud account
func (c *Client) CreateCloudAccount(input *client.CreateCloudAccountInput) (*client.CloudAccount, error) {
	ctx := NewContext()
	clt, err := c.MakeClient()
	if err != nil {
		return nil, err
	}
	cloudAccount, err := clt.CreateCloudAccount(ctx, input)
	if err != nil {
		return nil, err
	}

	return cloudAccount, nil

}

func (c *Client) UpdateCloudAccount(input *client.UpdateCloudAccountInput) (*client.CloudAccount, error) {
	ctx := NewContext()
	clt, err := c.MakeClient()
	if err != nil {
		return nil, err
	}
	cloudAccount, err := clt.UpdateCloudAccount(ctx, input)
	if err != nil {
		return nil, err
	}

	return cloudAccount, nil
}

//DeleteCloudAccountByID Delete cloud by ID
func (c *Client) DeleteCloudAccountByID(accountNumber, provider string) error {
	ctx := NewContext()
	clt, err := c.MakeClient()
	if err != nil {
		return err
	}

	err = clt.DeleteCloudAccountByID(ctx, accountNumber, provider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ReValidateRole(accountNumber, provider string) (*client.RoleReValidationResult, error) {
	ctx := NewContext()
	clt, err := c.MakeClient()
	if err != nil {
		return nil, err
	}
	result, err := clt.ReValidateRole(ctx, accountNumber, provider)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//GetEventStreamConfig gets event stream setup config
func (c *Client) GetEventStreamConfig(accountNumber, provider string) (*client.EventStreamConfig, error) {
	ctx := NewContext()
	clt, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	return clt.GetSetupConfig(ctx, accountNumber, provider)
}

func (c *Client) GetEventRemoveConfig(accountNumber, provider string) (*client.EventRemoveConfig, error) {
	ctx := NewContext()
	clt, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	return clt.GetRemoveConfig(ctx, accountNumber, provider)
}

func (c *Client) GetRoleCreationInfo(input *client.CreateCloudAccountInput) (*client.RoleCreationInfo, error) {
	ctx := NewContext()
	clt, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	return clt.GetRoleCreationInfo(ctx, input)
}
