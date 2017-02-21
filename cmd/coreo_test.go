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

package main

import (
	"github.com/CloudCoreo/cli/client"
)

type fakeReleaseClient struct {
	teams         []*client.Team
	tokens        []*client.Token
	cloudAccounts []*client.CloudAccount
	gitKeys       []*client.GitKey
	composites    []*client.Composite
	plans         []*client.Plan
	planConfig    []*client.PlanConfig
	panel         []*client.Panel
	err           error
}

func (c *fakeReleaseClient) ListTeams() ([]*client.Team, error) {
	resp := c.teams

	return resp, c.err
}

func (c *fakeReleaseClient) ShowTeamByID(teamID string) (*client.Team, error) {
	resp := &client.Team{}
	if len(c.teams) > 0 {

		resp = c.teams[0]
	}

	return resp, c.err
}

func (c *fakeReleaseClient) ListTokens() ([]*client.Token, error) {
	resp := c.tokens

	return resp, c.err
}

func (c *fakeReleaseClient) ShowTokenByID(tokenID string) (*client.Token, error) {
	resp := &client.Token{}
	if len(c.tokens) > 0 {

		resp = c.tokens[0]
	}

	return resp, c.err
}

func (c *fakeReleaseClient) DeleteTokenByID(tokenID string) error {
	return c.err
}

func (c *fakeReleaseClient) ListCloudAccounts(teamID string) ([]*client.CloudAccount, error) {
	resp := c.cloudAccounts

	return resp, c.err
}

func (c *fakeReleaseClient) ShowCloudAccountByID(teamID, cloudID string) (*client.CloudAccount, error) {
	resp := &client.CloudAccount{}
	if len(c.cloudAccounts) > 0 {

		resp = c.cloudAccounts[0]
	}

	return resp, c.err
}

func (c *fakeReleaseClient) CreateCloudAccount(teamID, resourceKey, resourceSecret, resourceName string) (*client.CloudAccount, error) {
	resp := &client.CloudAccount{}
	if len(c.cloudAccounts) > 0 {

		resp = c.cloudAccounts[0]
	}

	return resp, c.err
}

func (c *fakeReleaseClient) DeleteCloudAccountByID(teamID, cloudID string) error {
	return c.err
}

func (c *fakeReleaseClient) ListGitKeys(teamID string) ([]*client.GitKey, error) {
	resp := c.gitKeys

	return resp, c.err
}

func (c *fakeReleaseClient) ShowGitKeyByID(teamID, gitKeyID string) (*client.GitKey, error) {
	resp := &client.GitKey{}
	if len(c.gitKeys) > 0 {

		resp = c.gitKeys[0]
	}

	return resp, c.err
}

func (c *fakeReleaseClient) CreateGitKey(teamID, resourceSecret, resourceName string) (*client.GitKey, error) {
	resp := &client.GitKey{}
	if len(c.gitKeys) > 0 {

		resp = c.gitKeys[0]
	}

	return resp, c.err
}

func (c *fakeReleaseClient) DeleteGitKeyByID(teamID, gitKeyID string) error {
	return c.err
}

func (c *fakeReleaseClient) ListComposites(teamID string) ([]*client.Composite, error) {
	resp := c.composites

	return resp, c.err
}

func (c *fakeReleaseClient) ShowCompositeByID(teamID, compositeID string) (*client.Composite, error) {
	resp := &client.Composite{}
	if len(c.composites) > 0 {

		resp = c.composites[0]
	}

	return resp, c.err
}

func (c *fakeReleaseClient) CreateComposite(teamID, gitRepoURL, name string) (*client.Composite, error) {
	resp := &client.Composite{}
	if len(c.composites) > 0 {

		resp = c.composites[0]
	}

	return resp, c.err
}

func (c *fakeReleaseClient) ListPlans(teamID, compositeID string) ([]*client.Plan, error) {
	resp := c.plans

	return resp, c.err
}

func (c *fakeReleaseClient) ShowPlanByID(teamID, compositeID, planID string) (*client.Plan, error) {
	resp := &client.Plan{}
	if len(c.plans) > 0 {

		resp = c.plans[0]
	}

	return resp, c.err
}

func (c *fakeReleaseClient) EnablePlanByID(teamID, compositeID, planID string) (*client.Plan, error) {
	resp := &client.Plan{}
	if len(c.plans) > 0 {

		resp = c.plans[0]
	}

	return resp, c.err
}

func (c *fakeReleaseClient) DisablePlanByID(teamID, compositeID, planID string) (*client.Plan, error) {
	resp := &client.Plan{}
	if len(c.plans) > 0 {

		resp = c.plans[0]
	}

	return resp, c.err
}

func (c *fakeReleaseClient) DeletePlanByID(teamID, compositeID, planID string) error {
	return c.err
}

func (c *fakeReleaseClient) InitPlan(branch, name, region, teamID, cloudID, compositeID, revision string, interval int) (*client.PlanConfig, error) {
	resp := &client.PlanConfig{}
	if len(c.planConfig) > 0 {

		resp = c.planConfig[0]
	}
	return resp, c.err
}

func (c *fakeReleaseClient) CreatePlan(planConfigJSON []byte) (*client.Plan, error) {
	resp := &client.Plan{}
	if len(c.plans) > 0 {

		resp = c.plans[0]
	}

	return resp, c.err
}

func (c *fakeReleaseClient) GetPlanPanel(teamID, compositeID, planID string) (*client.Panel, error) {
	resp := &client.Panel{}

	if len(c.panel) > 0 {

		resp = c.panel[0]
	}

	return resp, c.err
}
