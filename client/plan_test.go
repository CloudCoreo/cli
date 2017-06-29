package client

import (
	"net/http"
	"testing"

	"fmt"

	"github.com/jharlap/httpstub"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

const compositeJSONPayloadForPlan = `[
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
				"href": "%s/api/composites/compositeID/plans"
			}
		]
	}
]`

const compositeJSONPayloadForPlanMissingPlanLinks = `[
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
			}
		]
	}
]`

const planJSONPayloadSingleCompletedFailed = `{
		"engineRunInfos": {
			"_id": "new_id",
			"engineState": "COMPLETED",
			"engineStatus": "ERROR",
			"runId": "70991609-1fab-4abf-9846-df1c4c83f4e8_new",
			"numberOfResources": 18,
			"createdAt": "2017-03-06T22:43:18.486Z",
			"engineStateMessage": {
				"error_message": "Some failed"
			}
		},
		"isDraft": true,
		"links": [
			{
				"ref": "self",
				"method": "GET",
				"href": "%s/api/plans/planid"
			},
			{
				"ref": "team",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/teams/teamID"
			},
			{
				"ref": "composite",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/composites/compositeID"
			},
			{
				"ref": "cloudAccount",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/cloudaccounts/cloudAccountID"
			},
			{
				"ref": "planconfig",
				"method": "GET",
				"href": "%s/api/plans/planID/planconfig"
			},
			{
				"ref": "runnow",
				"method": "GET",
				"href": "%s/api/plans/planID/runnow"
			},
			{
				"ref": "panel",
				"method": "GET",
				"href": "%s/api/plans/planID/panel"
			},
			{
				"ref": "compile_now",
				"method": "GET",
				"href": "%s/api/plans/planID/compile_now"
			}
		],
		"id": "planID"
	}`

const planJSONPayloadSingleCompiled = `{
		"engineRunInfos": {
			"_id": "new_id",
			"engineState": "COMPILED",
			"engineStatus": "OK",
			"runId": "70991609-1fab-4abf-9846-df1c4c83f4e8_new",
			"numberOfResources": 18,
			"createdAt": "2017-03-06T22:43:18.486Z"
		},
		"isDraft": true,
		"links": [
			{
				"ref": "self",
				"method": "GET",
				"href": "%s/api/plans/planid"
			},
			{
				"ref": "team",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/teams/teamID"
			},
			{
				"ref": "composite",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/composites/compositeID"
			},
			{
				"ref": "cloudAccount",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/cloudaccounts/cloudAccountID"
			},
			{
				"ref": "planconfig",
				"method": "GET",
				"href": "%s/api/plans/planID/planconfig"
			},
			{
				"ref": "runnow",
				"method": "GET",
				"href": "%s/api/plans/planID/runnow"
			},
			{
				"ref": "panel",
				"method": "GET",
				"href": "%s/api/plans/planID/panel"
			},
			{
				"ref": "compile_now",
				"method": "GET",
				"href": "%s/api/plans/planID/compile_now"
			}
		],
		"id": "planID"
	}`

const planJSONPayloadSingleCompleted = `{
		"engineRunInfos": {
			"_id": "new_id",
			"engineState": "COMPLETED",
			"engineStatus": "OK",
			"runId": "70991609-1fab-4abf-9846-df1c4c83f4e8_new",
			"numberOfResources": 18,
			"createdAt": "2017-03-06T22:43:18.486Z"
		},
		"isDraft": true,
		"links": [
			{
				"ref": "self",
				"method": "GET",
				"href": "%s/api/plans/planid"
			},
			{
				"ref": "team",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/teams/teamID"
			},
			{
				"ref": "composite",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/composites/compositeID"
			},
			{
				"ref": "cloudAccount",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/cloudaccounts/cloudAccountID"
			},
			{
				"ref": "planconfig",
				"method": "GET",
				"href": "%s/api/plans/planID/planconfig"
			},
			{
				"ref": "runnow",
				"method": "GET",
				"href": "%s/api/plans/planID/runnow"
			},
			{
				"ref": "panel",
				"method": "GET",
				"href": "%s/api/plans/planID/panel"
			},
			{
				"ref": "compile_now",
				"method": "GET",
				"href": "%s/api/plans/planID/compile_now"
			}
		],
		"id": "planID"
	}`

const planJSONPayloadSingle = `{
		"engineRunInfos": {
			"_id": "58bde606be8c31b09480f0d9",
			"engineState": "INITIALIZED",
			"engineStatus": "OK",
			"runId": "70991609-1fab-4abf-9846-df1c4c83f4e8",
			"numberOfResources": 18,
			"createdAt": "2017-03-06T22:43:18.486Z"
		},
		"isDraft": true,
		"links": [
			{
				"ref": "self",
				"method": "GET",
				"href": "%s/api/plans/planid"
			},
			{
				"ref": "team",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/teams/teamID"
			},
			{
				"ref": "composite",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/composites/compositeID"
			},
			{
				"ref": "cloudAccount",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/cloudaccounts/cloudAccountID"
			},
			{
				"ref": "planconfig",
				"method": "GET",
				"href": "%s/api/plans/planID/planconfig"
			},
			{
				"ref": "runnow",
				"method": "GET",
				"href": "%s/api/plans/planID/runnow"
			},
			{
				"ref": "panel",
				"method": "GET",
				"href": "%s/api/plans/planID/panel"
			},
			{
				"ref": "compile_now",
				"method": "GET",
				"href": "%s/api/plans/planID/compile_now"
			}
		],
		"id": "planID"
	}`

const planJSONPayload = `[{
		"defaultPanelRepo": "git@github.com:CloudCoreo/default-panel.git",
		"defaultPanelDirectory": "panel",
		"defaultPanelBranch": "master",
		"name": "Audit-S3",
		"iamUserAccessKeyId": "IAMUSERACCESSKEYID",
		"iamUserId": "iamUserId",
		"iamUserSecretAccessKey": "iamUserSecretAccessKey",
		"snsSubscriptionArn": "arn:aws:sns:us-west-1:12312123:coreo-asi-planID:f50ac4e8-82a8-4291-a0c3-d9e299f79d8d",
		"sqsArn": "arn:aws:sqs:us-west-1:12312123:coreo-asi-planID",
		"sqsUrl": "sqsUrl",
		"topicArn": "arn:aws:sns:us-west-1:12312123:coreo-asi-planID",
		"defaultRegion": "us-east-1",
		"refreshInterval": 1,
		"revision": "HEAD",
		"branch": "master",
		"enabled": false,
		"links": [
			{
				"ref": "self",
				"method": "GET",
				"href": "%s/api/plans/planID"
			},
			{
				"ref": "team",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/teams/teamID"
			},
			{
				"ref": "composite",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/composites/compositeID"
			},
			{
				"ref": "cloudAccount",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/cloudaccounts/cloudAccountID"
			},
			{
				"ref": "planconfig",
				"method": "GET",
				"href": "%s/api/plans/planID/planconfig"
			},
			{
				"ref": "panel",
				"method": "GET",
				"href": "%s/api/plans/planID/panel"
			},
			{
				"ref": "runnow",
				"method": "GET",
				"href": "%s/api/plans/planID/runnow"
			}
		],
		"id": "planID"
	}]`

const planJSONPayloadMissingLink = `[{
		"defaultPanelRepo": "git@github.com:CloudCoreo/default-panel.git",
		"defaultPanelDirectory": "panel",
		"defaultPanelBranch": "master",
		"name": "Audit-S3",
		"iamUserAccessKeyId": "IAMUSERACCESSKEYID",
		"iamUserId": "iamUserId",
		"iamUserSecretAccessKey": "iamUserSecretAccessKey",
		"snsSubscriptionArn": "arn:aws:sns:us-west-1:12312123:coreo-asi-planID:f50ac4e8-82a8-4291-a0c3-d9e299f79d8d",
		"sqsArn": "arn:aws:sqs:us-west-1:12312123:coreo-asi-planID",
		"sqsUrl": "sqsUrl",
		"topicArn": "arn:aws:sns:us-west-1:12312123:coreo-asi-planID",
		"defaultRegion": "us-east-1",
		"refreshInterval": 1,
		"revision": "HEAD",
		"branch": "master",
		"enabled": false,
		"id": "planID"
	}]`

const planJSONPayloadPlanEnabled = `{
		"defaultPanelRepo": "git@github.com:CloudCoreo/default-panel.git",
		"defaultPanelDirectory": "panel",
		"defaultPanelBranch": "master",
		"name": "Audit-S3",
		"iamUserAccessKeyId": "IAMUSERACCESSKEYID",
		"iamUserId": "iamUserId",
		"iamUserSecretAccessKey": "iamUserSecretAccessKey",
		"snsSubscriptionArn": "arn:aws:sns:us-west-1:12312123:coreo-asi-planID:f50ac4e8-82a8-4291-a0c3-d9e299f79d8d",
		"sqsArn": "arn:aws:sqs:us-west-1:12312123:coreo-asi-planID",
		"sqsUrl": "sqsUrl",
		"topicArn": "arn:aws:sns:us-west-1:12312123:coreo-asi-planID",
		"defaultRegion": "us-east-1",
		"refreshInterval": 1,
		"revision": "HEAD",
		"branch": "master",
		"enabled": true,
		"id": "planID"
	}`

const panelJSONPayload = `{
	"resourcesArray": [],
	"numberOfResources": 18,
	"planRefreshIntervalInHours": 0.016666666666666666,
	"lastExecutionTime": 1496256593240,
	"engineState": "COMPILED",
	"isEnabled": false
}`

const planJSONPayloadPlanDisabled = `{
		"defaultPanelRepo": "git@github.com:CloudCoreo/default-panel.git",
		"defaultPanelDirectory": "panel",
		"defaultPanelBranch": "master",
		"name": "Audit-S3",
		"iamUserAccessKeyId": "IAMUSERACCESSKEYID",
		"iamUserId": "iamUserId",
		"iamUserSecretAccessKey": "iamUserSecretAccessKey",
		"snsSubscriptionArn": "arn:aws:sns:us-west-1:12312123:coreo-asi-planID:f50ac4e8-82a8-4291-a0c3-d9e299f79d8d",
		"sqsArn": "arn:aws:sqs:us-west-1:12312123:coreo-asi-planID",
		"sqsUrl": "sqsUrl",
		"topicArn": "arn:aws:sns:us-west-1:12312123:coreo-asi-planID",
		"defaultRegion": "us-east-1",
		"refreshInterval": 1,
		"revision": "HEAD",
		"branch": "master",
		"enabled": false,
		"id": "planID"
	}`

const planConfigPayload = `{
  "gitRevision": "HEAD",
  "links": [
	{
		"ref": "self",
		"method": "GET",
		"href": "%s/api/planconfigs/planConfigID"
	},
	{
		"ref": "plan",
		"method": "GET",
		"href": "%s/api/plans/planID"
	}
  ],
  "gitBranch": "master",
  "variables": {
  },
  "planId": "planID",
  "id": "ID"
}`

func TestGetPlansSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(planJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetPlans(context.Background(), "teamID", "compositeID")
	assert.Nil(t, err, "GetPlans shouldn't return error.")
}

func TestGetPlansFailureInvalidCompositeID(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(planJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetPlans(context.Background(), "teamID", "invalidCompositeID")
	assert.NotNil(t, err, "GetPlans should return error.")
}

func TestGetPlansFailureInvalidMissingPlanLink(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(planJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(compositeJSONPayloadForPlanMissingPlanLinks).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetPlans(context.Background(), "teamID", "compositeID")
	assert.NotNil(t, err, "GetPlans should return error.")
	assert.Equal(t, "Resource for given ID not found.", err.Error())
}

func TestGetPlansFailureInvalidPlanResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetPlans(context.Background(), "teamID", "compositeID")
	assert.NotNil(t, err, "GetPlans should return error.")
}

func TestGetPlansFailureNoPlansFound(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(`[]`).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetPlans(context.Background(), "teamID", "compositeID")
	assert.NotNil(t, err, "GetPlans should return error.")
	assert.Equal(t, "No plans found under team team ID teamID and composite ID compositeID.", err.Error())
}

func TestGetPlanByIDSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayloadSingle, ts.URL, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayload, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetPlanByID(context.Background(), "teamID", "compositeID", "planID")
	assert.Nil(t, err, "GetPlanByID shouldn't return error.")
}

func TestGetPlanByIDFailurePlanNotFound(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(planJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetPlanByID(context.Background(), "teamID", "compositeID", "invalidPlanID")
	assert.NotNil(t, err, "GetPlanByID should return error.")
	assert.Equal(t, "No plan with ID invalidPlanID found under team ID teamID and composite ID compositeID.", err.Error())
}

func TestGetPlanByIDFailureInvalidPlanResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetPlanByID(context.Background(), "teamID", "compositeID", "invalidPlanID")
	assert.NotNil(t, err, "GetPlanByID should return error.")
}

func TestEnablePlanSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID").WithMethod("PUT").WithBody(planJSONPayloadPlanEnabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayload, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.EnablePlan(context.Background(), "teamID", "compositeID", "planID")
	assert.Nil(t, err, "EnablePlan shouldn't return error.")
}

func TestEnablePlanFailureInvalidPlanResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID").WithMethod("PUT").WithBody(planJSONPayloadPlanEnabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.EnablePlan(context.Background(), "teamID", "compositeID", "planID")
	assert.NotNil(t, err, "EnablePlan should return error.")
}

func TestEnablePlanFailureInvalidPlanID(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID").WithMethod("PUT").WithBody(planJSONPayloadPlanEnabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayload, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.EnablePlan(context.Background(), "teamID", "compositeID", "invalidPlanID")
	assert.NotNil(t, err, "EnablePlan should return error.")
	assert.Equal(t, "Failed to enable plan ID invalidPlanID found under team ID teamID and composite ID compositeID.", err.Error())
}

func TestEnablePlanFailureToUpdatePlan(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID").WithMethod("PUT").WithBody(planJSONPayloadPlanDisabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayload, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.EnablePlan(context.Background(), "teamID", "compositeID", "invalidPlanID")
	assert.NotNil(t, err, "EnablePlan should return error.")
	assert.Equal(t, "Failed to enable plan ID invalidPlanID found under team ID teamID and composite ID compositeID.", err.Error())
}

func TestEnablePlanFailureToUpdatePlanBadRequest(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID").WithMethod("PUT").WithBody(planJSONPayloadPlanDisabled).WithStatus(http.StatusBadRequest)
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayload, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.EnablePlan(context.Background(), "teamID", "compositeID", "planID")
	assert.NotNil(t, err, "EnablePlan should return error.")
}

