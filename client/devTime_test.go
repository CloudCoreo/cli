package client

import (
	"net/http"
	"testing"

	"fmt"

	"github.com/jharlap/httpstub"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

const teamDevTimeSONPayload = `[{
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
		"ref": "devtime",
		"method": "POST",
		"href": "%s/api/teams/teamID/devtime"
		}
	],
		"id": "teamID"
	}]`

const devTimeJSONPayload = `{
	"context": "context",
	"task": "task",
	"devTimeUrl": "https://2b80f18f-504c-4065-ac24-d722eddc8a0d.devtime.guru.com:9876",
	"links": [
		{
			"ref": "self",
			"method": "GET",
			"href": "http://localhost:3000/api/devtime/59c2c54dfd9b7a5d0c6fa4cc"
		},
		{
			"ref": "team",
			"method": "GET",
			"href": "http://localhost:3000/api/teams/5994c647af67725e01e4cb8f"
		},
		{
			"ref": "start",
			"method": "GET",
			"href": "http://localhost:3000/api/devtime/59c2c54dfd9b7a5d0c6fa4cc/start"
		},
		{
			"ref": "stop",
			"method": "GET",
			"href": "http://localhost:3000/api/devtime/59c2c54dfd9b7a5d0c6fa4cc/stop"
		},
		{
			"ref": "results",
			"method": "GET",
			"href": "http://localhost:3000/api/devtime/59c2c54dfd9b7a5d0c6fa4cc/results"
		},
		{
			"ref": "jobs",
			"method": "GET",
			"href": "http://localhost:3000/api/devtime/59c2c54dfd9b7a5d0c6fa4cc/jobs"
		}
	],
	"devTimeId": "2b80f18f-504c-4065-ac24-d722eddc8a0d"
}`

func TestCreateProxyTaskSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/devtime").WithMethod("POST").WithBody(devTimeJSONPayload).WithStatus(http.StatusCreated)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamDevTimeSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateDevTime(context.Background(), "teamID", "context", "task")
	assert.Nil(t, err, "CreateProxyTask shouldn't return error.")
}
