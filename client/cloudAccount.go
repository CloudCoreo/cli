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
	"strings"

	"github.com/imdario/mergo"

	"github.com/CloudCoreo/cli/client/content"
)

// CloudAccount Information
type CloudAccount struct {
	RoleID    string `json:"roleId"`
	RoleName  string `json:"roleName"`
	Links     []Link `json:"links"`
	ID        string `json:"id"`
	AccountID string `json:"accountId"`
	CloudPayLoad
}

// CreateCloudAccountInput for function CreateCloudAccount
type CreateCloudAccountInput struct {
	TeamID         string
	CloudName      string
	RoleName       string
	ExternalID     string
	RoleArn        string
	Policy         string
	IsDraft        bool
	Email          string
	UserName       string
	Environment    string
	ScanEnabled    bool
	Provider       string
	KeyValue       string
	ApplicationID  string
	DirectoryID    string
	SubscriptionID string
	Tags           string
}

//CloudInfo listed all info of cloud accounts
type CloudInfo struct {
	Name                string   `json:"name,omitempty"`
	Arn                 string   `json:"arn,omitempty"`
	ScanEnabled         bool     `json:"scanEnabled"`
	ScanInterval        string   `json:"scanInterval"`
	ScanRegion          string   `json:"scanRegion"`
	ExternalID          string   `json:"externalId,omitempty"`
	IsDraft             bool     `json:"isDraft"`
	Provider            string   `json:"provider"`
	Email               string   `json:"email,omitempty"`
	UserName            string   `json:"username,omitempty"`
	Environment         []string `json:"environment,omitempty"`
	KeyValue            string   `json:"key,omitempty"`
	ApplicationID       string   `json:"appId,omitempty"`
	DirectoryID         string   `json:"directoryId,omitempty"`
	SubscriptionID      string   `json:"subscriptionId,omitempty"`
	Tags                []string `json:"tags,omitempty"`
	IsValid             bool     `json:"isValid"`
	LastValidationCheck string   `json:"lastValidationCheck"`
}

// CloudPayLoad ...
type CloudPayLoad struct {
	CloudInfo
	TeamID string `json:"teamId,omitempty"`
}

type sendCloudCreateRequestInput struct {
	cloudLink Link
	CloudInfo
}

type defaultID struct {
	AccountID  string `json:"accountId"`
	ExternalID string `json:"externalId"`
	Domain     string `json:"domain"`
}

//RoleCreationInfo contains the info required for role creation
type RoleCreationInfo struct {
	AwsAccount string
	ExternalID string
	RoleName   string
	Policy     string
}

//RoleReValidationResult is the result for role re-validation
type RoleReValidationResult struct {
	Message string `json:"message"`
	IsValid bool   `json:"isValid"`
}

