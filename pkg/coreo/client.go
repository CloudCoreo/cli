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

//ListGitKeys Get list of git keys
func (c *Client) ListGitKeys(teamID string) ([]*client.GitKey, error) {
	ctx := NewContext()
	client, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	gitKeys, err := client.GetGitKeys(ctx, teamID)
	if err != nil {
		return nil, err
	}

	return gitKeys, nil
}

//ShowGitKeyByID Show gitkey with ID
func (c *Client) ShowGitKeyByID(teamID, gitKeyID string) (*client.GitKey, error) {
	ctx := NewContext()
	client, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	gitKey, err := client.GetGitKeyByID(ctx, teamID, gitKeyID)
	if err != nil {
		return nil, err
	}

	return gitKey, nil

}

//CreateGitKey Create git key
func (c *Client) CreateGitKey(teamID, resourceSecret, resourceName string) (*client.GitKey, error) {
	ctx := NewContext()
	client, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	gitKey, err := client.CreateGitKey(ctx, teamID, resourceSecret, resourceName)
	if err != nil {
		return nil, err
	}

	return gitKey, nil
}

//DeleteGitKeyByID delete git key
func (c *Client) DeleteGitKeyByID(teamID, gitKeyID string) error {
	ctx := NewContext()
	client, err := c.MakeClient()
	if err != nil {
		return err
	}

	err = client.DeleteGitKeyByID(ctx, teamID, gitKeyID)
	if err != nil {
		return err
	}

	return nil
}

//ListComposites List composite
func (c *Client) ListComposites(teamID string) ([]*client.Composite, error) {
	ctx := NewContext()
	client, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	composites, err := client.GetComposites(ctx, teamID)
	if err != nil {
		return nil, err
	}

	return composites, nil
}

//ShowCompositeByID show composite by ID
func (c *Client) ShowCompositeByID(teamID, compositeID string) (*client.Composite, error) {
	ctx := NewContext()
	client, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	composite, err := client.GetCompositeByID(ctx, teamID, compositeID)
	if err != nil {
		return nil, err
	}

	return composite, nil
}

//CreateComposite Create composite
func (c *Client) CreateComposite(teamID, gitRepoURL, name string) (*client.Composite, error) {
	ctx := NewContext()
	client, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	composite, err := client.CreateComposite(ctx, gitRepoURL, name, teamID)
	if err != nil {
		return nil, err
	}

	return composite, nil
}

//ListPlans List plans
func (c *Client) ListPlans(teamID, compositeID string) ([]*client.Plan, error) {
	ctx := NewContext()
	client, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	plans, err := client.GetPlans(ctx, teamID, compositeID)
	if err != nil {
		return nil, err
	}

	return plans, nil
}

//ShowPlanByID Show plan by ID
func (c *Client) ShowPlanByID(teamID, compositeID, planID string) (*client.Plan, error) {
	ctx := NewContext()
	client, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	plan, err := client.GetPlanByID(ctx, teamID, compositeID, planID)
	if err != nil {
		return nil, err
	}

	return plan, nil
}

//EnablePlanByID Enable plan by ID
func (c *Client) EnablePlanByID(teamID, compositeID, planID string) (*client.Plan, error) {
	ctx := NewContext()
	client, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	plan, err := client.EnablePlan(ctx, teamID, compositeID, planID)
	if err != nil {
		return nil, err
	}

	return plan, nil
}

//DisablePlanByID Disable plan by ID
func (c *Client) DisablePlanByID(teamID, compositeID, planID string) (*client.Plan, error) {
	ctx := NewContext()
	client, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	plan, err := client.DisablePlan(ctx, teamID, compositeID, planID)
	if err != nil {
		return nil, err
	}

	return plan, nil
}

//DeletePlanByID Delete by ID
func (c *Client) DeletePlanByID(teamID, compositeID, planID string) error {
	ctx := NewContext()
	client, err := c.MakeClient()
	if err != nil {
		return err
	}

	err = client.DeletePlanByID(ctx, teamID, compositeID, planID)
	if err != nil {
		return err
	}

	return nil
}

//InitPlan init a plan
func (c *Client) InitPlan(branch, name, region, teamID, cloudID, compositeID, revision string, interval int) (*client.PlanConfig, error) {
	ctx := NewContext()
	client, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	planConfig, err := client.InitPlan(ctx, branch, name, region, teamID, cloudID, compositeID, revision, interval)
	if err != nil {
		return nil, err
	}

	// Add value property
	for _, v := range planConfig.Variables {
		if v.Required {
			if v.Default != nil {
				v.Value = v.Default
			}
		}
	}

	return planConfig, nil
}

//CreatePlan create a plan
func (c *Client) CreatePlan(planConfigContent []byte) (*client.Plan, error) {

	ctx := NewContext()
	client, err := c.MakeClient()
	if err != nil {
		return nil, err
	}

	plan, err := client.CreatePlan(ctx, planConfigContent)
	if err != nil {
		return nil, err
	}

	return plan, nil
}
