package client

import (
	"golang.org/x/net/context"
)

type Test struct {
	Msg        string `json:"msg"`
	CloudCoreo int    `json:"cloudoreo"`
}

func (c *Client) GetTest(ctx context.Context) (Test, error) {
	t := Test{}
	err := c.Do(ctx, "POST", "test", nil, &t)
	if err != nil {
		return t, err
	}

	return t, nil
}
