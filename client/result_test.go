package client

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"

	"github.com/CloudCoreo/cli/cmd/content"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

const userJSONPayloadForResult = `{
	"username": "gitUser",
	"email": "user@cloudcoreo.com",
	"gravatarIconUrl": "http://www.gravatar.com/avatar/avatarID",
	"createdAt": "2016-11-26T09:22:40.356Z",
	"defaultTeamId": "teamID",
	"links": [
		{
			"ref": "result",
			"method": "GET",
			"href": "%s/users/userID/result"
		}
	],
	"id": "userID"
}`

const resultJSONPayload = `[
		{
			"ref": "rule",
			"method": "GET",
			"href": "%s/rule"
		},
		{
			"ref": "object",
			"method": "POST",
			"href": "%s/object"
		}]`

const iamInactiveKeyNoRotationRuleOutput = `{
"result": {"violatingRules": [{
		"id": "fake-rule-name",
		"info": {
			"suggested_action": "fake-suggestion",
			"link": "fake-link",
			"description": "fake-description",
			"display_name": "fake-display-name",
			"level": "Medium",
			"service": "iam",
			"name": "fake-name",
			"region": "global",
			"include_violations_in_count": true,
			"timestamp": "2018-10-11T17:21:54.448+00:00"
		},
		"teamAndPlan": [
			{
				"team": {
					"name": "fake-name",
					"id": "team-id"
				}
			}
		],
		"accounts": [
			"account-id"
		],
		"objects": 1528
	}
]}}`

const kmsKeyRotatesObjectOutput = `{
	"violations": [{
		"id": "fakeObjectId",
		"rule_report": {
			"suggested_action": "fake action",
			"link": "fake-link",
			"description": "fake-description",
			"display_name": "fake-display-name",
			"level": "Medium",
			"service": "kms",
			"name": "fake-name",
			"region": "us-east-1",
			"include_violations_in_count": true,
			"timestamp": "2018-10-11T17:21:55.111+00:00"
		},
		"team": {
			"name": "username",
			"id": "team-id"
		},
		"cloud_account": {
			"name": "new-test",
			"id": "account-id"
		},
		"run_id": "run-id"
	}],
	"totalItems": 10000}`

const NoObjectOutput = `{
	"totalItems": 10000}`

func TestGetAllResultRuleSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForResult, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/result", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(resultJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", defaultAPIEndpoint+"/rule", httpmock.NewStringResponder(http.StatusOK, iamInactiveKeyNoRotationRuleOutput))
	httpmock.RegisterResponder("POST", cspUrl+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	var request = &resultRuleRequest{}
	request.Filter = client.buildFilter("team-id", "account-id", "", "")
	_, err := client.getAllResultRule(context.Background(), request)
	assert.Nil(t, err, "GetAllResultRule shouldn't return error")
}

func TestGetAllResultObjectSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForResult, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/result", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(resultJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", defaultAPIEndpoint+"/object", httpmock.NewStringResponder(http.StatusOK, kmsKeyRotatesObjectOutput))
	httpmock.RegisterResponder("POST", cspUrl+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.getResultObjects(context.Background(), content.None, content.None, "", "", 0)
	assert.Nil(t, err, "GetAllResultObject shouldn't return error")
}

func TestGetAllResultRuleFailureNoViolatedRule(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForResult, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/result", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(resultJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", defaultAPIEndpoint+"/rule", httpmock.NewStringResponder(http.StatusOK, `{"rules":[]}`))
	httpmock.RegisterResponder("POST", cspUrl+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	var request = &resultRuleRequest{}
	request.Filter = client.buildFilter("team-id", "account-id", "", "")
	_, err := client.getAllResultRule(context.Background(), request)
	assert.NotNil(t, err, "GetAllResultRule should return error.")
	assert.Equal(t, "No violated rule", err.Error())
}

func TestGetAllResultRuleFailureBadRequest(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForResult, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/result", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(resultJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", defaultAPIEndpoint+"/rule", httpmock.NewStringResponder(http.StatusBadRequest, iamInactiveKeyNoRotationRuleOutput))
	httpmock.RegisterResponder("POST", cspUrl+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	var request = &resultRuleRequest{}
	request.Filter = client.buildFilter("team-id", "None", "", "")

	_, err := client.getAllResultRule(context.Background(), request)
	assert.NotNil(t, err, "GetAllResultRule should return error.")
}

func TestGetAllResultObjectFailureBadRequest(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForResult, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/result", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(resultJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", defaultAPIEndpoint+"/object", httpmock.NewStringResponder(http.StatusBadRequest, kmsKeyRotatesObjectOutput))
	httpmock.RegisterResponder("POST", cspUrl+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.getResultObjects(context.Background(), content.None, content.None, "", "", 0)
	assert.NotNil(t, err, "GetAllResultObject should return error.")
}

func TestGetResultRuleSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForResult, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/result", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(resultJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", defaultAPIEndpoint+"/rule", httpmock.NewStringResponder(http.StatusOK, iamInactiveKeyNoRotationRuleOutput))
	httpmock.RegisterResponder("POST", cspUrl+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.ShowResultRule(context.Background(), "team-id", "account-id", "Medium", "")
	assert.Nil(t, err, "GetResultRule shouldn't return error")
}

func TestShowResultObjectSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForResult, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/result", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(resultJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", defaultAPIEndpoint+"/object", httpmock.NewStringResponder(http.StatusOK, kmsKeyRotatesObjectOutput))
	httpmock.RegisterResponder("POST", cspUrl+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.ShowResultObject(context.Background(), "teamID", "cloudID", "", "", 0)
	assert.Nil(t, err, "GetResultObject shouldn't return error")
}

func TestGetResultRuleFailureNoViolatedRule(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForResult, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/result", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(resultJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", defaultAPIEndpoint+"/rule", httpmock.NewStringResponder(http.StatusOK, `{"rules":[]}`))
	httpmock.RegisterResponder("POST", cspUrl+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.ShowResultRule(context.Background(), "team-id", "account-id", "Medium", "")
	assert.NotNil(t, err, "GetResultRule should return error.")
	assert.Equal(t, "No violated rule", err.Error())
}

func TestShowResultObjectFailureNoViolatedObject(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForResult, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/result", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(resultJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", defaultAPIEndpoint+"/object", httpmock.NewStringResponder(http.StatusOK, NoObjectOutput))
	httpmock.RegisterResponder("POST", cspUrl+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.ShowResultObject(context.Background(), "teamID", "cloudID", "", "", 0)
	assert.Nil(t, err, "GetResultObject should not return error.")
}
