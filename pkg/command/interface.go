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

package command

import "github.com/CloudCoreo/cli/client"

// Interface for Coreo client for mocking in tests
type Interface interface {
	ListTeams() ([]*client.Team, error)
	ShowTeamByID(teamID string) (*client.Team, error)
	CreateTeam(teamName, teamDescripton string) (*client.Team, error)

	ListTokens() ([]*client.Token, error)
	ShowTokenByID(tokenID string) (*client.Token, error)
	DeleteTokenByID(tokenID string) error

	ListCloudAccounts(teamID string) ([]*client.CloudAccount, error)
	ShowCloudAccountByID(teamID, cloudID string) (*client.CloudAccount, error)
	CreateCloudAccount(input *client.CreateCloudAccountInput) (*client.CloudAccount, error)
	UpdateCloudAccount(input *client.UpdateCloudAccountInput) (*client.CloudAccount, error)
	DeleteCloudAccountByID(teamID, cloudID string) error
	ReValidateRole(teamID, cloudID string) (*client.RoleReValidationResult, error)

	GetEventStreamConfig(teamID, cloudID string) (*client.EventStreamConfig, error)
	GetEventRemoveConfig(teamID, cloudID string) (*client.EventRemoveConfig, error)
	GetRoleCreationInfo(input *client.CreateCloudAccountInput) (*client.RoleCreationInfo, error)
}

//CloudProvider for adding cloud account
type CloudProvider interface {
	SetupEventStream(input *client.EventStreamConfig) error
	GetOrgTree() ([]*TreeNode, error)
	CreateNewRole(input *client.RoleCreationInfo) (arn string, externalID string, err error)
	DeleteRole(roleName string)
	RemoveEventStream(input *client.EventRemoveConfig) error
}
