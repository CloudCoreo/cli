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

const devTimesJSONPayload = `[{
		"context": "context",
		"task": "task",
		"devTimeUrl": "https://2b80f18f-504c-4065-ac24-d722eddc8a0d.devtime.guru.com:9999",
		"links": [
			{
				"ref": "self",
				"method": "GET",
				"href": "http://localhost:3000/api/devtime/devtimeid"
			},
			{
				"ref": "team",
				"method": "GET",
				"href": "http://localhost:3000/api/teams/devtimeid"
			},
			{
				"ref": "start",
				"method": "GET",
				"href": "%s/api/devtime/devtimeid/start"
			},
			{
				"ref": "stop",
				"method": "GET",
				"href": "%s/api/devtime/devtimeid/stop"
			},
			{
				"ref": "results",
				"method": "GET",
				"href": "%s/api/devtime/devtimeid/results"
			},
			{
				"ref": "status",
				"method": "GET",
				"href": "%s/api/devtime/devtimeid/status"
			}
		],
		"devTimeId": "2b80f18f-504c-4065-ac24-d722eddc8a0d"
	}]`

const devTimeJSONPayload = `{
	"context": "context",
	"task": "task",
	"devTimeUrl": "https://2b80f18f-504c-4065-ac24-d722eddc8a0d.devtime.guru.com:9999",
	"links": [
		{
			"ref": "self",
			"method": "GET",
			"href": "http://localhost:3000/api/devtime/devtimeid"
		},
		{
			"ref": "team",
			"method": "GET",
			"href": "http://localhost:3000/api/teams/devtimeid"
		},
		{
			"ref": "start",
			"method": "GET",
			"href": "%s/api/devtime/devtimeid/start"
		},
		{
			"ref": "stop",
			"method": "GET",
			"href": "%s/api/devtime/devtimeid/stop"
		},
		{
			"ref": "results",
			"method": "GET",
			"href": "%s/api/devtime/devtimeid/results"
		},
		{
			"ref": "status",
			"method": "GET",
			"href": "%s/api/devtime/devtimeid/status"
		}
	],
	"devTimeId": "2b80f18f-504c-4065-ac24-d722eddc8a0d"
}`

func TestCreateDevTimeSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/devtime").WithMethod("POST").WithBody(devTimeJSONPayload).WithStatus(http.StatusCreated)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamDevTimeSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateDevTime(context.Background(), "teamID", "context", "task")
	assert.Nil(t, err, "CreateDevTime shouldn't return error.")
}

func TestDevTimeStartSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/devtime/devtimeid/start").WithMethod("GET").WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/devtime").WithMethod("GET").WithBody(fmt.Sprintf(devTimesJSONPayload, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamDevTimeSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	err := client.StartDevTime(context.Background(), "teamID", "2b80f18f-504c-4065-ac24-d722eddc8a0d")
	assert.Nil(t, err, "StartDevTime shouldn't return error.")
}

func TestDevTimeStopSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/devtime/devtimeid/start").WithMethod("GET").WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/devtime").WithMethod("GET").WithBody(fmt.Sprintf(devTimesJSONPayload, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamDevTimeSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	err := client.StopDevTime(context.Background(), "teamID", "2b80f18f-504c-4065-ac24-d722eddc8a0d")
	assert.Nil(t, err, "StopDevTime shouldn't return error.")
}
