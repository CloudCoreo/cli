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
	cloudAccounts    []*client.CloudAccount
	config           client.EventStreamConfig
	err              error
	info             client.RoleCreationInfo
	regions          []string
	validationResult client.RoleReValidationResult
}

func (c *fakeReleaseClient) ListCloudAccounts() ([]*client.CloudAccount, error) {
	resp := c.cloudAccounts

	return resp, c.err
}

func (c *fakeReleaseClient) ShowCloudAccountByID(accountNumber, provider string) (*client.CloudAccount, error) {
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

func (c *fakeReleaseClient) DeleteCloudAccountByID(accountNumber, provider string) error {
	return c.err
}

func (c *fakeReleaseClient) GetEventStreamConfig(accountNumber, provider string) (*client.EventStreamConfig, error) {
	return &client.EventStreamConfig{
		AWSEventStreamConfig: client.AWSEventStreamConfig{Regions: c.regions},
	}, c.err
}

func (c *fakeReleaseClient) GetEventRemoveConfig(accountNumber, provider string) (*client.EventRemoveConfig, error) {
	return &client.EventRemoveConfig{
		AWSEventRemoveConfig: client.AWSEventRemoveConfig{
			Regions: c.regions,
		},
	}, c.err
}

func (c *fakeReleaseClient) GetRoleCreationInfo(input *client.CreateCloudAccountInput) (*client.RoleCreationInfo, error) {
	resp := c.info
	return &resp, c.err
}

func (c *fakeReleaseClient) UpdateCloudAccount(input *client.UpdateCloudAccountInput) (*client.CloudAccount, error) {
	resp := &client.CloudAccount{}
	if len(c.cloudAccounts) > 0 {

		resp = c.cloudAccounts[0]
	}
	return resp, c.err
}

func (c *fakeReleaseClient) ReValidateRole(accountNumber, provider string) (*client.RoleReValidationResult, error) {
	resp := c.validationResult
	return &resp, c.err
}

type fakeCloudProvider struct {
	err        error
	arn        string
	externalID string
}

func (c *fakeCloudProvider) SetupEventStream(input *client.EventStreamConfig) error {

	return c.err
}

func (c *fakeCloudProvider) CreateNewRole(input *client.RoleCreationInfo) (arn string, externalID string, err error) {
	return c.arn, c.externalID, c.err
}

func (c *fakeCloudProvider) DeleteRole(roleName string) {

}
func (c *fakeCloudProvider) RemoveEventStream(input *client.EventRemoveConfig) error {
	return c.err
}
