package client

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

const defaultAPIEndpoint = "https://app.hack.securestate.vmware.com/api"

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

const refreshTokenJSONPayload = `{
    "id_token": "fake-id",
    "token_type": "bearer",
    "expires_in": 1799,
    "scope": "ALL_PERMISSIONS customer_number openid",
    "access_token": "fake-access-token",
    "refresh_token": "fake-refresh-token"
}`

func TestGetUserSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, userJSONPayload))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetUser(context.Background())
	assert.Nil(t, err, "GetUser shouldn't return error.")
}

func TestGetUserFailure(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, ``))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetUser(context.Background())
	assert.NotNil(t, err, "GetUser should return error.")
}
