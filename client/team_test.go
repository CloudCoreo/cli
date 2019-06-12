package client

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"

	"fmt"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

const userJSONPayloadForTeam = `{
	"username": "gitUser",
	"email": "user@cloudcoreo.com",
	"gravatarIconUrl": "http://www.gravatar.com/avatar/avatarID",
	"createdAt": "2016-11-26T09:22:40.356Z",
	"defaultTeamId": "teamID",
	"links": [
		{
			"ref": "teams",
			"method": "GET",
			"href": "%s/users/userID/teams"
		}
	],
	"id": "userID"
}`

const userJSONPayloadForTeamMissingTeamData = `{
	"username": "gitUser",
	"email": "user@cloudcoreo.com",
	"gravatarIconUrl": "http://www.gravatar.com/avatar/avatarID",
	"createdAt": "2016-11-26T09:22:40.356Z",
	"defaultTeamId": "teamID",
	"id": "userID"
}`

const teamJSONPayload = `[{
		"teamName": "gitUser-default",
			"ownerId": "userID",
			"teamIcon": "periodic-bg-5.png",
			"teamDescription": null,
			"default": true,
			"links": [
		{
		"ref": "self",
		"method": "GET",
		"href": "https://app.cloudcoreo.com/api/teams/teamID"
		},
		{
		"ref": "owner",
		"method": "GET",
		"href": "https://app.cloudcoreo.com/api/users/userID"
		},
		{
		"ref": "composites",
		"method": "GET",
		"href": "https://app.cloudcoreo.com/api/teams/teamID/composites"
		},
		{
		"ref": "users",
		"method": "GET",
		"href": "https://app.cloudcoreo.com/api/teams/teamID/users"
		},
		{
		"ref": "gitKeys",
		"method": "GET",
		"href": "https://app.cloudcoreo.com/api/teams/teamID/gitkeys"
		},
		{
		"ref": "cloudAccounts",
		"method": "GET",
		"href": "https://app.cloudcoreo.com/api/teams/teamID/cloudaccounts"
		}
	],
		"id": "teamID"
	}]`

const singleTeamJSONPayload = `{
		"teamName": "gitUser-default",
			"ownerId": "userID",
			"teamIcon": "periodic-bg-5.png",
			"teamDescription": null,
			"default": true,
			"links": [
		{
		"ref": "self",
		"method": "GET",
		"href": "https://app.cloudcoreo.com/api/teams/teamID"
		},
		{
		"ref": "owner",
		"method": "GET",
		"href": "https://app.cloudcoreo.com/api/users/userID"
		},
		{
		"ref": "composites",
		"method": "GET",
		"href": "https://app.cloudcoreo.com/api/teams/teamID/composites"
		},
		{
		"ref": "users",
		"method": "GET",
		"href": "https://app.cloudcoreo.com/api/teams/teamID/users"
		},
		{
		"ref": "gitKeys",
		"method": "GET",
		"href": "https://app.cloudcoreo.com/api/teams/teamID/gitkeys"
		},
		{
		"ref": "cloudAccounts",
		"method": "GET",
		"href": "https://app.cloudcoreo.com/api/teams/teamID/cloudaccounts"
		}
	],
		"id": "teamID"
	}`

func TestGetTeamsSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, teamJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspUrl+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetTeams(context.Background())
	assert.Nil(t, err, "GetTeams shouldn't return error.")
}

func TestGetTeamsFailureBadTeamResponse(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, ``))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspUrl+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetTeams(context.Background())
	assert.NotNil(t, err, "GetTeams should return error.")
}

func TestGetTeamsFailureBadUserResponse(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, ``))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, ``))
	httpmock.RegisterResponder("POST", cspUrl+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetTeams(context.Background())
	assert.NotNil(t, err, "GetTeams should return error.")
}

func TestGetTeamsFailureMissingTeamLinkInResponse(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, ``))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, userJSONPayloadForTeamMissingTeamData))
	httpmock.RegisterResponder("POST", cspUrl+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetTeams(context.Background())
	assert.NotNil(t, err, "GetTeams should return error.")
}

func TestGetTeamByIDSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, teamJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspUrl+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetTeamByID(context.Background(), "teamID")
	assert.Nil(t, err, "GetTeams shouldn't return error.")
}

func TestGetTeamByIDFailure(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, ``))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspUrl+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetTeamByID(context.Background(), "teamID")
	assert.NotNil(t, err, "GetTeams should return error.")
}

func TestGetTeamByIDFailureTeamIDNotFound(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, teamJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspUrl+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetTeamByID(context.Background(), "583bc8dbca5e631017ed46")
	assert.NotNil(t, err, "GetTeams should return error.")
}

func TestCreateTeamSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, singleTeamJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspUrl+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.CreateTeam(context.Background(), "TestName", "TestDescription")
	assert.Nil(t, err, "CreateTeam shouldn't return error.")
}
