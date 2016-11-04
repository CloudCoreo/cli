package client

import (
	"bytes"
	"fmt"

	"golang.org/x/net/context"
)

// Cloud struct for api payload
type CloudAccount struct {
	TeamID string `json:"teamId"`
	Name string `json:"name"`
	RoleID string `json:"roleId"`
	RoleName string `json:"roleName"`
	Links []Link `json:"links"`
	ID string `json:"id"`
}

// GetCloudAccounts method for cloud command
func (c *Client) GetCloudAccounts(ctx context.Context, teamID string) ([]CloudAccount, error) {
	clouds := []CloudAccount{}
	teams, err := c.GetTeams(ctx)

	if err != nil {
		return clouds, err
	}

	for _,team := range teams {
		if team.ID == teamID {
			cloudLink := GetLinkByRef(team.Links, "cloudAccounts")

			err := c.Do(ctx, "GET", cloudLink.Href, nil, &clouds)
			if err != nil {
				return clouds, err
			}
		}
	}

	return clouds, nil
}

// GetCloudAccount method for cloud command
func (c *Client) GetCloudAccountByID(ctx context.Context, teamID, cloudID string) (CloudAccount, error) {
	cloudAccount := CloudAccount{}

	cloudAccounts, err := c.GetCloudAccounts(ctx, teamID)

	if err != nil {
		return cloudAccount, err
	}

	for _, c := range cloudAccounts {
		if c.ID == cloudID {
			cloudAccount = c
			break
		}
	}

	return cloudAccount, nil
}

// CreateCloudAccount method to create a cloud object
func (c *Client) CreateCloudAccount(ctx context.Context, teamID, accessKeyID, secretAccessKey, cloudName string) (CloudAccount, error) {
	cloudAccount := CloudAccount{}
	teams, err := c.GetTeams(ctx)

	if err != nil {
		return cloudAccount, err
	}

	for _,team := range teams {


		if team.ID == teamID {
			cloudPlayLoad := fmt.Sprintf(`{"name":"%s","accessKeyId":"%s","secretAccessKey":"%s","teamId":"%s"}`, cloudName, accessKeyID, secretAccessKey, teamID)
			var jsonStr = []byte(cloudPlayLoad)

			cloudLink := GetLinkByRef(team.Links, "cloudAccounts")

			err := c.Do(ctx, "POST", cloudLink.Href, bytes.NewBuffer(jsonStr), &cloudAccount)
			if err != nil {
				return cloudAccount, err
			}
			break
		}
	}

	return cloudAccount, nil
}

// DeleteCloudAccountByID method to delete cloud object
func (c *Client) DeleteCloudAccountByID(ctx context.Context, teamID, cloudID string) error {
	cloudAccounts, err := c.GetCloudAccounts(ctx, teamID)

	if err != nil {
		return err
	}

	for _, cloudAccount := range cloudAccounts {
		if cloudAccount.ID == cloudID {
			cloudLink := GetLinkByRef(cloudAccount.Links, "self")
			err := c.Do(ctx, "DELETE", cloudLink.Href, nil, nil)
			if err != nil {
				return err
			}
			break
		}
	}

	return nil
}