func TestEnablePlanFailureMissingSelfLinks(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID").WithMethod("PUT").WithBody(planJSONPayloadPlanDisabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(planJSONPayloadMissingLink).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.EnablePlan(context.Background(), "teamID", "compositeID", "planID")
	assert.NotNil(t, err, "EnablePlan should return error.")
	assert.Equal(t, "Resource for given ID not found.", err.Error())
}

func TestDisablePlanSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID").WithMethod("PUT").WithBody(planJSONPayloadPlanDisabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayload, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.DisablePlan(context.Background(), "teamID", "compositeID", "planID")
	assert.Nil(t, err, "DisablePlan shouldn't return error.")
}

func TestDisablePlanFailureInvalidPlanResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID").WithMethod("PUT").WithBody(planJSONPayloadPlanDisabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.DisablePlan(context.Background(), "teamID", "compositeID", "planID")
	assert.NotNil(t, err, "DisablePlan should return error.")
}

func TestDisablePlanFailureInvalidPlanID(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID").WithMethod("PUT").WithBody(planJSONPayloadPlanDisabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayload, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.DisablePlan(context.Background(), "teamID", "compositeID", "invalidPlanID")
	assert.NotNil(t, err, "DisablePlan should return error.")
	assert.Equal(t, "Failed to disable plan ID invalidPlanID found under team ID teamID and composite ID compositeID.", err.Error())
}

