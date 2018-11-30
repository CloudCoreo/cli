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
	TeamID   string `json:"teamId"`
	Name     string `json:"name"`
	RoleID   string `json:"roleId"`
	RoleName string `json:"roleName"`
	Links    []Link `json:"links"`
	ID       string `json:"id"`
}

// CreateCloudAccountInput for function CreateCloudAccount
type CreateCloudAccountInput struct {
	TeamID          string
	AccessKeyID     string
	SecretAccessKey string
	CloudName       string
	RoleName        string
	ExternalID      string
	RoleArn         string
	Policy          string
}

// CloudPayLoad ...
type CloudPayLoad struct {
	Name         string `json:"name"`
	Arn          string `json:"arn"`
	ScanEnabled  bool   `json:"scanEnabled"`
	ScanInterval string `json:"scanInterval"`
	ScanRegion   string `json:"scanRegion"`
	ExternalID   string `json:"externalId"`
	IsDraft      bool   `json:"isDraft"`
	Provider     string `json:"provider"`
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

/*
func (c *Client) createAssumeRolePolicyDocument(awsAccount string, externalID string) string {
	return `{
	"Version": "2012-10-17",
	"Statement": [
		{
			"Effect": "Allow",
			"Principal": {
				"AWS": "arn:aws:iam::` + awsAccount + `:root"
			},
			"Action": "sts:AssumeRole",
			"Condition": {
				"StringEquals": {
					"sts:ExternalId": "` + externalID + `"
				}
			}
		}
	]
}`
}
*/

// GetCloudAccounts method for cloud command
func (c *Client) GetCloudAccounts(ctx context.Context, teamID string) ([]*CloudAccount, error) {
	clouds := []*CloudAccount{}
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

/*
func (c *Client) createNewRole(ctx context.Context, input *RoleCreationInfo, sess *session.Session) (*string, error) {
	svc := iam.New(sess)

	// Create a new session for iam
	result, err := c.createNewAwsRole(input.AwsAccount, input.ExternalID, input.RoleName, svc)
	if err != nil {
		return nil, err
	}
	roleArn := result.Role.Arn
	c.attachRolePolicy(svc, input.Policy, input.RoleName)
	return roleArn, nil
}
*/

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

/*
func (c *Client) checkRolePolicy(roleName, policy string, sess *session.Session) (bool, error) {
	svc := iam.New(sess)
	input := &iam.ListAttachedRolePoliciesInput{}
	input.SetRoleName(roleName)
	res, err := svc.ListAttachedRolePolicies(input)
	if err != nil {
		return false, err
	}
	for i := range res.AttachedPolicies {
		if *res.AttachedPolicies[i].PolicyName == policy {
			return true, nil
		}
	}
	return false, nil
}
*/

/*
func (c *Client) newSession(input *CreateCloudAccountInput) (*session.Session, error) {
	var sess *session.Session
	var err error
	if input.AwsProfile != "" {
		sess, err = session.NewSession(&aws.Config{Credentials: credentials.NewSharedCredentials(input.AwsProfilePath, input.AwsProfile)})
		if err != nil {
			return nil, err
		}
	} else {
		sess, err = session.NewSession()
		if err != nil {
			return nil, err
		}
	}
	return sess, nil
}
*/

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
			c.Do(ctx, "GET", ref.Href, nil, &defaultId)

			createNewRoleInfo := &RoleCreationInfo{
				RoleName:   input.RoleName,
				ExternalID: input.TeamID + "-" + defaultId.ExternalID,
				AwsAccount: defaultId.AccountID,
				Policy:     input.Policy,
			}

			return createNewRoleInfo, nil
		}
	}
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
			/*
				var roleArn *string
				var externalID string


				// Check aws credentials
				sess, err := c.newSession(input)
				if err != nil {
					return nil, err
				}
			*/

			cloudLink, err := GetLinkByRef(team.Links, "cloudAccounts")
			if err != nil {
				return nil, NewError(err.Error())
			}
			/*
				// Check whether the rolearn and externalID are provided
				if input.RoleArn != "" && input.ExternalID != "" {
					roleArn = &input.RoleArn
					externalID = input.ExternalID

					// Check whether SecurityAudit is Enabled
					// roleNames := strings.Split(input.RoleArn, `/`)
					// boolean, err := c.checkRolePolicy(roleNames[len(roleNames)-1], "SecurityAudit", sess)
					// if err != nil {
					//	 return nil, err
					// }
					// if !boolean {
					// 	 return nil, NewError("SecurityAudit is not enable in role " + *roleArn)
					// }
				} else if input.RoleName != "" {
					// Create a new role for the user
					ref, err := GetLinkByRef(team.Links, "defaultid")
					defaultid := defaultID{}
					c.Do(ctx, "GET", ref.Href, nil, &defaultid)
					externalID = input.TeamID + "-" + defaultid.ExternalID

					createNewRoleInput := &RoleCreationInfo{
						RoleName:   input.RoleName,
						ExternalID: externalID,
						AwsAccount: defaultid.AccountID,
						Policy:     input.Policy,
					}
					roleArn, err = c.createNewRole(ctx, createNewRoleInput, sess)

					if err != nil {
						return nil, NewError(err.Error())
					}
					time.Sleep(10 * time.Second)
				} else {
					return nil, NewError(content.ErrorMissingRoleInformation)
				}
			*/

			cloudCreateInput := &sendCloudCreateRequestInput{
				cloudLink:       cloudLink,
				externalID:      input.ExternalID,
				cloudName:       input.CloudName,
				accessKeyID:     input.AccessKeyID,
				secretAccessKey: input.SecretAccessKey,
				roleArn:         input.RoleArn,
				scanEnabled:     true,
				scanInterval:    "Daily",
				scanRegion:      "All",
				isDraft:         false,
				provider:        "AWS",
			}
			cloudAccount, err = c.sendCloudCreateRequest(ctx, cloudCreateInput)
			/*
				if err != nil {
					if input.RoleName != "" {
						c.deleteRole(ctx, sess, input.RoleName, input.Policy)
					}
					return nil, err
				}
			*/
			break
		}
	}
	if cloudAccount.ID == "" {
		return nil, NewError(fmt.Sprintf(content.ErrorFailedToCreateCloudAccount, input.TeamID))
	}
	return cloudAccount, nil
}

