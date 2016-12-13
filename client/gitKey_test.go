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
		"ref": "users",
		"method": "GET",
		"href": "https://app.cloudcoreo.com/api/teams/583bc8dbca5e631017ed46c9/users"
		},
		{
		"ref": "gitKeys",
		"method": "GET",
		"href": "%s/api/teams/583bc8dbca5e631017ed46c9/gitkeys"
		}
	],
		"id": "583bc8dbca5e631017ed46c9"
	}]`

const teamGitKeyJSONPayloadMissingGitKeyLink = `[{
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
		"ref": "users",
		"method": "GET",
		"href": "https://app.cloudcoreo.com/api/teams/583bc8dbca5e631017ed46c9/users"
		}
	],
		"id": "583bc8dbca5e631017ed46c9"
	}]`

const GitKeyJSONPayload = `[
	{
		"teamId": "583bc8dbca5e631017ed46c9",
		"name": "Test key",
		"sha256fingerprint": "SHA256:Jbl/YDLNXwIu9w1eqm6CN4HSEsWJplkDydy7AQ9al9g",
		"md5fingerprint": "39:18:b9:c9:aa:b6:5c:15:8f:f9:a2:7a:cf:24:23:c1",
		"links": [
			{
				"ref": "self",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/gitkeys/585049537bb23bb35859ee5e"
			},
			{
				"ref": "team",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/teams/583bc8dbca5e631017ed46c9"
			}
		],
		"id": "gitKeyID"
	}
]`

const GitKeyJSONPayloadMissingSelfData = `[{
		"teamId": "583bc8dbca5e631017ed46c9",
		"name": "Test key",
		"sha256fingerprint": "SHA256:Jbl/YDLNXwIu9w1eqm6CN4HSEsWJplkDydy7AQ9al9g",
		"md5fingerprint": "39:18:b9:c9:aa:b6:5c:15:8f:f9:a2:7a:cf:24:23:c1",
		"id": "gitKeyID"
	}]`

const createdGitKeyJSONPayload = `{
		"id": "583bca6dca5e631017ed46fb"
	}`

func TestGetGitKeysSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/gitkeys").WithMethod("GET").WithBody(GitKeyJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamGitKeyJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetGitKeys(context.Background(), "583bc8dbca5e631017ed46c9")
	assert.Nil(t, err, "GetGitKeys shouldn't return error.")
}

func TestGetGitKeysFailureInvalidUserResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/gitkeys").WithMethod("GET").WithBody(GitKeyJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamGitKeyJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetGitKeys(context.Background(), "583bc8dbca5e631017ed46c9")
	assert.NotNil(t, err, "GetGitKeys should return error.")
}

func TestGetGitKeysFailureInvalidTeamResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/gitkeys").WithMethod("GET").WithBody(GitKeyJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetGitKeys(context.Background(), "583bc8dbca5e631017ed46c9")
	assert.NotNil(t, err, "GetGitKeys should return error.")
}

func TestGetGitKeysFailureInvalidGitKeyresponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/gitkeys").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamGitKeyJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetGitKeys(context.Background(), "583bc8dbca5e631017ed46c9")
	assert.NotNil(t, err, "GetGitKeys should return error.")
}

func TestGetGitKeysFailureMissingGitKeysLink(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/gitkeys").WithMethod("GET").WithBody(GitKeyJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(teamGitKeyJSONPayloadMissingGitKeyLink).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetGitKeys(context.Background(), "583bc8dbca5e631017ed46c9")
	assert.NotNil(t, err, "GetGitKeys should return error.")
	assert.Equal(t, "Resource for given ID not found.", err.Error())
}

func TestGetGitKeysFailedNoGitKeysFound(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/gitkeys").WithMethod("GET").WithBody(`[]`).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamGitKeyJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetGitKeys(context.Background(), "583bc8dbca5e631017ed46c9")
	assert.NotNil(t, err, "GetGitKeys should return error.")
	assert.Equal(t, "No git keys found under team ID 583bc8dbca5e631017ed46c9.", err.Error())

}

func TestGetGitKeyByIDSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/gitkeys").WithMethod("GET").WithBody(GitKeyJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamGitKeyJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetGitKeyByID(context.Background(), "583bc8dbca5e631017ed46c9", "gitKeyID")
	assert.Nil(t, err, "GetGitKeyByID shouldn't return error.")
}

func TestGetGitKeyByIDFailureInvalidTeamID(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/gitkeys").WithMethod("GET").WithBody(GitKeyJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamGitKeyJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetGitKeyByID(context.Background(), "invalidTeamID", "gitKeyID")
	assert.NotNil(t, err, "GetGitKeyByID should return error.")
	assert.Equal(t, "No git keys found under team ID invalidTeamID.", err.Error())

}

func TestGetGitKeyByIDFailureInvalidGitKeyID(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/gitkeys").WithMethod("GET").WithBody(GitKeyJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamGitKeyJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetGitKeyByID(context.Background(), "583bc8dbca5e631017ed46c9", "invalidGitKeyID")
	assert.NotNil(t, err, "GetGitKeyByID should return error.")
	assert.Equal(t, "No git key with ID invalidGitKeyID found under team ID 583bc8dbca5e631017ed46c9.", err.Error())
}

func TestCreateGitKeySuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/gitkeys").WithMethod("POST").WithBody(createdGitKeyJSONPayload).WithStatus(http.StatusCreated)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamGitKeyJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateGitKey(context.Background(), "583bc8dbca5e631017ed46c9", "keyMaterial", "name")
	assert.Nil(t, err, "CreateGitKey shouldn't return error.")
}

func TestCreateGitKeyFailureBadRequest(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/gitkeys").WithMethod("POST").WithBody(createdGitKeyJSONPayload).WithStatus(http.StatusBadRequest)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamGitKeyJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateGitKey(context.Background(), "583bc8dbca5e631017ed46c9", "keyMaterial", "name")
	assert.NotNil(t, err, "CreateGitKey should return error.")
}

func TestCreateGitKeyFailedToParseUser(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/gitkeys").WithMethod("POST").WithBody(createdGitKeyJSONPayload).WithStatus(http.StatusCreated)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamGitKeyJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateGitKey(context.Background(), "583bc8dbca5e631017ed46c9", "keyMaterial", "name")
	assert.NotNil(t, err, "CreateGitKey should return error.")
}

func TestCreateGitKeyFailedToParseGitKeyLink(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/gitkeys").WithMethod("POST").WithBody(createdGitKeyJSONPayload).WithStatus(http.StatusCreated)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateGitKey(context.Background(), "583bc8dbca5e631017ed46c9", "keyMaterial", "name")
	assert.NotNil(t, err, "CreateGitKey should return error.")
	assert.Equal(t, "json: cannot unmarshal object into Go value of type []*client.Team", err.Error())
}

func TestCreateGitKeysFailureMissingGitKeysLink(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/gitkeys").WithMethod("POST").WithBody(createdGitKeyJSONPayload).WithStatus(http.StatusCreated)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(teamGitKeyJSONPayloadMissingGitKeyLink).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateGitKey(context.Background(), "583bc8dbca5e631017ed46c9", "keyMaterial", "name")
	assert.NotNil(t, err, "CreateGitKey should return error.")
	assert.Equal(t, "Resource for given ID not found.", err.Error())
}

func TestCreateGitKeyFailureGitKeyNotCreated(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/gitkeys").WithMethod("POST").WithBody(`{}`).WithStatus(http.StatusCreated)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamGitKeyJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateGitKey(context.Background(), "583bc8dbca5e631017ed46c9", "keyMaterial", "name")
	assert.NotNil(t, err, "CreateGitKey should return error.")
	assert.Equal(t, "Failed to create git key under team ID 583bc8dbca5e631017ed46c9.", err.Error())
}
