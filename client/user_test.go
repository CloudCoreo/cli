package client

import (
	"net/http"
	"testing"

	"github.com/jharlap/httpstub"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

const userJSONPayload = `{
	"username": "gitUser",
	"email": "user@cloudcoreo.com",
	"gravatarIconUrl": "http://www.gravatar.com/avatar/avatarID",
	"createdAt": "2016-11-26T09:22:40.356Z",
	"defaultTeamId": "teamID",
	"links": [
		{
			"ref": "self",
			"method": "GET",
			"href": "https://app.cloudcoreo.com/api/users/userID"
		},
		{
			"ref": "defaultTeam",
			"method": "GET",
			"href": "https://app.cloudcoreo.com/api/teams/teamID"
		},
		{
			"ref": "teams",
			"method": "GET",
			"href": "https://app.cloudcoreo.com/api/users/userID/teams"
		},
		{
			"ref": "tokens",
			"method": "GET",
			"href": "https://app.cloudcoreo.com/api/users/userID/tokens"
		}
	],
	"id": "userID"
}`

func TestGetUserSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/me").WithBody(userJSONPayload).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetUser(context.Background())
	assert.Nil(t, err, "GetUser shouldn't return error.")
}

func TestGetUserFailure(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/me").WithBody(``).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetUser(context.Background())
	assert.NotNil(t, err, "GetUser should return error.")
}
