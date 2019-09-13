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

//ListTeams get list of teams
func (c *Client) ListTeams() ([]*client.Team, error) {
	ctx := NewContext()
	clt, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	teams, err := clt.GetTeams(ctx)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

//ShowTeamByID show team with ID
func (c *Client) ShowTeamByID(teamID string) (*client.Team, error) {
	ctx := NewContext()
	clt, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	team, err := clt.GetTeamByID(ctx, teamID)
	if err != nil {
		return nil, err
	}

	return team, nil
}

//CreateTeam Create a new team
func (c *Client) CreateTeam(teamName, teamDescription string) (*client.Team, error) {
	ctx := NewContext()
	clt, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	team, err := clt.CreateTeam(ctx, teamName, teamDescription)
	if err != nil {
		return nil, err
	}

	return team, nil
}

//ListTokens get tokens list
func (c *Client) ListTokens() ([]*client.Token, error) {
	ctx := NewContext()
	clt, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	tokens, err := clt.GetTokens(ctx)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

//ShowTokenByID show token by ID
func (c *Client) ShowTokenByID(tokenID string) (*client.Token, error) {
	ctx := NewContext()
	clt, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	token, err := clt.GetTokenByID(ctx, tokenID)
	if err != nil {
		return nil, err
	}

	return token, nil
}

//DeleteTokenByID Delete token by ID
func (c *Client) DeleteTokenByID(tokenID string) error {
	ctx := NewContext()
	clt, err := c.MakeClient()
	if err != nil {
		return err
	}

	err = clt.DeleteTokenByID(ctx, tokenID)
	if err != nil {
		return err
	}

	return nil
}

//ListCloudAccounts Get list of cloud accounts
func (c *Client) ListCloudAccounts(teamID string) ([]*client.CloudAccount, error) {
	ctx := NewContext()
	clt, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	cloudAccounts, err := clt.GetCloudAccounts(ctx, teamID)
	if err != nil {
		return nil, err
	}

	return cloudAccounts, nil
}

//ShowCloudAccountByID show cloud account by ID
func (c *Client) ShowCloudAccountByID(teamID, cloudID string) (*client.CloudAccount, error) {
	ctx := NewContext()
	clt, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	cloudAccount, err := clt.GetCloudAccountByID(ctx, teamID, cloudID)
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
func (c *Client) DeleteCloudAccountByID(teamID, cloudID string) error {
	ctx := NewContext()
	clt, err := c.MakeClient()
	if err != nil {
		return err
	}

	err = clt.DeleteCloudAccountByID(ctx, teamID, cloudID)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ReValidateRole(teamID, cloudID string) (*client.RoleReValidationResult, error) {
	ctx := NewContext()
	clt, err := c.MakeClient()
	if err != nil {
		return nil, err
	}
	result, err := clt.ReValidateRole(ctx, teamID, cloudID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//GetEventStreamConfig gets event stream setup config
func (c *Client) GetEventStreamConfig(teamID, cloudID string) (*client.EventStreamConfig, error) {
	ctx := NewContext()
	clt, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	return clt.GetSetupConfig(ctx, teamID, cloudID)
}

func (c *Client) GetEventRemoveConfig(teamID, cloudID string) (*client.EventRemoveConfig, error) {
	ctx := NewContext()
	clt, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	return clt.GetRemoveConfig(ctx, teamID, cloudID)
}

func (c *Client) GetRoleCreationInfo(input *client.CreateCloudAccountInput) (*client.RoleCreationInfo, error) {
	ctx := NewContext()
	clt, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	return clt.GetRoleCreationInfo(ctx, input)
}
