package client

import (
	"golang.org/x/net/context"
)

// Me struct for api payload
type Me struct {
	Defaultteamuri  string `json:"defaultTeamUri"`
	Teamsuri string `json:"teamsUri"`
	ID string `json:"id"`
	Self string `json:"self"`
	Createdat string `json:"createdAt"`
	Defaultteam string `json:"defaultTeam"`
	Gravatariconurl string `json:"gravatarIconUrl"`
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
