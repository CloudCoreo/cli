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
	"github.com/CloudCoreo/cli/pkg/command"
)

type fakeReleaseClient struct {
	teams         []*command.Team
	tokens        []*command.Token
	cloudAccounts []*command.CloudAccount
	objects       []*command.ResultObject
	rules         []*command.ResultRule
	err           error
}

func (c *fakeReleaseClient) ListTeams() ([]*command.Team, error) {
	resp := c.teams

	return resp, c.err
}

func (c *fakeReleaseClient) ShowTeamByID(teamID string) (*command.Team, error) {
	resp := &command.Team{}
	if len(c.teams) > 0 {

		resp = c.teams[0]
	}

	return resp, c.err
}

func (c *fakeReleaseClient) CreateTeam(teamName, teamDescription string) (*command.Team, error) {
	resp := &command.Team{}
	if len(c.teams) > 0 {
		resp = c.teams[0]
	}

	return resp, c.err
}

func (c *fakeReleaseClient) ListTokens() ([]*command.Token, error) {
	resp := c.tokens

	return resp, c.err
}

func (c *fakeReleaseClient) ShowTokenByID(tokenID string) (*command.Token, error) {
	resp := &command.Token{}
	if len(c.tokens) > 0 {

		resp = c.tokens[0]
	}

	return resp, c.err
}

func (c *fakeReleaseClient) DeleteTokenByID(tokenID string) error {
	return c.err
}

func (c *fakeReleaseClient) ListCloudAccounts(teamID string) ([]*command.CloudAccount, error) {
	resp := c.cloudAccounts

	return resp, c.err
}

func (c *fakeReleaseClient) ShowCloudAccountByID(teamID, cloudID string) (*command.CloudAccount, error) {
	resp := &command.CloudAccount{}
	if len(c.cloudAccounts) > 0 {

		resp = c.cloudAccounts[0]
	}

	return resp, c.err
}

func (c *fakeReleaseClient) CreateCloudAccount(input *command.CreateCloudAccountInput) (*command.CloudAccount, error) {
	resp := &command.CloudAccount{}
	if len(c.cloudAccounts) > 0 {

		resp = c.cloudAccounts[0]
	}

	return resp, c.err
}

func (c *fakeReleaseClient) DeleteCloudAccountByID(teamID, cloudID string) error {
	return c.err
}

func (c *fakeReleaseClient) ShowResultRule(teamID, cloudID, level string) ([]*command.ResultRule, error) {
	resp := c.rules
	return resp, c.err
}

func (c *fakeReleaseClient) ShowResultObject(teamID, cloudID, level string) ([]*command.ResultObject, error) {
	resp := c.objects
	return resp, c.err
}

func (c *fakeReleaseClient) SetupEventStream(input *command.SetupEventStreamInput) error {
	return c.err
}
