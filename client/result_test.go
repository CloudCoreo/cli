package client

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/CloudCoreo/cli/cmd/content"

	"github.com/jharlap/httpstub"
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
			"href": "%s/api/users/userID/result"
		}
	],
	"id": "userID"
}`

const resultJSONPayload = `[
		{
			"ref": "rule",
			"method": "GET",
			"href": "%s/api/rule"
		},
		{
			"ref": "object",
			"method": "GET",
			"href": "%s/api/object"
		}]`

const iamInactiveKeyNoRotationRuleOutput = `[
	{
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
		"teams": [
			{
				"name": "username",
				"id": "team-id"
			}
		],
		"accounts": [
			{
				"name": "new-test",
				"id": "account-id"
			}
		],
		"objects": 1528
	}
]`

const kmsKeyRotatesObjectOutput = `[{
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
	}]`

const rolePolicy = `{
	"Version": "2012-10-17",
	"Statement": [
		{
			"Effect": "Allow",
			"Principal": {
				"AWS": "arn:aws:iam::` + "accountID" + `:root"
			},
			"Action": "sts:AssumeRole",
			"Condition": {
				"StringEquals": {
					"sts:ExternalId": "` + "externalID" + `"
				}
			}
		}
	]
}`

func TestGetAllResultRuleSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForResult, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/result").WithMethod("GET").WithBody(fmt.Sprintf(resultJSONPayload, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/rule").WithMethod("GET").WithBody(iamInactiveKeyNoRotationRuleOutput).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.getAllResultRule(context.Background())
	assert.Nil(t, err, "GetAllResultRule shouldn't return error")
}

func TestGetAllResultObjectSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForResult, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/result").WithMethod("GET").WithBody(fmt.Sprintf(resultJSONPayload, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/object").WithMethod("GET").WithBody(kmsKeyRotatesObjectOutput).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.getAllResultObject(context.Background())
	assert.Nil(t, err, "GetAllResultObject shouldn't return error")
}

func TestGetAllResultRuleFailureNoViolatedRule(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForResult, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/result").WithMethod("GET").WithBody(fmt.Sprintf(resultJSONPayload, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/rule").WithMethod("GET").WithBody("[]").WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.getAllResultRule(context.Background())
	assert.NotNil(t, err, "GetAllResultRule should return error.")
	assert.Equal(t, "No violated rule", err.Error())
}

func TestGetAllResultRuleFailureBadRequest(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForResult, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/result").WithMethod("GET").WithBody(fmt.Sprintf(resultJSONPayload, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/rule").WithMethod("GET").WithBody(iamInactiveKeyNoRotationRuleOutput).WithStatus(http.StatusBadRequest)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.getAllResultRule(context.Background())
	assert.NotNil(t, err, "GetAllResultRule should return error.")
}

func TestGetAllResultObjectFailureBadRequest(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForResult, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/result").WithMethod("GET").WithBody(fmt.Sprintf(resultJSONPayload, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/object").WithMethod("GET").WithBody(kmsKeyRotatesObjectOutput).WithStatus(http.StatusBadRequest)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.getAllResultObject(context.Background())
	assert.NotNil(t, err, "GetAllResultObject should return error.")
}

func TestGetResultRuleSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForResult, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/result").WithMethod("GET").WithBody(fmt.Sprintf(resultJSONPayload, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/rule").WithMethod("GET").WithBody(iamInactiveKeyNoRotationRuleOutput).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.ShowResultRule(context.Background(), "team-id", "account-id", "Medium")
	assert.Nil(t, err, "GetResultRule shouldn't return error")
}

func TestShowResultObjectSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForResult, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/result").WithMethod("GET").WithBody(fmt.Sprintf(resultJSONPayload, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/object").WithMethod("GET").WithBody(kmsKeyRotatesObjectOutput).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.ShowResultObject(context.Background(), content.None, content.None, content.None)
	assert.Nil(t, err, "GetResultObject shouldn't return error")
}

func TestGetResultRuleFailureNoViolatedRule(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForResult, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/result").WithMethod("GET").WithBody(fmt.Sprintf(resultJSONPayload, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/rule").WithMethod("GET").WithBody("[]").WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.ShowResultRule(context.Background(), "team-id", "account-id", "Medium")
	assert.NotNil(t, err, "GetResultRule should return error.")
	assert.Equal(t, "No violated rule", err.Error())
}

func TestShowResultObjectFailureNoViolatedObject(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForResult, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/result").WithMethod("GET").WithBody(fmt.Sprintf(resultJSONPayload, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/object").WithMethod("GET").WithBody("[]").WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.ShowResultObject(context.Background(), content.None, content.None, content.None)
	assert.NotNil(t, err, "GetResultObject should return error.")
	assert.Equal(t, "No violated object", err.Error())
}

func TestCreateAssumeRolePolicyDocument(t *testing.T) {
	ts := httpstub.New()
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	res := client.createAssumeRolePolicyDocument("accountID", "externalID")
	assert.Equal(t, rolePolicy, res)
}
