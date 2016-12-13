package client

import (
	"net/http"
	"testing"

	"fmt"

	"github.com/jharlap/httpstub"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

const teamGitKeyJSONPayload = `[{
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
		"ref": "users",
		"method": "GET",
		"href": "https://app.cloudcoreo.com/api/teams/teamID/users"
		},
		{
		"ref": "gitKeys",
		"method": "GET",
		"href": "%s/api/teams/teamID/gitkeys"
		}
	],
		"id": "teamID"
	}]`

const teamGitKeyJSONPayloadMissingGitKeyLink = `[{
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
		"ref": "users",
		"method": "GET",
		"href": "https://app.cloudcoreo.com/api/teams/teamID/users"
		}
	],
		"id": "teamID"
	}]`

const GitKeyJSONPayload = `[
	{
		"teamId": "teamID",
		"name": "Test key",
		"sha256fingerprint": "SHA256:Jbl/sadfasdfasf",
		"md5fingerprint": "123:123:123:123:123:123:123:123:123:3d:asdf:asdf:asdf:asdf:asdf:asf",
		"links": [
			{
				"ref": "self",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/gitkeys/gitkeyID"
			},
			{
				"ref": "team",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/teams/teamID"
			}
		],
		"id": "gitKeyID"
	}
]`

const GitKeyJSONPayloadMissingSelfData = `[{
		"teamId": "teamID",
		"name": "Test key",
		"sha256fingerprint": "SHA256:Jbl/sadfasdfasf",
		"md5fingerprint": "123:123:123:123:123:123:123:123:123:3d:asdf:asdf:asdf:asdf:asdf:asf",
		"id": "gitKeyID"
	}]`

const createdGitKeyJSONPayload = `{
		"id": "compositeID"
	}`

func TestGetGitKeysSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/gitkeys").WithMethod("GET").WithBody(GitKeyJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamGitKeyJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetGitKeys(context.Background(), "teamID")
	assert.Nil(t, err, "GetGitKeys shouldn't return error.")
}

func TestGetGitKeysFailureInvalidUserResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/gitkeys").WithMethod("GET").WithBody(GitKeyJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamGitKeyJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetGitKeys(context.Background(), "teamID")
	assert.NotNil(t, err, "GetGitKeys should return error.")
}

func TestGetGitKeysFailureInvalidTeamResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/gitkeys").WithMethod("GET").WithBody(GitKeyJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetGitKeys(context.Background(), "teamID")
	assert.NotNil(t, err, "GetGitKeys should return error.")
}

func TestGetGitKeysFailureInvalidGitKeyresponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/gitkeys").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamGitKeyJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetGitKeys(context.Background(), "teamID")
	assert.NotNil(t, err, "GetGitKeys should return error.")
}

func TestGetGitKeysFailureMissingGitKeysLink(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/gitkeys").WithMethod("GET").WithBody(GitKeyJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(teamGitKeyJSONPayloadMissingGitKeyLink).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetGitKeys(context.Background(), "teamID")
	assert.NotNil(t, err, "GetGitKeys should return error.")
	assert.Equal(t, "Resource for given ID not found.", err.Error())
}

func TestGetGitKeysFailedNoGitKeysFound(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/gitkeys").WithMethod("GET").WithBody(`[]`).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamGitKeyJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetGitKeys(context.Background(), "teamID")
	assert.NotNil(t, err, "GetGitKeys should return error.")
	assert.Equal(t, "No git keys found under team ID teamID.", err.Error())

}

func TestGetGitKeyByIDSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/gitkeys").WithMethod("GET").WithBody(GitKeyJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamGitKeyJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetGitKeyByID(context.Background(), "teamID", "gitKeyID")
	assert.Nil(t, err, "GetGitKeyByID shouldn't return error.")
}

func TestGetGitKeyByIDFailureInvalidTeamID(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/gitkeys").WithMethod("GET").WithBody(GitKeyJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamGitKeyJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetGitKeyByID(context.Background(), "invalidTeamID", "gitKeyID")
	assert.NotNil(t, err, "GetGitKeyByID should return error.")
	assert.Equal(t, "No git keys found under team ID invalidTeamID.", err.Error())

}

func TestGetGitKeyByIDFailureInvalidGitKeyID(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/gitkeys").WithMethod("GET").WithBody(GitKeyJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamGitKeyJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetGitKeyByID(context.Background(), "teamID", "invalidGitKeyID")
	assert.NotNil(t, err, "GetGitKeyByID should return error.")
	assert.Equal(t, "No git key with ID invalidGitKeyID found under team ID teamID.", err.Error())
}

func TestCreateGitKeySuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/gitkeys").WithMethod("POST").WithBody(createdGitKeyJSONPayload).WithStatus(http.StatusCreated)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamGitKeyJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateGitKey(context.Background(), "teamID", "keyMaterial", "name")
	assert.Nil(t, err, "CreateGitKey shouldn't return error.")
}

func TestCreateGitKeyFailureBadRequest(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/gitkeys").WithMethod("POST").WithBody(createdGitKeyJSONPayload).WithStatus(http.StatusBadRequest)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamGitKeyJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateGitKey(context.Background(), "teamID", "keyMaterial", "name")
	assert.NotNil(t, err, "CreateGitKey should return error.")
}

func TestCreateGitKeyFailedToParseUser(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/gitkeys").WithMethod("POST").WithBody(createdGitKeyJSONPayload).WithStatus(http.StatusCreated)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamGitKeyJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateGitKey(context.Background(), "teamID", "keyMaterial", "name")
	assert.NotNil(t, err, "CreateGitKey should return error.")
}

func TestCreateGitKeyFailedToParseGitKeyLink(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/gitkeys").WithMethod("POST").WithBody(createdGitKeyJSONPayload).WithStatus(http.StatusCreated)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateGitKey(context.Background(), "teamID", "keyMaterial", "name")
	assert.NotNil(t, err, "CreateGitKey should return error.")
	assert.Equal(t, "json: cannot unmarshal object into Go value of type []*client.Team", err.Error())
}

func TestCreateGitKeysFailureMissingGitKeysLink(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/gitkeys").WithMethod("POST").WithBody(createdGitKeyJSONPayload).WithStatus(http.StatusCreated)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(teamGitKeyJSONPayloadMissingGitKeyLink).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateGitKey(context.Background(), "teamID", "keyMaterial", "name")
	assert.NotNil(t, err, "CreateGitKey should return error.")
	assert.Equal(t, "Resource for given ID not found.", err.Error())
}

func TestCreateGitKeyFailureGitKeyNotCreated(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/gitkeys").WithMethod("POST").WithBody(`{}`).WithStatus(http.StatusCreated)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamGitKeyJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateGitKey(context.Background(), "teamID", "keyMaterial", "name")
	assert.NotNil(t, err, "CreateGitKey should return error.")
	assert.Equal(t, "Failed to create git key under team ID teamID.", err.Error())
}
