package client

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/jharlap/httpstub"
	"github.com/stretchr/testify/assert"
)

const EventConfigureResponse = `{
"templateURL": "fakeURL",
"topicName": "fakeName",
"stackName": "fakeStackName",
"devtimeQueueArn": "fakeDevtimeQueueArn",
"version": "1",
"monitorRule": "fakeMonitorRule"}`

const RemoveConfigureResponse = `{
"topicName": "fakeName",
"stackName": "fakeStackName",
"arnType": "fakeArnType"}`

const CloudAccountJSONPayloadNoSetup = `[
	{
		"teamId": "teamID",
		"name": "aws cloud account",
		"roleId": "asdf",
		"roleName": "CloudCoreoAssumedRole",
		"links": [
			{
				"ref": "self",
				"method": "GET",
				"href": "%s/api/cloudaccounts/cloudAccountID"
			},
			{
				"ref": "team",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/teams/teamID"
			}
		],
		"id": "cloudAccountID"
	}]`

func TestGetSetupConfigSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/cloudaccounts").WithMethod("GET").WithBody(fmt.Sprintf(CloudAccountJSONPayload, ts.URL, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCloudAccountJSONPayload, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/cloudaccounts/cloudAccountID/event/setup").WithMethod("GET").WithBody(EventConfigureResponse).WithStatus(http.StatusOK)
	defer ts.Close()
	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	config, err := client.GetSetupConfig(context.Background(), "teamID", "cloudAccountID")
	assert.Nil(t, err, "getSetupConfig shouldn't return error")
	assert.Equal(t, "fakeStackName", config.StackName)
}

func TestGetSetupConfigFailureWithNoResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/cloudaccounts").WithMethod("GET").WithBody(fmt.Sprintf(CloudAccountJSONPayload, ts.URL, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCloudAccountJSONPayload, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/cloudaccounts/cloudAccountID/event/setup").WithMethod("GET").WithBody("").WithStatus(http.StatusBadRequest)
	defer ts.Close()
	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetSetupConfig(context.Background(), "teamID", "cloudAccountID")
	assert.NotNil(t, err, "getSetupConfig should return error")
}

func TestGetSetupConfigFailureWithNoLink(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/cloudaccounts").WithMethod("GET").WithBody(fmt.Sprintf(CloudAccountJSONPayloadNoSetup, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCloudAccountJSONPayload, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetSetupConfig(context.Background(), "teamID", "cloudAccountID")
	assert.NotNil(t, err, "getSetupConfig should return error")
	assert.Equal(t, "resource for given ID not found", err.Error())
}

func TestGetRemoveConfigSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/cloudaccounts").WithMethod("GET").WithBody(fmt.Sprintf(CloudAccountJSONPayload, ts.URL, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCloudAccountJSONPayload, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/cloudaccounts/cloudAccountID/event/remove").WithMethod("GET").WithBody(RemoveConfigureResponse).WithStatus(http.StatusOK)

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	config, err := client.GetRemoveConfig(context.Background(), "teamID", "cloudAccountID")
	assert.Nil(t, err, "getRemoveConfig shouldn't return error")
	assert.Equal(t, "fakeStackName", config.StackName)
}

func TestGetRemoveConfigFailureWithNoResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/cloudaccounts").WithMethod("GET").WithBody(fmt.Sprintf(CloudAccountJSONPayload, ts.URL, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCloudAccountJSONPayload, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/cloudaccounts/cloudAccountID/event/remove").WithMethod("GET").WithBody("").WithStatus(http.StatusBadRequest)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetRemoveConfig(context.Background(), "teamID", "cloudAccountID")
	assert.NotNil(t, err, "getRemovepConfig should return error")
}

func TestGetRemoveConfigFailureWithNoLink(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/cloudaccounts").WithMethod("GET").WithBody(fmt.Sprintf(CloudAccountJSONPayloadNoSetup, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCloudAccountJSONPayload, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetRemoveConfig(context.Background(), "teamID", "cloudAccountID")
	assert.NotNil(t, err, "getRemoveConfig should return error")
	assert.Equal(t, "resource for given ID not found", err.Error())
}
