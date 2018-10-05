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
	return client.MakeClient(c.opts.apiKey, c.opts.secretKey, c.opts.host)
}

//ListTeams get list of teams
func (c *Client) ListTeams() ([]*client.Team, error) {
	ctx := NewContext()
	client, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	teams, err := client.GetTeams(ctx)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

//ShowTeamByID show team with ID
func (c *Client) ShowTeamByID(teamID string) (*client.Team, error) {
	ctx := NewContext()
	client, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	team, err := client.GetTeamByID(ctx, teamID)
	if err != nil {
		return nil, err
	}

	return team, nil
}

//CreateTeam Create a new team
func (c *Client) CreateTeam(teamName, teamDescription string) (*client.Team, error) {
	ctx := NewContext()
	client, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	team, err := client.CreateTeam(ctx, teamName, teamDescription)
	if err != nil {
		return nil, err
	}

	return team, nil
}

//ListTokens get tokens list
func (c *Client) ListTokens() ([]*client.Token, error) {
	ctx := NewContext()
	client, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	tokens, err := client.GetTokens(ctx)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

//ShowTokenByID show token by ID
func (c *Client) ShowTokenByID(tokenID string) (*client.Token, error) {
	ctx := NewContext()
	client, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	token, err := client.GetTokenByID(ctx, tokenID)
	if err != nil {
		return nil, err
	}

	return token, nil
}

//DeleteTokenByID Delete token by ID
func (c *Client) DeleteTokenByID(tokenID string) error {
	ctx := NewContext()
	client, err := c.MakeClient()
	if err != nil {
		return err
	}

	err = client.DeleteTokenByID(ctx, tokenID)
	if err != nil {
		return err
	}

	return nil
}

//ListCloudAccounts Get list of cloud accounts
func (c *Client) ListCloudAccounts(teamID string) ([]*client.CloudAccount, error) {
	ctx := NewContext()
	client, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	cloudAccounts, err := client.GetCloudAccounts(ctx, teamID)
	println("Get cloud accounts ")
	if err != nil {
		return nil, err
	}

	return cloudAccounts, nil
}

//ShowCloudAccountByID show cloud account by ID
func (c *Client) ShowCloudAccountByID(teamID, cloudID string) (*client.CloudAccount, error) {
	ctx := NewContext()
	client, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	cloudAccount, err := client.GetCloudAccountByID(ctx, teamID, cloudID)
	if err != nil {
		return nil, err
	}

	return cloudAccount, nil
}

//CreateCloudAccount Create cloud account
func (c *Client) CreateCloudAccount(teamID, resourceKey, resourceSecret, resourceName string) (*client.CloudAccount, error) {
	ctx := NewContext()
	client, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	cloudAccount, err := client.CreateCloudAccount(ctx, teamID, resourceKey, resourceSecret, resourceName)
	if err != nil {
		return nil, err
	}

	return cloudAccount, nil

}

//DeleteCloudAccountByID Delete cloud by ID
func (c *Client) DeleteCloudAccountByID(teamID, cloudID string) error {
	ctx := NewContext()
	client, err := c.MakeClient()
	if err != nil {
		return err
	}

	err = client.DeleteCloudAccountByID(ctx, teamID, cloudID)
	if err != nil {
		return err
	}

	return nil
}

//Show violated rules. If the filter condition (teamID, cloudID in this case) is valid,
//rules will be filtered. Otherwise return all violation rules under this user account.
func (c *Client) ShowResultRule(teamID, cloudID string) ([]* client.ResultRule, error) {
	//TODO
	ctx := NewContext()
	client, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	result, err := client.ShowResultRule(ctx, teamID, cloudID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//Show violated objects. If the filter condition (teamID, cloudID in this case) is valid,
//objects will be filtered. Otherwise return all violation objects under this user account.
func (c *Client) ShowResultObject(teamID, cloudID string) ([]* client.ResultObject, error) {
	//TODO
	ctx := NewContext()
	client, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	result, err := client.ShowResultObject(ctx, teamID, cloudID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
