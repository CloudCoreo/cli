package client

import (
	"time"

	"golang.org/x/net/context"
)

// User struct for api payload
type User struct {
	Username        string    `json:"username"`
	Email           string    `json:"email"`
	GravatarIconURL string    `json:"gravatarIconUrl"`
	CreatedAt       time.Time `json:"createdAt"`
	DefaultTeamID   string    `json:"defaultTeamId"`
	Links           []Link    `json:"links"`
	ID              string    `json:"id"`
}

// GetUser method for getting user info command
func (c *Client) GetUser(ctx context.Context) (User, error) {
	t := User{}
	err := c.Do(ctx, "GET", c.endpoint+"/me", nil, &t)
	if err != nil {
		return t, err
	}

	return t, nil
}
