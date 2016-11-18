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
	"context"
)

// Team struct for api payload
type Team struct {
	TeamName        string      `json:"teamName"`
	OwnerID         string      `json:"ownerId"`
	TeamIcon        string      `json:"teamIcon"`
	TeamDescription interface{} `json:"teamDescription"`
	Default         bool        `json:"default"`
	Links           []Link      `json:"links"`
	ID              string      `json:"id"`
}

// GetTeams method to get Teams info array object
func (c *Client) GetTeams(ctx context.Context) ([]*Team, error) {
	t := []*Team{}
	u, err := c.GetUser(ctx)

	if err != nil {
		return t, err
	}

	teamLink, err := GetLinkByRef(u.Links, "teams")
	if err != nil {
		return t, err
	}

	err = c.Do(ctx, "GET", teamLink.Href, nil, &t)
	if err != nil {
		return t, err
	}

	return t, nil
}

// GetTeamByID method to get Team info object by team ID
func (c *Client) GetTeamByID(ctx context.Context, teamID string) (*Team, error) {
	team := &Team{}
	teams, err := c.GetTeams(ctx)
	if err != nil {
		return team, err
	}

	for _, t := range teams {
		if t.ID == teamID {
			team = t
			break
		}
	}
	return team, nil
}
