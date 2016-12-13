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
	"username": "coolguru",
	"email": "nandesh@cloudcoreo.com",
	"gravatarIconUrl": "http://www.gravatar.com/avatar/ff813888fedf42ea2f849b2fba9e9de8",
	"createdAt": "2016-11-26T09:22:40.356Z",
	"defaultTeamId": "583bc8dbca5e631017ed46c9",
	"links": [
		{
			"ref": "teams",
			"method": "GET",
			"href": "%s/api/users/583bc8dbca5e631017ed46c8/teams"
		}
	],
	"id": "583bc8dbca5e631017ed46c8"
}`

const userJSONPayloadForTeamMissingTeamData = `{
	"username": "coolguru",
	"email": "nandesh@cloudcoreo.com",
	"gravatarIconUrl": "http://www.gravatar.com/avatar/ff813888fedf42ea2f849b2fba9e9de8",
	"createdAt": "2016-11-26T09:22:40.356Z",
	"defaultTeamId": "583bc8dbca5e631017ed46c9",
	"id": "583bc8dbca5e631017ed46c8"
}`

const teamJSONPayload = `[{
		"teamName": "coolguru-default",
			"ownerId": "583bc8dbca5e631017ed46c8",
			"teamIcon": "periodic-bg-5.png",
			"teamDescription": null,
			"default": true,
			"links": [
		{
		"ref": "self",
		"method": "GET",
		"href": "https://app.cloudcoreo.com/api/teams/583bc8dbca5e631017ed46c9"
		},
		{
		"ref": "owner",
		"method": "GET",
		"href": "https://app.cloudcoreo.com/api/users/583bc8dbca5e631017ed46c8"
		},
		{
		"ref": "composites",
		"method": "GET",
		"href": "https://app.cloudcoreo.com/api/teams/583bc8dbca5e631017ed46c9/composites"
		},
		{
		"ref": "users",
		"method": "GET",
		"href": "https://app.cloudcoreo.com/api/teams/583bc8dbca5e631017ed46c9/users"
		},
		{
		"ref": "gitKeys",
		"method": "GET",
		"href": "https://app.cloudcoreo.com/api/teams/583bc8dbca5e631017ed46c9/gitkeys"
		},
		{
		"ref": "cloudAccounts",
		"method": "GET",
		"href": "https://app.cloudcoreo.com/api/teams/583bc8dbca5e631017ed46c9/cloudaccounts"
		}
	],
		"id": "583bc8dbca5e631017ed46c9"
	}]`

func TestGetTeamsSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(teamJSONPayload).WithStatus(http.StatusOK)

	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetTeams(context.Background())
	assert.Nil(t, err, "GetTeams shouldn't return error.")
}

func TestGetTeamsFailureBadTeamResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(``).WithStatus(http.StatusOK)

	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetTeams(context.Background())
	assert.NotNil(t, err, "GetTeams should return error.")
}

func TestGetTeamsFailureBadUserResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(``).WithStatus(http.StatusOK)

	ts.Path("/me").WithMethod("GET").WithBody(``).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetTeams(context.Background())
	assert.NotNil(t, err, "GetTeams should return error.")
}

func TestGetTeamsFailureMissingTeamLinkInResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(``).WithStatus(http.StatusOK)

	ts.Path("/me").WithMethod("GET").WithBody(userJSONPayloadForTeamMissingTeamData).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetTeams(context.Background())
	assert.NotNil(t, err, "GetTeams should return error.")
}

func TestGetTeamByIDSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(teamJSONPayload).WithStatus(http.StatusOK)

	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetTeamByID(context.Background(), "583bc8dbca5e631017ed46c9")
	assert.Nil(t, err, "GetTeams shouldn't return error.")
}

func TestGetTeamByIDFailure(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(``).WithStatus(http.StatusOK)

	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetTeamByID(context.Background(), "583bc8dbca5e631017ed46c9")
	assert.NotNil(t, err, "GetTeams should return error.")
}

func TestGetTeamByIDFailureTeamIDNotFound(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(teamJSONPayload).WithStatus(http.StatusOK)

	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetTeamByID(context.Background(), "583bc8dbca5e631017ed46")
	assert.NotNil(t, err, "GetTeams should return error.")
}
