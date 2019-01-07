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

package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/CloudCoreo/cli/client/content"
)

// CloudAccount Information
type CloudAccount struct {
	TeamID    string `json:"teamId"`
	Name      string `json:"name"`
	RoleID    string `json:"roleId"`
	RoleName  string `json:"roleName"`
	Links     []Link `json:"links"`
	ID        string `json:"id"`
	AccountID string `json:"accountId"`
	Arn       string `json:"arn"`
}

// CreateCloudAccountInput for function CreateCloudAccount
type CreateCloudAccountInput struct {
	TeamID      string
	CloudName   string
	RoleName    string
	ExternalID  string
	RoleArn     string
	Policy      string
	IsDraft     bool
	Email       string
	UserName    string
	Environment string
}

// CloudPayLoad ...
type CloudPayLoad struct {
	Name         string   `json:"name"`
	Arn          string   `json:"arn"`
	ScanEnabled  bool     `json:"scanEnabled"`
	ScanInterval string   `json:"scanInterval"`
	ScanRegion   string   `json:"scanRegion"`
	ExternalID   string   `json:"externalId"`
	IsDraft      bool     `json:"isDraft"`
	Provider     string   `json:"provider"`
	Email        string   `json:"email"`
	UserName     string   `json:"username"`
	Environment  []string `json:"environment"`
}

type sendCloudCreateRequestInput struct {
	cloudLink       Link
	externalID      string
	cloudName       string
	accessKeyID     string
	secretAccessKey string
	roleArn         string
	scanEnabled     bool
	scanInterval    string
	scanRegion      string
	isDraft         bool
	provider        string
	email           string
	username        string
	environment     []string
}

type defaultID struct {
	AccountID  string `json:"accountId"`
	ExternalID string `json:"externalId"`
	Domain     string `json:"domain"`
}

type RoleCreationInfo struct {
	AwsAccount string
	ExternalID string
	RoleName   string
	Policy     string
}

// GetCloudAccounts method for cloud command
func (c *Client) GetCloudAccounts(ctx context.Context, teamID string) ([]*CloudAccount, error) {
	// clouds := []*CloudAccount{}
	clouds := make([]*CloudAccount, 0)
	teams, err := c.GetTeams(ctx)

	if err != nil {
		return nil, err
	}

	for _, team := range teams {
		if team.ID == teamID {
			cloudLink, e := GetLinkByRef(team.Links, "cloudAccounts")

			if e != nil {
				return nil, NewError(e.Error())
			}

			err = c.Do(ctx, "GET", cloudLink.Href, nil, &clouds)
			if err != nil {
				return nil, NewError(err.Error())
			}
		}
	}

	if len(clouds) == 0 {
		return nil, NewError(fmt.Sprintf(content.ErrorNoCloudAccountsFound, teamID))
	}

	return clouds, nil
}

// GetCloudAccountByID method getting cloud account by user ID
func (c *Client) GetCloudAccountByID(ctx context.Context, teamID, cloudID string) (*CloudAccount, error) {
	cloudAccount := &CloudAccount{}

	cloudAccounts, err := c.GetCloudAccounts(ctx, teamID)

	if err != nil {
		return nil, NewError(err.Error())
	}

	for _, c := range cloudAccounts {
		if c.ID == cloudID {
			cloudAccount = c
			break
		}
	}

	if cloudAccount.ID == "" {
		return nil, NewError(fmt.Sprintf(content.ErrorNoCloudAccountWithIDFound, cloudID, teamID))
	}

	return cloudAccount, nil
}

