package client

import (
	"bytes"
	"fmt"

	"golang.org/x/net/context"
)

// GitKey struct for api payload
type GitKey struct {
	TeamID string `json:"teamId"`
	Name string `json:"name"`
	Sha256Fingerprint string `json:"sha256fingerprint"`
	Md5Fingerprint string `json:"md5fingerprint"`
	Links []Link `json:"links"`
	ID string `json:"id"`
}

// GetGitKeys method for gitKey command
func (c *Client) GetGitKeys(ctx context.Context, teamID string) ([]GitKey, error) {
	gitKeys := []GitKey{}
	teams, err := c.GetTeams(ctx)

	if err != nil {
		return gitKeys, err
	}

	for _,team := range teams {
		if team.ID == teamID {
			gitKeyLink, err := GetLinkByRef(team.Links, "gitKeys")

			if err != nil {
				return gitKeys, err
			}

			err = c.Do(ctx, "GET", gitKeyLink.Href, nil, &gitKeys)
			if err != nil {
				return gitKeys, err
			}
		}
	}

	return gitKeys, nil
}

// GetGitKeyByID method for gitKey command
func (c *Client) GetGitKeyByID(ctx context.Context, teamID, gitKeyID string) (GitKey, error) {
	gitKey := GitKey{}

	gitKeys, err := c.GetGitKeys(ctx, teamID)

	if err != nil {
		return gitKey, err
	}

	for _, g := range gitKeys {
		if g.ID == gitKeyID {
			gitKey = g
			break
		}
	}

	return gitKey, nil
}

// CreateGitKey method to create a gitKey object
func (c *Client) CreateGitKey(ctx context.Context, teamID, keyMaterial, name string) (GitKey, error) {
	gitKey := GitKey{}
	teams, err := c.GetTeams(ctx)

	if err != nil {
		return gitKey, err
	}

	for _,team := range teams {
		if team.ID == teamID {
			gitKeyPlayLoad := fmt.Sprintf(`{"keyMaterial":"%s","name":"%s","teamId":"%s"}`, keyMaterial, name, teamID)
			var jsonStr = []byte(gitKeyPlayLoad)
			gitKeyLink, err := GetLinkByRef(team.Links, "gitKeys")
			if err != nil {
				return gitKey, err
			}

			err = c.Do(ctx, "POST", gitKeyLink.Href, bytes.NewBuffer(jsonStr), &gitKey)
			if err != nil {
				return gitKey, err
			}
			break
		}
	}

	return gitKey, nil
}

// DeleteGitKeyByID method to delete gitKey object
func (c *Client) DeleteGitKeyByID(ctx context.Context, teamID, gitKeyID string) error {
	gitKeys, err := c.GetGitKeys(ctx, teamID)

	if err != nil {
		return err
	}

	for _, gitKey := range gitKeys {
		if gitKey.ID == gitKeyID {
			gitKeyLink, err := GetLinkByRef(gitKey.Links, "self")

			if err != nil {
				return err
			}

			err = c.Do(ctx, "DELETE", gitKeyLink.Href, nil, nil)
			if err != nil {
				return err
			}
			break
		}
	}

	return nil
}