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
