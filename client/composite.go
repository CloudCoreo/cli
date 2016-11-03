package client

import (
	"time"

	"golang.org/x/net/context"
)

// Composite struct for api payload
type Composite struct {
	TeamID string `json:"teamId"`
	Name string `json:"name"`
	GitURL string `json:"gitUrl"`
	HasCustomDashboard bool `json:"hasCustomDashboard"`
	CreatedAt time.Time `json:"createdAt"`
	Self string `json:"self"`
	TeamURI string `json:"teamUri"`
	ID string `json:"id"`
	GitKeyURI string `json:"gitKeyUri"`
	PlansURI string `json:"plansUri"`
}

// GetComposites method to get composite info array object
func (c *Client) GetComposites(ctx context.Context, teamID string) ([]Composite, error) {
	composites := []Composite{}
	team, err := c.GetTeam(ctx, teamID)

	if err != nil {
		return composites, err
	}

	teamLink := GetLinkByRef(team.Links, "composite")

	err = c.Do(ctx, "GET", teamLink.Href, nil, &composites)
	if err != nil {
		return composites, err
	}

	return composites, nil
}

// GetComposite method to get composite info object
func (c *Client) GetComposite(ctx context.Context, teamID, compositeID string) (Composite, error) {
	composite := Composite{}
	composites, err := c.GetComposites(ctx, teamID)
	if err != nil {
		return composite, err
	}

	for _, comp := range composites {
		if comp.ID == compositeID {
			composite = comp
			break
		}
	}
	return composite, nil
}