/*
func (c *Client) createNewAwsRole(awsAccount, externalID, roleName string, svc *iam.IAM) (*iam.CreateRoleOutput, error) {
	input := &iam.CreateRoleInput{
		AssumeRolePolicyDocument: aws.String(c.createAssumeRolePolicyDocument(awsAccount, externalID)),
		Path:     aws.String("/"),
		RoleName: aws.String(roleName),
	}
	result, err := svc.CreateRole(input)
	if err != nil {
		return nil, err
	}
	return result, nil
}
*/

/*
func (c *Client) deleteRole(ctx context.Context, sess *session.Session, roleName, policyArn string) {
	svc := iam.New(sess)
	detachPolicyInput := &iam.DetachRolePolicyInput{
		PolicyArn: aws.String(policyArn),
		RoleName:  aws.String(roleName),
	}
	_, err := svc.DetachRolePolicy(detachPolicyInput)
	if err != nil {
		fmt.Println("Detach role policy " + policyArn + "for " + roleName + " failed, " + err.Error())
		return
	}

	deleteRoleInput := &iam.DeleteRoleInput{
		RoleName: aws.String(roleName),
	}
	_, err = svc.DeleteRole(deleteRoleInput)
	if err != nil {
		fmt.Println("Delete role " + roleName + " failed, " + err.Error())
	}
}
*/

/*
func (c *Client) attachRolePolicy(svc *iam.IAM, policyArn, roleName string) (*iam.AttachRolePolicyOutput, error) {
	input := &iam.AttachRolePolicyInput{
		PolicyArn: aws.String(policyArn),
		RoleName:  aws.String(roleName),
	}

	result, err := svc.AttachRolePolicy(input)
	return result, err
}
*/

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
