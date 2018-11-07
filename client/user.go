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
	"time"

	"github.com/CloudCoreo/cli/pkg/command"
)

// User struct for api payload
type User struct {
	Username        string         `json:"username"`
	Email           string         `json:"email"`
	GravatarIconURL string         `json:"gravatarIconUrl"`
	CreatedAt       time.Time      `json:"createdAt"`
	DefaultTeamID   string         `json:"defaultTeamId"`
	Links           []command.Link `json:"links"`
	ID              string         `json:"id"`
}

// GetUser method for getting user info command
func (c *Client) GetUser(ctx context.Context) (*User, error) {
	t := &User{}
	err := c.Do(ctx, "GET", c.endpoint+"/me", nil, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}
