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

// GitKey struct for api payload
type GitKey struct {
	TeamID            string `json:"teamId"`
	Name              string `json:"name"`
	Sha256Fingerprint string `json:"sha256fingerprint"`
	Md5Fingerprint    string `json:"md5fingerprint"`
	Links             []Link `json:"links"`
	ID                string `json:"id"`
}

// GetGitKeys method for gitKey command
func (c *Client) GetGitKeys(ctx context.Context, teamID string) ([]*GitKey, error) {
	gitKeys := []*GitKey{}
	teams, err := c.GetTeams(ctx)

	if err != nil {
		return nil, NewError(err.Error())
	}

	for _, team := range teams {
		if team.ID == teamID {
			gitKeyLink, e := GetLinkByRef(team.Links, "gitKeys")

			if e != nil {
				return nil, NewError(e.Error())
			}

			e = c.Do(ctx, "GET", gitKeyLink.Href, nil, &gitKeys)
			if e != nil {
				return nil, NewError(e.Error())
			}
		}
	}

	if len(gitKeys) == 0 {
		return nil, NewError(fmt.Sprintf(content.ErrorNoGitKeysFound, teamID))
	}

	return gitKeys, nil
}

// GetGitKeyByID method for gitKey command
func (c *Client) GetGitKeyByID(ctx context.Context, teamID, gitKeyID string) (*GitKey, error) {
	gitKey := &GitKey{}

	gitKeys, err := c.GetGitKeys(ctx, teamID)

	if err != nil {
		return nil, NewError(err.Error())
	}

	for _, g := range gitKeys {
		if g.ID == gitKeyID {
			gitKey = g
			break
		}
	}

	if gitKey.ID == "" {
		return nil, NewError(fmt.Sprintf(content.ErrorNoGitKeyWithIDFound, gitKeyID, teamID))
	}

	return gitKey, nil
}

// CreateGitKey method to create a gitKey object
func (c *Client) CreateGitKey(ctx context.Context, teamID, keyMaterial, name string) (*GitKey, error) {
	gitKey := &GitKey{}
	teams, err := c.GetTeams(ctx)

	if err != nil {
		return nil, NewError(err.Error())
	}

	for _, team := range teams {
		if team.ID == teamID {
			gitKeyPayLoad := fmt.Sprintf(`{"keyMaterial":"%s","name":"%s","teamId":"%s"}`, keyMaterial, name, teamID)
			var jsonStr = []byte(gitKeyPayLoad)
			gitKeyLink, err := GetLinkByRef(team.Links, "gitKeys")
			if err != nil {
				return nil, NewError(err.Error())
			}

			err = c.Do(ctx, "POST", gitKeyLink.Href, bytes.NewBuffer(jsonStr), &gitKey)
			if err != nil {
				return nil, NewError(err.Error())
			}
			break
		}
	}

	if gitKey.ID == "" {
		return nil, NewError(fmt.Sprintf(content.ErrorFailedToCreateGitKey, teamID))
	}

	return gitKey, nil
}

// DeleteGitKeyByID method to delete gitKey object
func (c *Client) DeleteGitKeyByID(ctx context.Context, teamID, gitKeyID string) error {
	gitKeys, err := c.GetGitKeys(ctx, teamID)

	if err != nil {
		return err
	}

	gitKeyFound := false

	for _, gitKey := range gitKeys {
		if gitKey.ID == gitKeyID {
			gitKeyFound = true
			gitKeyLink, err := GetLinkByRef(gitKey.Links, "self")

			if err != nil {
				return NewError(err.Error())
			}

			err = c.Do(ctx, "DELETE", gitKeyLink.Href, nil, nil)
			if err != nil {
				return NewError(err.Error())
			}
			break
		}
	}

	if !gitKeyFound {
		return NewError(fmt.Sprintf(content.ErrorFailedToDeleteGitKey, gitKeyID, teamID))
	}

	return nil
}
