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
	"github.com/CloudCoreo/cli/pkg/command"
)

type fakeReleaseClient struct {
	teams         []*client.Team
	tokens        []*client.Token
	cloudAccounts []*client.CloudAccount
	objects       []*client.ResultObject
	rules         []*client.ResultRule
	config        client.EventStreamConfig
	err           error
	info          client.RoleCreationInfo
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

func (c *fakeReleaseClient) CreateTeam(teamName, teamDescription string) (*client.Team, error) {
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

func (c *fakeReleaseClient) CreateCloudAccount(input *client.CreateCloudAccountInput) (*client.CloudAccount, error) {
	resp := &client.CloudAccount{}
	if len(c.cloudAccounts) > 0 {

		resp = c.cloudAccounts[0]
	}

	return resp, c.err
}

func (c *fakeReleaseClient) DeleteCloudAccountByID(teamID, cloudID string) error {
	return c.err
}

func (c *fakeReleaseClient) ShowResultRule(teamID, cloudID, level string) ([]*client.ResultRule, error) {
	resp := c.rules
	return resp, c.err
}

func (c *fakeReleaseClient) ShowResultObject(teamID, cloudID, level string) ([]*client.ResultObject, error) {
	resp := c.objects
	return resp, c.err
}

func (c *fakeReleaseClient) GetEventStreamConfig(teamID, cloudID string) (*client.EventStreamConfig, error) {
	return &client.EventStreamConfig{}, c.err
}

func (c *fakeReleaseClient) GetRoleCreationInfo(input *client.CreateCloudAccountInput) (*client.RoleCreationInfo, error) {
	resp := c.info
	return &resp, c.err
}

type fakeCloudProvider struct {
	err        error
	arn        string
	externalID string
	root       command.TreeNode
}

func (c *fakeCloudProvider) SetupEventStream(input *client.EventStreamConfig) error {

	return c.err
}

func (c *fakeCloudProvider) GetOrgTree() ([]*command.TreeNode, error) {
	resp := []*command.TreeNode{&c.root}
	return resp, c.err
}

func (c *fakeCloudProvider) CreateNewRole(input *client.RoleCreationInfo) (arn string, externalID string, err error) {
	return c.arn, c.externalID, c.err
}

func (c *fakeCloudProvider) DeleteRole(roleName, policyArn string) {

}