func TestDisablelanFailureToUpdatePlan(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID").WithMethod("PUT").WithBody(planJSONPayloadPlanEnabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayload, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.DisablePlan(context.Background(), "teamID", "compositeID", "invalidPlanID")
	assert.NotNil(t, err, "DisablePlan should return error.")
	assert.Equal(t, "Failed to disable plan ID invalidPlanID found under team ID teamID and composite ID compositeID.", err.Error())
}

func TestDisablePlanFailureToUpdatePlanBadRequest(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID").WithMethod("PUT").WithBody(planJSONPayloadPlanEnabled).WithStatus(http.StatusBadRequest)
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayload, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.DisablePlan(context.Background(), "teamID", "compositeID", "planID")
	assert.NotNil(t, err, "DisablePlan should return error.")
}

func TestDisablePlanFailureMissingSelfLinks(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID").WithMethod("PUT").WithBody(planJSONPayloadPlanEnabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(planJSONPayloadMissingLink).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.DisablePlan(context.Background(), "teamID", "compositeID", "planID")
	assert.NotNil(t, err, "DisablePlan should return error.")
	assert.Equal(t, "Resource for given ID not found.", err.Error())
}

func TestInitPlanSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID/planconfig").WithMethod("GET").WithBody(planConfigPayload).WithStatus(http.StatusOK)
	ts.Path("/api/plans/planid").WithMethod("PUT").WithBody(fmt.Sprintf(planJSONPayloadSingle, ts.URL, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/composites/compositeID/plans").WithMethod("POST").WithBody(fmt.Sprintf(planJSONPayloadSingle, ts.URL, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.InitPlan(context.Background(), "branch", "name", "region", "teamID", "cloudID", "compositeID", "revision", 2)

	assert.Nil(t, err, "Plan init failed")
}

func TestInitPlanFailureRefreshIntervalLessThan2(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID/planconfig").WithMethod("GET").WithBody(planConfigPayload).WithStatus(http.StatusOK)
	ts.Path("/api/plans/planid").WithMethod("PUT").WithBody(fmt.Sprintf(planJSONPayloadSingle, ts.URL, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/composites/compositeID/plans").WithMethod("POST").WithBody(fmt.Sprintf(planJSONPayloadSingle, ts.URL, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.InitPlan(context.Background(), "branch", "name", "region", "teamID", "cloudID", "compositeID", "revision", 1)

	assert.NotNil(t, err, "Returns error due to refresh interval less than 2.")
}

func TestInitPlanFailureRefreshIntervalMoreThan525600(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID/planconfig").WithMethod("GET").WithBody(planConfigPayload).WithStatus(http.StatusOK)
	ts.Path("/api/plans/planid").WithMethod("PUT").WithBody(fmt.Sprintf(planJSONPayloadSingle, ts.URL, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/composites/compositeID/plans").WithMethod("POST").WithBody(fmt.Sprintf(planJSONPayloadSingle, ts.URL, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.InitPlan(context.Background(), "branch", "name", "region", "teamID", "cloudID", "compositeID", "revision", 525601)

	assert.NotNil(t, err, "Returns error due to refresh interval more than 525600.")
}

func TestCreatePlanSuccessCompleted(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planid").WithMethod("PUT").WithBody(fmt.Sprintf(planJSONPayloadSingleCompleted, ts.URL, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/plans/planID").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayloadSingleCompleted, ts.URL, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/plans/planID").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayloadSingle, ts.URL, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/plans/planID").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayloadSingle, ts.URL, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/planconfigs/planConfigID").WithMethod("PUT").WithBody(fmt.Sprintf(planConfigPayload, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreatePlan(context.Background(), []byte(fmt.Sprintf(planConfigPayload, ts.URL, ts.URL)), "", "")
	assert.Nil(t, err, "Plan creation failed")
}

func TestCreatePlanFailedCompleted(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planid").WithMethod("PUT").WithBody(fmt.Sprintf(planJSONPayloadSingleCompleted, ts.URL, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/plans/planID").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayloadSingleCompletedFailed, ts.URL, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/plans/planID").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayloadSingle, ts.URL, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/plans/planID").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayloadSingle, ts.URL, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/planconfigs/planConfigID").WithMethod("PUT").WithBody(fmt.Sprintf(planConfigPayload, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreatePlan(context.Background(), []byte(fmt.Sprintf(planConfigPayload, ts.URL, ts.URL)), "", "")
	assert.NotNil(t, err, "Returns compile failed error.")
}

func TestCreatePlanSuccessCompiled(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planid").WithMethod("PUT").WithBody(fmt.Sprintf(planJSONPayloadSingleCompleted, ts.URL, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/plans/planID").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayloadSingleCompiled, ts.URL, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/plans/planID").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayloadSingle, ts.URL, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/plans/planID").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayloadSingle, ts.URL, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/planconfigs/planConfigID").WithMethod("PUT").WithBody(fmt.Sprintf(planConfigPayload, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreatePlan(context.Background(), []byte(fmt.Sprintf(planConfigPayload, ts.URL, ts.URL)), "", "")
	assert.Nil(t, err, "Plan creation failed")
}

func TestPanelSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planID/panel").WithMethod("GET").WithBody(panelJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/plans/planID").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayloadSingle, ts.URL, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayload, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetPanelInfo(context.Background(), "teamID", "compositeID", "planID")

	assert.Nil(t, err, "Plan planel failed")
}

func TestRunNowPlanByIDSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/planid").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayloadSingleCompleted, ts.URL, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/plans/planID/runnow").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayloadSingle, ts.URL, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/plans/planID").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayloadSingle, ts.URL, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/composites/compositeID/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayload, ts.URL, ts.URL, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.RunNowPlanByID(context.Background(), "teamID", "compositeID", "planID", true)
	assert.Nil(t, err, "RunNowPlanByID shouldn't return error.")
}
