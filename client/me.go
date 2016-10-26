package client

import (
	"golang.org/x/net/context"
)

// Me struct for api payload
type Me struct {
	Defaultteamuri  string `json:"defaultteamuri"`
	Teamsuri string `json:"teamsuri"`
	ID string `json:"id"`
	Self string `json:"self"`
	Createdat string `json:"createdat"`
	Defaultteam string `json:"defaultteam"`
	Gravatariconurl string `json:"gravatariconurl"`
	Email string `json:"email"`
	Username string `json:"username"`
}

// GetMe method for Me command
func (c *Client) GetMe(ctx context.Context) (Me, error) {
	t := Me{}
	err := c.Do(ctx, "GET", "me", nil, &t)
	if err != nil {
		return t, err
	}

	return t, nil
}
