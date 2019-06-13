package client

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"

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
				"href": "%s/cloudaccounts/cloudAccountID"
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
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(CloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/cloudaccounts/cloudAccountID/event/setup", httpmock.NewStringResponder(http.StatusOK, EventConfigureResponse))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	config, err := client.GetSetupConfig(context.Background(), "teamID", "cloudAccountID")
	assert.Nil(t, err, "getSetupConfig shouldn't return error")
	assert.Equal(t, "fakeStackName", config.StackName)
}

func TestGetSetupConfigFailureWithNoResponse(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(CloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/cloudaccounts/cloudAccountID/event/setup", httpmock.NewStringResponder(http.StatusBadRequest, ""))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetSetupConfig(context.Background(), "teamID", "cloudAccountID")
	assert.NotNil(t, err, "getSetupConfig should return error")
}

func TestGetSetupConfigFailureWithNoLink(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(CloudAccountJSONPayloadNoSetup, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetSetupConfig(context.Background(), "teamID", "cloudAccountID")
	assert.NotNil(t, err, "getSetupConfig should return error")
	assert.Equal(t, "resource for given ID not found", err.Error())
}

func TestGetRemoveConfigSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(CloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/cloudaccounts/cloudAccountID/event/remove", httpmock.NewStringResponder(http.StatusOK, RemoveConfigureResponse))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	config, err := client.GetRemoveConfig(context.Background(), "teamID", "cloudAccountID")
	assert.Nil(t, err, "getRemoveConfig shouldn't return error")
	assert.Equal(t, "fakeStackName", config.StackName)
}

func TestGetRemoveConfigFailureWithNoResponse(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(CloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/cloudaccounts/cloudAccountID/event/remove", httpmock.NewStringResponder(http.StatusBadRequest, ""))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetRemoveConfig(context.Background(), "teamID", "cloudAccountID")
	assert.NotNil(t, err, "getRemovepConfig should return error")
}

func TestGetRemoveConfigFailureWithNoLink(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(CloudAccountJSONPayloadNoSetup, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetRemoveConfig(context.Background(), "teamID", "cloudAccountID")
	assert.NotNil(t, err, "getRemoveConfig should return error")
	assert.Equal(t, "resource for given ID not found", err.Error())
}
