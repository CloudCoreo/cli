package client

import (
	"net/http"
	"testing"

	"fmt"

	"github.com/jharlap/httpstub"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

const teamCompositeJSONPayload = `[{
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
		"href": "%s/api/teams/583bc8dbca5e631017ed46c9/composites"
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
		}
	],
		"id": "583bc8dbca5e631017ed46c9"
	}]`

const teamCompositeJSONPayloadMissingCompositeLink = `[{
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
		"href": "https://app.cloudcoreo.com/api/teams/583bc8dbca5e631017ed46c9/gitkeys"
		}
	],
		"id": "583bc8dbca5e631017ed46c9"
	}]`

const CompositeJSONPayload = `[
	{
		"name": "audit-aws-s3",
		"gitUrl": "git@github.com:CloudCoreo/audit-aws-s3.git",
		"hasCustomDashboard": false,
		"createdAt": "2016-11-28T06:10:53.903Z",
		"gitKeyId": "0",
		"teamId": "583bc8dbca5e631017ed46c9",
		"id": "583bca6dca5e631017ed46fb",
		"links": [
			{
				"ref": "self",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/composites/583bca6dca5e631017ed46fb"
			},
			{
				"ref": "team",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/teams/583bc8dbca5e631017ed46c9"
			},
			{
				"ref": "gitKey",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/gitkeys/0"
			},
			{
				"ref": "plans",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/composites/583bca6dca5e631017ed46fb/plans"
			}
		]
	}
]`

const CompositeJSONPayloadMissingSelfData = `[{
		"name": "audit-aws-s3",
		"gitUrl": "git@github.com:CloudCoreo/audit-aws-s3.git",
		"hasCustomDashboard": false,
		"createdAt": "2016-11-28T06:10:53.903Z",
		"gitKeyId": "0",
		"teamId": "583bc8dbca5e631017ed46c9",
		"id": "583bca6dca5e631017ed46fb"
	}]`

const createdCompositeJSONPayload = `{
		"id": "583bca6dca5e631017ed46fb"
	}`

func TestGetCompositesSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(CompositeJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetComposites(context.Background(), "583bc8dbca5e631017ed46c9")
	assert.Nil(t, err, "GetComposites shouldn't return error.")
}

func TestGetCompositesFailureInvalidUserResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(CompositeJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetComposites(context.Background(), "583bc8dbca5e631017ed46c9")
	assert.NotNil(t, err, "GetComposites should return error.")
}

func TestGetCompositesFailureInvalidTeamResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(CompositeJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetComposites(context.Background(), "583bc8dbca5e631017ed46c9")
	assert.NotNil(t, err, "GetComposites should return error.")
}

func TestGetCompositesFailureInvalidCompositeresponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetComposites(context.Background(), "583bc8dbca5e631017ed46c9")
	assert.NotNil(t, err, "GetComposites should return error.")
}

func TestGetCompositesFailureMissingCompositesLink(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(CompositeJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(teamCompositeJSONPayloadMissingCompositeLink).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetComposites(context.Background(), "583bc8dbca5e631017ed46c9")
	assert.NotNil(t, err, "GetComposites should return error.")
	assert.Equal(t, "Resource for given ID not found.", err.Error())
}

func TestGetCompositesFailedNoCompositesFound(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(`[]`).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetComposites(context.Background(), "583bc8dbca5e631017ed46c9")
	assert.NotNil(t, err, "GetComposites should return error.")
	assert.Equal(t, "No composites found under team ID 583bc8dbca5e631017ed46c9.", err.Error())

}

func TestGetCompositeByIDSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(CompositeJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetCompositeByID(context.Background(), "583bc8dbca5e631017ed46c9", "583bca6dca5e631017ed46fb")
	assert.Nil(t, err, "GetCompositeByID shouldn't return error.")
}

func TestGetCompositeByIDFailureInvalidTeamID(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(CompositeJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetCompositeByID(context.Background(), "invalidTeamID", "583bca6dca5e631017ed46fb")
	assert.NotNil(t, err, "GetCompositeByID should return error.")
	assert.Equal(t, "No team with ID invalidTeamID found.", err.Error())

}

func TestGetCompositeByIDFailureInvalidCompositeID(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(CompositeJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetCompositeByID(context.Background(), "583bc8dbca5e631017ed46c9", "invalidCompositeID")
	assert.NotNil(t, err, "GetCompositeByID should return error.")
	assert.Equal(t, "No composite with ID invalidCompositeID found under team ID 583bc8dbca5e631017ed46c9.", err.Error())
}

func TestCreateCompositeSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("POST").WithBody(createdCompositeJSONPayload).WithStatus(http.StatusCreated)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateComposite(context.Background(), "gitURL", "name", "583bc8dbca5e631017ed46c9")
	assert.Nil(t, err, "CreateComposite shouldn't return error.")
}

func TestCreateCompositeFailureBadRequest(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("POST").WithBody(createdCompositeJSONPayload).WithStatus(http.StatusBadRequest)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateComposite(context.Background(), "gitURL", "name", "583bc8dbca5e631017ed46c9")
	assert.NotNil(t, err, "CreateComposite should return error.")
}

func TestCreateCompositeFailedToParseUser(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/Composites").WithMethod("POST").WithBody(createdCompositeJSONPayload).WithStatus(http.StatusCreated)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateComposite(context.Background(), "gitURL", "name", "583bc8dbca5e631017ed46c9")
	assert.NotNil(t, err, "CreateComposite should return error.")
}

func TestCreateCompositeFailedToParseCompositeLink(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("POST").WithBody(createdCompositeJSONPayload).WithStatus(http.StatusCreated)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateComposite(context.Background(), "gitURL", "name", "583bc8dbca5e631017ed46c9")
	assert.NotNil(t, err, "CreateComposite should return error.")
	assert.Equal(t, "json: cannot unmarshal object into Go value of type []*client.Team", err.Error())
}

func TestCreateCompositesFailureMissingCompositesLink(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("POST").WithBody(createdCompositeJSONPayload).WithStatus(http.StatusCreated)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(teamCompositeJSONPayloadMissingCompositeLink).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateComposite(context.Background(), "gitURL", "name", "583bc8dbca5e631017ed46c9")
	assert.NotNil(t, err, "CreateComposite should return error.")
	assert.Equal(t, "Resource for given ID not found.", err.Error())
}

func TestCreateCompositeFailureCompositeNotCreated(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("POST").WithBody(`{}`).WithStatus(http.StatusCreated)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateComposite(context.Background(), "gitURL", "name", "583bc8dbca5e631017ed46c9")
	assert.NotNil(t, err, "CreateComposite should return error.")
	assert.Equal(t, "Failed to create composite under team ID 583bc8dbca5e631017ed46c9.", err.Error())
}
