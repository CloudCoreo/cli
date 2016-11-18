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
	"time"
)

// Composite struct for api payload
type Composite struct {
	Name               string    `json:"name"`
	GitURL             string    `json:"gitUrl"`
	HasCustomDashboard bool      `json:"hasCustomDashboard"`
	CreatedAt          time.Time `json:"createdAt"`
	GitKeyID           string    `json:"gitKeyId"`
	TeamID             string    `json:"teamId"`
	ID                 string    `json:"id"`
	Links              []Link    `json:"links"`
}

// GetComposites method to get composite info array object
func (c *Client) GetComposites(ctx context.Context, teamID string) ([]*Composite, error) {
	composites := []*Composite{}
	team, err := c.GetTeamByID(ctx, teamID)

	if err != nil {
		return nil, NewError(err.Error())
	}

	compsoitesLink, err := GetLinkByRef(team.Links, "composites")
	if err != nil {
		return nil, NewError(err.Error())
	}

	err = c.Do(ctx, "GET", compsoitesLink.Href, nil, &composites)
	if err != nil {
		return nil, NewError(err.Error())
	}

	if len(composites) == 0 {
		if err != nil {
			return nil, NewError(fmt.Sprintf("No composites found under team ID %s.", teamID))
		}
	}

	return composites, nil
}

// GetCompositeByID method to get composite info object
func (c *Client) GetCompositeByID(ctx context.Context, teamID, compositeID string) (*Composite, error) {
	composite := &Composite{}
	composites, err := c.GetComposites(ctx, teamID)
	if err != nil {
		return nil, NewError(err.Error())
	}

	for _, comp := range composites {
		if comp.ID == compositeID {
			composite = comp
			break
		}
	}

	if composite.ID == "" {
		return nil, NewError(fmt.Sprintf("No composite with ID %s found under team ID %s.", compositeID, teamID))
	}

	return composite, nil
}

// CreateComposite method to create a composite object
func (c *Client) CreateComposite(ctx context.Context, gitURL, name, teamID string) (*Composite, error) {
	team, err := c.GetTeamByID(ctx, teamID)
	composite := &Composite{}
	if err != nil {
		return composite, err
	}

	compositesLink, err := GetLinkByRef(team.Links, "composites")

	if err != nil {
		return composite, err
	}

	compositePayLoad := fmt.Sprintf(`{"name":"%s","gitUrl":"%s","teamId":"%s"}`, name, gitURL, teamID)
	var jsonStr = []byte(compositePayLoad)
	err = c.Do(ctx, "POST", compositesLink.Href, bytes.NewBuffer(jsonStr), &composite)
	if err != nil {
		return composite, err
	}

	return composite, nil
}
