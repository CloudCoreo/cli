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
	CreateCloudAccount(teamID, resourceKey, resourceSecret, resourceName string) (*client.CloudAccount, error)
	DeleteCloudAccountByID(teamID, cloudID string) error

	ListGitKeys(teamID string) ([]*client.GitKey, error)
	ShowGitKeyByID(teamID, gitKeyID string) (*client.GitKey, error)
	CreateGitKey(teamID, resourceSecret, resourceName string) (*client.GitKey, error)
	DeleteGitKeyByID(teamID, gitKeyID string) error

	ListComposites(teamID string) ([]*client.Composite, error)
	ShowCompositeByID(teamID, compositeID string) (*client.Composite, error)
	CreateComposite(teamID, gitRepoURL, name, gitKeyId string) (*client.Composite, error)
	DeleteCompositeByID(teamID, compositeID string) error

	ListPlans(teamID, compositeID string) ([]*client.Plan, error)
	ShowPlanByID(teamID, compositeID, planID string) (*client.Plan, error)
	EnablePlanByID(teamID, compositeID, planID string) (*client.Plan, error)
	DisablePlanByID(teamID, compositeID, planID string) (*client.Plan, error)
	DeletePlanByID(teamID, compositeID, planID string) error
	InitPlan(branch, name, region, teamID, cloudID, compositeID, revision string, interval int) (*client.PlanConfig, error)
	CreatePlan(planConfigJSON []byte) (*client.Plan, error)
	GetPlanPanel(teamID, compositeID, planID string) (*client.Panel, error)
	RunNowPlanByID(teamID, compositeID, planID string, block bool) (*client.Plan, error)
}
