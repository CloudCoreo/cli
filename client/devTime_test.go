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
	"id": "compositeID"
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
