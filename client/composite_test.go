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
		"href": "%s/api/teams/teamID/composites"
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
		}
	],
		"id": "teamID"
	}]`

const teamCompositeJSONPayloadMissingCompositeLink = `[{
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
		"href": "https://app.cloudcoreo.com/api/teams/teamID/gitkeys"
		}
	],
		"id": "teamID"
	}]`

const CompositeJSONPayload = `[
	{
		"name": "audit-aws-s3",
		"gitUrl": "git@github.com:CloudCoreo/audit-aws-s3.git",
		"hasCustomDashboard": false,
		"createdAt": "2016-11-28T06:10:53.903Z",
		"gitKeyId": "0",
		"teamId": "teamID",
		"id": "compositeID",
		"links": [
			{
				"ref": "self",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/composites/compositeID"
			},
			{
				"ref": "team",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/teams/teamID"
			},
			{
				"ref": "gitKey",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/gitkeys/0"
			},
			{
				"ref": "plans",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/composites/compositeID/plans"
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
		"teamId": "teamID",
		"id": "compositeID"
	}]`

const createdCompositeJSONPayload = `{
		"id": "compositeID"
	}`

func TestGetCompositesSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(CompositeJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetComposites(context.Background(), "teamID")
	assert.Nil(t, err, "GetComposites shouldn't return error.")
}

func TestGetCompositesFailureInvalidUserResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(CompositeJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetComposites(context.Background(), "teamID")
	assert.NotNil(t, err, "GetComposites should return error.")
}

func TestGetCompositesFailureInvalidTeamResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(CompositeJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetComposites(context.Background(), "teamID")
	assert.NotNil(t, err, "GetComposites should return error.")
}

func TestGetCompositesFailureInvalidCompositeresponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetComposites(context.Background(), "teamID")
	assert.NotNil(t, err, "GetComposites should return error.")
}

func TestGetCompositesFailureMissingCompositesLink(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(CompositeJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(teamCompositeJSONPayloadMissingCompositeLink).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetComposites(context.Background(), "teamID")
	assert.NotNil(t, err, "GetComposites should return error.")
	assert.Equal(t, "Resource for given ID not found.", err.Error())
}

func TestGetCompositesFailedNoCompositesFound(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(`[]`).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetComposites(context.Background(), "teamID")
	assert.NotNil(t, err, "GetComposites should return error.")
	assert.Equal(t, "No composites found under team ID teamID.", err.Error())

}

func TestGetCompositeByIDSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(CompositeJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetCompositeByID(context.Background(), "teamID", "compositeID")
	assert.Nil(t, err, "GetCompositeByID shouldn't return error.")
}

func TestGetCompositeByIDFailureInvalidTeamID(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(CompositeJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetCompositeByID(context.Background(), "invalidTeamID", "compositeID")
	assert.NotNil(t, err, "GetCompositeByID should return error.")
	assert.Equal(t, "No team with ID invalidTeamID found.", err.Error())

}

func TestGetCompositeByIDFailureInvalidCompositeID(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(CompositeJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetCompositeByID(context.Background(), "teamID", "invalidCompositeID")
	assert.NotNil(t, err, "GetCompositeByID should return error.")
	assert.Equal(t, "No composite with ID invalidCompositeID found under team ID teamID.", err.Error())
}

func TestCreateCompositeSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/composites").WithMethod("POST").WithBody(createdCompositeJSONPayload).WithStatus(http.StatusCreated)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateComposite(context.Background(), "gitURL", "name", "teamID", "gitKeyId")
	assert.Nil(t, err, "CreateComposite shouldn't return error.")
}

func TestCreateCompositeFailureBadRequest(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/composites").WithMethod("POST").WithBody(createdCompositeJSONPayload).WithStatus(http.StatusBadRequest)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateComposite(context.Background(), "gitURL", "name", "teamID", "gitKeyId")
	assert.NotNil(t, err, "CreateComposite should return error.")
}

func TestCreateCompositeFailedToParseUser(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/Composites").WithMethod("POST").WithBody(createdCompositeJSONPayload).WithStatus(http.StatusCreated)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateComposite(context.Background(), "gitURL", "name", "teamID", "gitKeyId")
	assert.NotNil(t, err, "CreateComposite should return error.")
}

func TestCreateCompositeFailedToParseCompositeLink(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/composites").WithMethod("POST").WithBody(createdCompositeJSONPayload).WithStatus(http.StatusCreated)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateComposite(context.Background(), "gitURL", "name", "teamID", "gitKeyId")
	assert.NotNil(t, err, "CreateComposite should return error.")
	assert.Equal(t, "json: cannot unmarshal object into Go value of type []*client.Team", err.Error())
}

func TestCreateCompositesFailureMissingCompositesLink(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/composites").WithMethod("POST").WithBody(createdCompositeJSONPayload).WithStatus(http.StatusCreated)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(teamCompositeJSONPayloadMissingCompositeLink).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateComposite(context.Background(), "gitURL", "name", "teamID", "gitKeyId")
	assert.NotNil(t, err, "CreateComposite should return error.")
	assert.Equal(t, "Resource for given ID not found.", err.Error())
}

func TestCreateCompositeFailureCompositeNotCreated(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/composites").WithMethod("POST").WithBody(`{}`).WithStatus(http.StatusCreated)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateComposite(context.Background(), "gitURL", "name", "teamID", "gitKeyId")
	assert.NotNil(t, err, "CreateComposite should return error.")
	assert.Equal(t, "Failed to create composite under team ID teamID.", err.Error())
}