func (c *Client) sendCloudCreateRequest(ctx context.Context, input *sendCloudCreateRequestInput) (*CloudAccount, error) {
	// Connect with webapp to add the new cloud account into team
	// Do not include space in cloudPayLoad!!! Otherwise the whitespace would be removed at some point and
	// the authentication would fail!!!
	cloudAccount := &CloudAccount{}
	cloudPayLoad := CloudPayLoad{
		Name:         input.cloudName,
		Arn:          input.roleArn,
		ScanEnabled:  input.scanEnabled,
		ScanInterval: input.scanInterval,
		ScanRegion:   input.scanRegion,
		ExternalID:   input.externalID,
		IsDraft:      input.isDraft,
		Provider:     input.provider,
		Email:        input.email,
		UserName:     input.username,
		Environment:  input.environment,
	}

	jsonStr, err := json.Marshal(cloudPayLoad)
	if err != nil {
		return nil, err
	}

	err = c.Do(ctx, "POST", input.cloudLink.Href, bytes.NewBuffer(jsonStr), &cloudAccount)
	if err != nil {
		return nil, err
	}
	return cloudAccount, nil
}

// GetRoleCreationInfo returns the configuration for creating a new role
func (c *Client) GetRoleCreationInfo(ctx context.Context, input *CreateCloudAccountInput) (*RoleCreationInfo, error) {
	teams, err := c.GetTeams(ctx)
	if err != nil {
		return nil, err
	}

	for _, team := range teams {
		if team.ID == input.TeamID {
			ref, err := GetLinkByRef(team.Links, "defaultid")

			if err != nil {
				return nil, err
			}

			defaultId := defaultID{}
			err = c.Do(ctx, "GET", ref.Href, nil, &defaultId)
			if err != nil {
				return nil, err
			}

			createNewRoleInfo := &RoleCreationInfo{
				RoleName:   input.RoleName,
				ExternalID: input.TeamID + "-" + defaultId.ExternalID,
				AwsAccount: defaultId.AccountID,
				Policy:     input.Policy,
			}

			return createNewRoleInfo, nil
		}
	}
	return nil, NewError("No team id match")
}

// CreateCloudAccount method to create a cloud object
func (c *Client) CreateCloudAccount(ctx context.Context, input *CreateCloudAccountInput) (*CloudAccount, error) {
	var cloudAccount = &CloudAccount{}
	teams, err := c.GetTeams(ctx)
	if err != nil {
		return nil, err
	}
	for _, team := range teams {
		if team.ID == input.TeamID {
			cloudLink, err := GetLinkByRef(team.Links, "cloudAccounts")
			if err != nil {
				return nil, NewError(err.Error())
			}

			if input.RoleArn == "" {
				return nil, NewError(content.ErrorMissingRoleInformation)
			}
			cloudCreateInput := &sendCloudCreateRequestInput{
				cloudLink:    cloudLink,
				externalID:   input.ExternalID,
				cloudName:    input.CloudName,
				roleArn:      input.RoleArn,
				scanEnabled:  true,
				scanInterval: "Daily",
				scanRegion:   "All",
				isDraft:      input.IsDraft,
				provider:     "AWS",
				email:        input.Email,
				username:     input.UserName,
			}
			if input.Environment != "" {
				cloudCreateInput.environment = []string{input.Environment}
			}
			cloudAccount, err = c.sendCloudCreateRequest(ctx, cloudCreateInput)
			if err != nil {
				return nil, err
			}

			break
		}
	}
	if cloudAccount.ID == "" {
		return nil, NewError(fmt.Sprintf(content.ErrorFailedToCreateCloudAccount, input.TeamID))
	}
	return cloudAccount, nil
}

// DeleteCloudAccountByID method to delete cloud object
func (c *Client) DeleteCloudAccountByID(ctx context.Context, teamID, cloudID string) error {
	cloudAccounts, err := c.GetCloudAccounts(ctx, teamID)
	cloudAccountFound := false
	if err != nil {
		return err
	}

	for _, cloudAccount := range cloudAccounts {
		if cloudAccount.ID == cloudID {
			cloudAccountFound = true
			cloudLink, err := GetLinkByRef(cloudAccount.Links, "self")
			if err != nil {
				return NewError(err.Error())
			}

			err = c.Do(ctx, "DELETE", cloudLink.Href, nil, nil)
			if err != nil {
				return NewError(err.Error())
			}
			break
		}
	}

	if !cloudAccountFound {
		return NewError(fmt.Sprintf(content.ErrorFailedToDeleteCloudAccount, cloudID, teamID))
	}

	return nil
}
