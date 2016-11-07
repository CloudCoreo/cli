package client

import (
	"bytes"
	"fmt"
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
	team, err := c.GetTeamByID(ctx, teamID)

	if err != nil {
		return composites, err
	}

	compsoitesLink, err := GetLinkByRef(team.Links, "composites")
	if err != nil {
		return composites, err
	}

	err = c.Do(ctx, "GET", compsoitesLink.Href, nil, &composites)
	if err != nil {
		return composites, err
	}

	return composites, nil
}

// GetComposite method to get composite info object
func (c *Client) GetCompositeByID(ctx context.Context, teamID, compositeID string) (Composite, error) {
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

// CreateComposite method to create a composite object
func (c *Client) CreateComposite(ctx context.Context, gitURL, name, teamID string) (Composite, error)  {
	team, err := c.GetTeamByID(ctx, teamID)
	composite := Composite{}
	if err != nil {
		return composite, err
	}

	compositesLink, err := GetLinkByRef(team.Links, "composites")

	if err != nil {
		return composite, err
	}

	compositePlayLoad := fmt.Sprintf(`{"name":"%s","gitUrl":"%s","teamId":"%s"}`, name, gitURL, teamID)
	var jsonStr = []byte(compositePlayLoad)
	err = c.Do(ctx, "POST", compositesLink.Href, bytes.NewBuffer(jsonStr), &composite)
	if err != nil {
		return composite, err
	}

	return composite, nil
}