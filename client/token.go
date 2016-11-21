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

	"github.com/CloudCoreo/cli/client/content"
)

// Token struct
type Token struct {
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	CreationDate time.Time `json:"creationDate"`
	Links        []Link    `json:"links"`
	ID           string    `json:"id"`
}

// GetTokens method for token command
func (c *Client) GetTokens(ctx context.Context) ([]*Token, error) {
	tokens := []*Token{}
	u, err := c.GetUser(ctx)

	if err != nil {
		return nil, err
	}

	tokenLink, err := GetLinkByRef(u.Links, "tokens")

	if err != nil {
		return nil, err
	}

	err = c.Do(ctx, "GET", tokenLink.Href, nil, &tokens)
	if err != nil {
		return nil, err
	}

	if len(tokens) == 0 {
		return nil, NewError(fmt.Sprintf(content.ErrorNoTokensFound))
	}

	return tokens, nil
}

// GetTokenByID method for token command
func (c *Client) GetTokenByID(ctx context.Context, tokenID string) (*Token, error) {
	tokens, err := c.GetTokens(ctx)

	if err != nil {
		return nil, err
	}

	token := &Token{}

	for _, t := range tokens {
		if t.ID == tokenID {
			token = t
			break
		}
	}

	if token.ID == "" {
		return nil, NewError(fmt.Sprintf(content.ErrorNoTokenWithIDFound, tokenID))
	}

	return token, nil
}

// CreateToken method to create a token object
func (c *Client) CreateToken(ctx context.Context, name, description string) (*Token, error) {
	u, err := c.GetUser(ctx)
	token := &Token{}
	if err != nil {
		return token, err
	}

	tokenLink, err := GetLinkByRef(u.Links, "tokens")

	if err != nil {
		return token, err
	}

	tokenPayLoad := fmt.Sprintf(`{"description":"%s","name":"%s"}`, description, name)
	var jsonStr = []byte(tokenPayLoad)
	err = c.Do(ctx, "POST", tokenLink.Href, bytes.NewBuffer(jsonStr), &token)
	if err != nil {
		return token, err
	}

	if token.ID == "" {
		return nil, NewError(fmt.Sprintf(content.ErrorFailedTokenCreation))
	}

	return token, nil
}

// DeleteTokenByID method to delete token object
func (c *Client) DeleteTokenByID(ctx context.Context, tokenID string) error {
	tokens, err := c.GetTokens(ctx)

	if err != nil {
		return err
	}

	tokenFound := false
	for _, token := range tokens {
		if token.ID == tokenID {
			tokenFound = true
			tokenslink, err := GetLinkByRef(token.Links, "self")

			if err != nil {
				return err
			}

			err = c.Do(ctx, "DELETE", tokenslink.Href, nil, nil)
			if err != nil {
				return err
			}
			break
		}
	}

	if !tokenFound {
		return NewError(fmt.Sprintf(content.ErrorFailedToDeleteToken, tokenID))
	}

	return nil
}
