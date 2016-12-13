package client

import (
	"net/http"
	"testing"

	"github.com/jharlap/httpstub"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

const userJSONPayload = `{
	"username": "coolguru",
	"email": "nandesh@cloudcoreo.com",
	"gravatarIconUrl": "http://www.gravatar.com/avatar/ff813888fedf42ea2f849b2fba9e9de8",
	"createdAt": "2016-11-26T09:22:40.356Z",
	"defaultTeamId": "583bc8dbca5e631017ed46c9",
	"links": [
		{
			"ref": "self",
			"method": "GET",
			"href": "https://app.cloudcoreo.com/api/users/583bc8dbca5e631017ed46c8"
		},
		{
			"ref": "defaultTeam",
			"method": "GET",
			"href": "https://app.cloudcoreo.com/api/teams/583bc8dbca5e631017ed46c9"
		},
		{
			"ref": "teams",
			"method": "GET",
			"href": "https://app.cloudcoreo.com/api/users/583bc8dbca5e631017ed46c8/teams"
		},
		{
			"ref": "tokens",
			"method": "GET",
			"href": "https://app.cloudcoreo.com/api/users/583bc8dbca5e631017ed46c8/tokens"
		}
	],
	"id": "583bc8dbca5e631017ed46c8"
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
