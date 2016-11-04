package client

import (
	"time"
	"fmt"
	"bytes"

	"golang.org/x/net/context"
)

type Token struct {
	Name string `json:"name"`
	Description string `json:"description"`
	CreationDate time.Time `json:"creationDate"`
	Links []Link `json:"links"`
	ID string `json:"id"`
}

// GetTokens method for token command
func (c *Client) GetTokens(ctx context.Context) ([]Token, error) {
	t := []Token{}
	u, err := c.GetUser(ctx)

	if err != nil {
		return t, err
	}

	tokenLink := GetLinkByRef(u.Links, "tokens")

	err = c.Do(ctx, "GET", tokenLink.Href, nil, &t)
	if err != nil {
		return t, err
	}

	return t, nil
}

// GetTokenByID method for token command
func (c *Client) GetTokenByID(ctx context.Context, tokenID string) (Token, error) {
	token := Token{}

	tokens, err := c.GetTokens(ctx)

	if err != nil {
		return token, err
	}

	for _, t := range tokens {
		if t.ID == tokenID {
			token = t
			break
		}
	}

	return token, nil
}

// CreateToken method to create a token object
func (c *Client) CreateToken(ctx context.Context, description, name string) (Token, error)  {
	u, err := c.GetUser(ctx)
	token := Token{}
	if err != nil {
		return token, err
	}

	tokenLink := GetLinkByRef(u.Links, "tokens")

	tokenPlayLoad := fmt.Sprintf(`{"description":"%s","name":"%s"}`, description, name)
	var jsonStr = []byte(tokenPlayLoad)
	err = c.Do(ctx, "POST", tokenLink.Href, bytes.NewBuffer(jsonStr), &token)
	if err != nil {
		return token, err
	}

	return token, nil
}

// DeleteTokenByID method to delete token object
func (c *Client) DeleteTokenByID(ctx context.Context, tokenID string) error {
	tokens, err := c.GetTokens(ctx)

	if err != nil {
		return err
	}

	for _, token := range tokens {
		if token.ID == tokenID {
			tokensink := GetLinkByRef(token.Links, "self")
			err := c.Do(ctx, "DELETE", tokensink.Href, nil, nil)
			if err != nil {
				return err
			}
			break
		}
	}

	return nil
}