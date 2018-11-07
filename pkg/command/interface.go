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

// Interface for Coreo client for mocking in tests
type Interface interface {
	ListTeams() ([]*Team, error)
	ShowTeamByID(teamID string) (*Team, error)
	CreateTeam(teamName, teamDescripton string) (*Team, error)

	ListTokens() ([]*Token, error)
	ShowTokenByID(tokenID string) (*Token, error)
	DeleteTokenByID(tokenID string) error

	ListCloudAccounts(teamID string) ([]*CloudAccount, error)
	ShowCloudAccountByID(teamID, cloudID string) (*CloudAccount, error)
	CreateCloudAccount(input *CreateCloudAccountInput) (*CloudAccount, error)
	DeleteCloudAccountByID(teamID, cloudID string) error

	ShowResultObject(teamID, cloudID, level string) ([]*ResultObject, error)
	ShowResultRule(teamID, cloudID, level string) ([]*ResultRule, error)

	SetupEventStream(input *SetupEventStreamInput) error
}
