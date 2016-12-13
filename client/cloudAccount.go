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
	"fmt"

	"github.com/CloudCoreo/cli/client/content"
)

// CloudAccount struct for api payload
type CloudAccount struct {
	TeamID   string `json:"teamId"`
	Name     string `json:"name"`
	RoleID   string `json:"roleId"`
	RoleName string `json:"roleName"`
	Links    []Link `json:"links"`
	ID       string `json:"id"`
}

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

// CreateCloudAccount method to create a cloud object
func (c *Client) CreateCloudAccount(ctx context.Context, teamID, accessKeyID, secretAccessKey, cloudName string) (*CloudAccount, error) {
	cloudAccount := &CloudAccount{}
	teams, err := c.GetTeams(ctx)

	if err != nil {
		return cloudAccount, err
	}

	for _, team := range teams {

		if team.ID == teamID {
			cloudPayLoad := fmt.Sprintf(`{"name":"%s","accessKeyId":"%s","secretAccessKey":"%s","teamId":"%s"}`, cloudName, accessKeyID, secretAccessKey, teamID)
			var jsonStr = []byte(cloudPayLoad)

			cloudLink, err := GetLinkByRef(team.Links, "cloudAccounts")
			if err != nil {
				return nil, NewError(err.Error())
			}

			err = c.Do(ctx, "POST", cloudLink.Href, bytes.NewBuffer(jsonStr), &cloudAccount)
			if err != nil {
				return nil, NewError(err.Error())
			}
			break
		}
	}

	if cloudAccount.ID == "" {
		return nil, NewError(fmt.Sprintf(content.ErrorFailedToCreateCloudAccount, teamID))
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