//UpdateCloudAccountInput is the info needed for update cloud account
type UpdateCloudAccountInput struct {
	CreateCloudAccountInput
	CloudID string
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
		CloudInfo: input.CloudInfo,
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

			id := defaultID{}
			err = c.Do(ctx, "GET", ref.Href, nil, &id)
			if err != nil {
				return nil, err
			}

			createNewRoleInfo := &RoleCreationInfo{
				RoleName:   input.RoleName,
				ExternalID: input.TeamID + "-" + id.ExternalID,
				AwsAccount: id.AccountID,
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

			if input.Provider == "AWS" && input.RoleArn == "" {
				return nil, NewError(content.ErrorMissingRoleInformation)
			}
			cloudCreateInput := &sendCloudCreateRequestInput{
				cloudLink: cloudLink,
				CloudInfo: CloudInfo{
					ExternalID:     input.ExternalID,
					Name:           input.CloudName,
					Arn:            input.RoleArn,
					ScanEnabled:    true,
					ScanInterval:   "Daily",
					ScanRegion:     "All",
					IsDraft:        input.IsDraft,
					Provider:       input.Provider,
					Email:          input.Email,
					UserName:       input.UserName,
					KeyValue:       input.KeyValue,
					ApplicationID:  input.ApplicationID,
					DirectoryID:    input.DirectoryID,
					SubscriptionID: input.SubscriptionID,
				},
			}
			if input.Environment != "" {
				cloudCreateInput.Environment = []string{input.Environment}
			}
			if input.Tags != "" {
				cloudCreateInput.Tags = strings.Split(input.Tags, "|")
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

//ReValidateRole checks role validation and re-validate it
func (c *Client) ReValidateRole(ctx context.Context, teamID, cloudID string) (*RoleReValidationResult, error) {
	result := new(RoleReValidationResult)

	accounts, err := c.GetCloudAccountByID(ctx, teamID, cloudID)
	if err != nil {
		return nil, err
	}

	link, err := GetLinkByRef(accounts.Links, "test")
	if err != nil {
		return nil, err
	}

	err = c.Do(ctx, "GET", link.Href, nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//UpdateCloudAccount updates cloud account
func (c *Client) UpdateCloudAccount(ctx context.Context, input *UpdateCloudAccountInput) (*CloudAccount, error) {
	result := new(CloudAccount)
	account, err := c.GetCloudAccountByID(ctx, input.TeamID, input.CloudID)
	if err != nil {
		return nil, err
	}

	link, err := GetLinkByRef(account.Links, "update")
	if err != nil {
		return nil, err
	}

	updateInfo, err := input.mergeAndGetJSON(account)
	if err != nil {
		return nil, err
	}
	err = c.Do(ctx, "POST", link.Href, bytes.NewBuffer(updateInfo), result)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (t *UpdateCloudAccountInput) mergeAndGetJSON(account *CloudAccount) ([]byte, error) {
	updateInfo := t.toCloudPayLoad()
	err := mergo.Merge(updateInfo, *(account.toCloudPayLoad()))
	if err != nil {
		return nil, err
	}
	// mergo package will override false to true
	updateInfo.IsDraft = t.IsDraft
	// Add teamID to pass webapp check for cloud account creation
	updateInfo.TeamID = t.TeamID
	if t.Environment != "" {
		updateInfo.Environment = []string{t.Environment}
	}
	if t.Tags != "" {
		updateInfo.Tags = strings.Split(t.Tags, "|")
	}
	jsonStr, err := json.Marshal(updateInfo)
	if err != nil {
		return nil, err
	}

	return jsonStr, nil
}

func (t *UpdateCloudAccountInput) toCloudPayLoad() *CloudPayLoad {
	cloudPayLoad := &CloudPayLoad{
		CloudInfo: CloudInfo{
			Name:           t.CloudName,
			Arn:            t.RoleArn,
			ScanEnabled:    t.ScanEnabled,
			ScanInterval:   "Daily",
			ScanRegion:     "All",
			ExternalID:     t.ExternalID,
			IsDraft:        t.IsDraft,
			Email:          t.Email,
			UserName:       t.UserName,
			SubscriptionID: t.SubscriptionID,
			KeyValue:       t.KeyValue,
			ApplicationID:  t.ApplicationID,
			DirectoryID:    t.DirectoryID,
		},
	}
	return cloudPayLoad
}

func (t *CloudAccount) toCloudPayLoad() *CloudPayLoad {
	cloudPayLoad := &CloudPayLoad{
		CloudInfo: CloudInfo{
			Name:           t.Name,
			Arn:            t.Arn,
			ScanEnabled:    t.ScanEnabled,
			ScanInterval:   t.ScanInterval,
			ScanRegion:     t.ScanRegion,
			ExternalID:     t.ExternalID,
			IsDraft:        t.IsDraft,
			Provider:       t.Provider,
			Email:          t.Email,
			UserName:       t.UserName,
			Environment:    t.Environment,
			SubscriptionID: t.SubscriptionID,
			KeyValue:       t.KeyValue,
			ApplicationID:  t.ApplicationID,
			DirectoryID:    t.DirectoryID,
			Tags:           t.Tags,
		},
	}
	return cloudPayLoad
}
