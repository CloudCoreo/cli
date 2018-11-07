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
	"fmt"

	"github.com/CloudCoreo/cli/pkg/command"

	"bytes"

	"github.com/CloudCoreo/cli/client/content"
)

// GetTeams method to get Teams info array object
func (c *Client) GetTeams(ctx context.Context) ([]*command.Team, error) {
	u, err := c.GetUser(ctx)

	if err != nil {
		return nil, err
	}

	teamLink, err := GetLinkByRef(u.Links, "teams")
	if err != nil {
		return nil, err
	}

	t := []*command.Team{}
	err = c.Do(ctx, "GET", teamLink.Href, nil, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// GetTeamByID method to get Team info object by team ID
func (c *Client) GetTeamByID(ctx context.Context, teamID string) (*command.Team, error) {
	teams, err := c.GetTeams(ctx)
	if err != nil {
		return nil, err
	}

	team := &command.Team{}
	for _, t := range teams {
		if t.ID == teamID {
			team = t
			break
		}
	}

	if team.ID == "" {
		return nil, NewError(fmt.Sprintf(content.ErrorNoTeamWithIDFound, teamID))
	}

	return team, nil
}

// CreateTeam method to create a new team
func (c *Client) CreateTeam(ctx context.Context, teamName, teamDescription string) (*command.Team, error) {

	u, err := c.GetUser(ctx)

	if err != nil {
		return nil, err
	}

	teamLink, err := GetLinkByRef(u.Links, "teams")
	if err != nil {
		return nil, err
	}

	teamPayLoad := fmt.Sprintf(`{"teamName":"%s","teamDescription":"%s"}`, teamName, teamDescription)

	var jsonStr = []byte(teamPayLoad)

	team := &command.Team{}
	err = c.Do(ctx, "POST", teamLink.Href, bytes.NewBuffer(jsonStr), &team)
	if err != nil {
		return nil, err
	}

	return team, nil
}
