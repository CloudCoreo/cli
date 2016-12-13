package client

import (
	"net/http"
	"testing"

	"fmt"

	"github.com/jharlap/httpstub"
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
			"href": "%s/api/users/userID/teams"
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

func TestGetTeamsSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(teamJSONPayload).WithStatus(http.StatusOK)

	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetTeams(context.Background())
	assert.Nil(t, err, "GetTeams shouldn't return error.")
}

func TestGetTeamsFailureBadTeamResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(``).WithStatus(http.StatusOK)

	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetTeams(context.Background())
	assert.NotNil(t, err, "GetTeams should return error.")
}

func TestGetTeamsFailureBadUserResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(``).WithStatus(http.StatusOK)

	ts.Path("/me").WithMethod("GET").WithBody(``).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetTeams(context.Background())
	assert.NotNil(t, err, "GetTeams should return error.")
}

func TestGetTeamsFailureMissingTeamLinkInResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(``).WithStatus(http.StatusOK)

	ts.Path("/me").WithMethod("GET").WithBody(userJSONPayloadForTeamMissingTeamData).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetTeams(context.Background())
	assert.NotNil(t, err, "GetTeams should return error.")
}

func TestGetTeamByIDSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(teamJSONPayload).WithStatus(http.StatusOK)

	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetTeamByID(context.Background(), "teamID")
	assert.Nil(t, err, "GetTeams shouldn't return error.")
}

func TestGetTeamByIDFailure(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(``).WithStatus(http.StatusOK)

	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetTeamByID(context.Background(), "teamID")
	assert.NotNil(t, err, "GetTeams should return error.")
}

func TestGetTeamByIDFailureTeamIDNotFound(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(teamJSONPayload).WithStatus(http.StatusOK)

	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetTeamByID(context.Background(), "583bc8dbca5e631017ed46")
	assert.NotNil(t, err, "GetTeams should return error.")
}
