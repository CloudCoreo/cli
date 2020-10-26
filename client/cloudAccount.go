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
	"math/rand"
	"strings"

	"github.com/imdario/mergo"

	"github.com/CloudCoreo/cli/client/content"
)

// CloudAccount Information
type CloudAccount struct {
	RoleID    string `json:"roleId"`
	RoleName  string `json:"roleName"`
	ID        string `json:"_id"`
	AccountID string `json:"accountId"`
	CloudInfo
}

// CreateCloudAccountInput for function CreateCloudAccount
type CreateCloudAccountInput struct {
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
	CSPProjectID   string
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
	Environment         string   `json:"environment,omitempty"`
	KeyValue            string   `json:"key,omitempty"`
	ApplicationID       string   `json:"appId,omitempty"`
	DirectoryID         string   `json:"directoryId,omitempty"`
	SubscriptionID      string   `json:"subscriptionId,omitempty"`
	Tags                []string `json:"tags,omitempty"`
	IsValid             bool     `json:"isValid"`
	LastValidationCheck string   `json:"lastValidationCheck"`
	CSPProjectID        string   `json:"cspProjectId"`
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
func (c *Client) GetCloudAccounts(ctx context.Context) ([]*CloudAccount, error) {
	// clouds := []*CloudAccount{}
	clouds := make([]*CloudAccount, 0)

	err := c.Do(ctx, "GET", "cloudaccounts", nil, &clouds)
	if err != nil {
		return nil, NewError(err.Error())
	}
	for _, account := range clouds {
		if account.Provider == "Azure" {
			account.AccountID = account.SubscriptionID
		}
	}

	if len(clouds) == 0 {
		return nil, NewError(content.ErrorNoCloudAccountsFound)
	}

	return clouds, nil
}

// GetCloudAccountByID method getting cloud account by user ID
func (c *Client) GetCloudAccountByID(ctx context.Context, cloudID string) (*CloudAccount, error) {
	cloudAccount := &CloudAccount{}

	err := c.Do(ctx, "GET", fmt.Sprintf("cloudaccounts/%s", cloudID), nil, cloudAccount)
	if err != nil {
		return nil, err
	}

	if cloudAccount.ID == "" {
		return nil, NewError(fmt.Sprintf(content.ErrorNoCloudAccountWithIDFound, cloudID))
	}
	return cloudAccount, nil
}

func (c *Client) sendCloudCreateRequest(ctx context.Context, input *CloudInfo) (*CloudAccount, error) {
	// Connect with webapp to add the new cloud account into team
	// Do not include space in cloudPayLoad!!! Otherwise the whitespace would be removed at some point and
	// the authentication would fail!!!
	cloudAccount := CloudAccount{}

	jsonStr, err := json.Marshal(*input)
	if err != nil {
		return nil, err
	}

	err = c.Do(ctx, "POST", "cloudaccounts", bytes.NewBuffer(jsonStr), &cloudAccount)
	if err != nil {
		return nil, err
	}
	return &cloudAccount, nil
}

// GetRoleCreationInfo returns the configuration for creating a new role
func (c *Client) GetRoleCreationInfo(ctx context.Context, input *CreateCloudAccountInput) (*RoleCreationInfo, error) {

	id := defaultID{}
	err := c.Do(ctx, "GET", ".well-known/vss-configuration", nil, &id)
	if err != nil {
		return nil, err
	}

	createNewRoleInfo := &RoleCreationInfo{
		RoleName: input.RoleName,
		//Need to find out the right way to create external id.
		ExternalID: c.genRandomString(10) + id.ExternalID,
		AwsAccount: id.AccountID,
		Policy:     input.Policy,
	}

	return createNewRoleInfo, nil
}

func (c *Client) genRandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// CreateCloudAccount method to create a cloud object
func (c *Client) CreateCloudAccount(ctx context.Context, input *CreateCloudAccountInput) (*CloudAccount, error) {
	var cloudAccount *CloudAccount

	if input.Provider == "AWS" && input.RoleArn == "" {
		return nil, NewError(content.ErrorMissingRoleInformation)
	}
	cloudCreateInput := CloudInfo{
		ExternalID:     input.ExternalID,
		Name:           input.CloudName,
		Arn:            input.RoleArn,
		ScanEnabled:    true,
		ScanRegion:     "All",
		IsDraft:        input.IsDraft,
		Provider:       input.Provider,
		Email:          input.Email,
		UserName:       input.UserName,
		KeyValue:       input.KeyValue,
		ApplicationID:  input.ApplicationID,
		DirectoryID:    input.DirectoryID,
		SubscriptionID: input.SubscriptionID,
		Environment:    input.Environment,
		CSPProjectID:   input.CSPProjectID,
	}
	if input.Tags != "" {
		cloudCreateInput.Tags = strings.Split(input.Tags, "|")
	}
	if input.Provider == "AWS" {
		cloudCreateInput.ScanInterval = "Weekly"
	} else if input.Provider == "Azure" {
		cloudCreateInput.ScanInterval = "Daily"
	} else {
		return nil, NewError("Unsupported CloudAccount type")
	}
	cloudAccount, err := c.sendCloudCreateRequest(ctx, &cloudCreateInput)
	if err != nil {
		return nil, err
	}

	if cloudAccount.ID == "" {
		return nil, NewError(content.ErrorFailedToCreateCloudAccount)
	}
	if cloudAccount.Provider == "Azure" {
		cloudAccount.AccountID = cloudAccount.SubscriptionID
	}
	return cloudAccount, nil
}

// DeleteCloudAccountByID method to delete cloud object
func (c *Client) DeleteCloudAccountByID(ctx context.Context, cloudID string) error {
	err := c.Do(ctx, "DELETE", fmt.Sprintf("cloudaccounts/%s", cloudID), nil, nil)
	return err
}

//ReValidateRole checks role validation and re-validate it
func (c *Client) ReValidateRole(ctx context.Context, cloudID string) (*RoleReValidationResult, error) {
	result := new(RoleReValidationResult)
	err := c.Do(ctx, "GET", fmt.Sprintf("cloudaccounts/%s/re-validate", cloudID), nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//UpdateCloudAccount updates cloud account
func (c *Client) UpdateCloudAccount(ctx context.Context, input *UpdateCloudAccountInput) (*CloudAccount, error) {
	result := new(CloudAccount)
	account, err := c.GetCloudAccountByID(ctx, input.CloudID)
	if err != nil {
		return nil, err
	}

	updateInfo, err := input.mergeAndGetJSON(account)
	if err != nil {
		return nil, err
	}
	err = c.Do(ctx, "PUT", fmt.Sprintf("cloudaccounts/%s", input.CloudID), bytes.NewBuffer(updateInfo), result)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (t *UpdateCloudAccountInput) mergeAndGetJSON(account *CloudAccount) ([]byte, error) {
	updateInfo := t.toCloudInfo()
	err := mergo.Merge(updateInfo, *(account.toCloudInfo()))
	if err != nil {
		return nil, err
	}
	// mergo package will override false to true
	updateInfo.IsDraft = t.IsDraft
	updateInfo.Environment = t.Environment
	if t.Tags != "" {
		updateInfo.Tags = strings.Split(t.Tags, "|")
	}
	jsonStr, err := json.Marshal(updateInfo)
	if err != nil {
		return nil, err
	}

	return jsonStr, nil
}

func (t *UpdateCloudAccountInput) toCloudInfo() *CloudInfo {
	cloudInfo := &CloudInfo{
		Name:           t.CloudName,
		Arn:            t.RoleArn,
		ScanEnabled:    t.ScanEnabled,
		ScanRegion:     "All",
		ExternalID:     t.ExternalID,
		IsDraft:        t.IsDraft,
		Email:          t.Email,
		UserName:       t.UserName,
		SubscriptionID: t.SubscriptionID,
		KeyValue:       t.KeyValue,
		ApplicationID:  t.ApplicationID,
		DirectoryID:    t.DirectoryID,
		CSPProjectID:   t.CSPProjectID,
	}

	return cloudInfo
}

func (t *CloudAccount) toCloudInfo() *CloudInfo {
	cloudInfo := &CloudInfo{
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
		CSPProjectID:   t.CSPProjectID,
	}
	return cloudInfo
}
