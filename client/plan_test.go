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
				"href": "%s/api/composites/583bca6dca5e631017ed46fb/plans"
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
			}
		]
	}
]`

const planJSONPayload = `[{
		"defaultPanelRepo": "git@github.com:CloudCoreo/default-panel.git",
		"defaultPanelDirectory": "panel",
		"defaultPanelBranch": "master",
		"name": "Audit-S3",
		"iamUserAccessKeyId": "AKIAIHBYYNZ44JK535IA",
		"iamUserId": "AIDAJC3FTU33BVH22AJTE",
		"iamUserSecretAccessKey": "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",
		"snsSubscriptionArn": "arn:aws:sns:us-west-1:910887748405:coreo-asi-583bcab7ca5e631017ed470b:f50ac4e8-82a8-4291-a0c3-d9e299f79d8d",
		"sqsArn": "arn:aws:sqs:us-west-1:910887748405:coreo-asi-583bcab7ca5e631017ed470b",
		"sqsUrl": "https://sqs.us-west-1.amazonaws.com/910887748405/coreo-asi-583bcab7ca5e631017ed470b",
		"topicArn": "arn:aws:sns:us-west-1:910887748405:coreo-asi-583bcab7ca5e631017ed470b",
		"defaultRegion": "us-east-1",
		"refreshInterval": 1,
		"revision": "HEAD",
		"branch": "master",
		"enabled": false,
		"links": [
			{
				"ref": "self",
				"method": "GET",
				"href": "%s/api/plans/583bcab7ca5e631017ed470b"
			},
			{
				"ref": "team",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/teams/583bc8dbca5e631017ed46c9"
			},
			{
				"ref": "composite",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/composites/583bca6dca5e631017ed46fb"
			},
			{
				"ref": "cloudAccount",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/cloudaccounts/583bca5eca5e631017ed46f8"
			},
			{
				"ref": "planconfig",
				"method": "GET",
				"href": "%s/api/plans/583bcab7ca5e631017ed470b/planconfig"
			},
			{
				"ref": "runnow",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/plans/583bcab7ca5e631017ed470b/runnow"
			}
		],
		"id": "583bcab7ca5e631017ed470b"
	}]`

const planJSONPayloadMissingLink = `[{
		"defaultPanelRepo": "git@github.com:CloudCoreo/default-panel.git",
		"defaultPanelDirectory": "panel",
		"defaultPanelBranch": "master",
		"name": "Audit-S3",
		"iamUserAccessKeyId": "AKIAIHBYYNZ44JK535IA",
		"iamUserId": "AIDAJC3FTU33BVH22AJTE",
		"iamUserSecretAccessKey": "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",
		"snsSubscriptionArn": "arn:aws:sns:us-west-1:910887748405:coreo-asi-583bcab7ca5e631017ed470b:f50ac4e8-82a8-4291-a0c3-d9e299f79d8d",
		"sqsArn": "arn:aws:sqs:us-west-1:910887748405:coreo-asi-583bcab7ca5e631017ed470b",
		"sqsUrl": "https://sqs.us-west-1.amazonaws.com/910887748405/coreo-asi-583bcab7ca5e631017ed470b",
		"topicArn": "arn:aws:sns:us-west-1:910887748405:coreo-asi-583bcab7ca5e631017ed470b",
		"defaultRegion": "us-east-1",
		"refreshInterval": 1,
		"revision": "HEAD",
		"branch": "master",
		"enabled": false,
		"id": "583bcab7ca5e631017ed470b"
	}]`

const planJSONPayloadPlanEnabled = `{
		"defaultPanelRepo": "git@github.com:CloudCoreo/default-panel.git",
		"defaultPanelDirectory": "panel",
		"defaultPanelBranch": "master",
		"name": "Audit-S3",
		"iamUserAccessKeyId": "AKIAIHBYYNZ44JK535IA",
		"iamUserId": "AIDAJC3FTU33BVH22AJTE",
		"iamUserSecretAccessKey": "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",
		"snsSubscriptionArn": "arn:aws:sns:us-west-1:910887748405:coreo-asi-583bcab7ca5e631017ed470b:f50ac4e8-82a8-4291-a0c3-d9e299f79d8d",
		"sqsArn": "arn:aws:sqs:us-west-1:910887748405:coreo-asi-583bcab7ca5e631017ed470b",
		"sqsUrl": "https://sqs.us-west-1.amazonaws.com/910887748405/coreo-asi-583bcab7ca5e631017ed470b",
		"topicArn": "arn:aws:sns:us-west-1:910887748405:coreo-asi-583bcab7ca5e631017ed470b",
		"defaultRegion": "us-east-1",
		"refreshInterval": 1,
		"revision": "HEAD",
		"branch": "master",
		"enabled": true,
		"id": "583bcab7ca5e631017ed470b"
	}`

const planJSONPayloadPlanDisabled = `{
		"defaultPanelRepo": "git@github.com:CloudCoreo/default-panel.git",
		"defaultPanelDirectory": "panel",
		"defaultPanelBranch": "master",
		"name": "Audit-S3",
		"iamUserAccessKeyId": "AKIAIHBYYNZ44JK535IA",
		"iamUserId": "AIDAJC3FTU33BVH22AJTE",
		"iamUserSecretAccessKey": "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",
		"snsSubscriptionArn": "arn:aws:sns:us-west-1:910887748405:coreo-asi-583bcab7ca5e631017ed470b:f50ac4e8-82a8-4291-a0c3-d9e299f79d8d",
		"sqsArn": "arn:aws:sqs:us-west-1:910887748405:coreo-asi-583bcab7ca5e631017ed470b",
		"sqsUrl": "https://sqs.us-west-1.amazonaws.com/910887748405/coreo-asi-583bcab7ca5e631017ed470b",
		"topicArn": "arn:aws:sns:us-west-1:910887748405:coreo-asi-583bcab7ca5e631017ed470b",
		"defaultRegion": "us-east-1",
		"refreshInterval": 1,
		"revision": "HEAD",
		"branch": "master",
		"enabled": false,
		"id": "583bcab7ca5e631017ed470b"
	}`

func TestGetPlansSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/composites/583bca6dca5e631017ed46fb/plans").WithMethod("GET").WithBody(planJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetPlans(context.Background(), "583bc8dbca5e631017ed46c9", "583bca6dca5e631017ed46fb")
	assert.Nil(t, err, "GetPlans shouldn't return error.")
}

func TestGetPlansFailureInvalidCompositeID(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/composites/583bca6dca5e631017ed46fb/plans").WithMethod("GET").WithBody(planJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetPlans(context.Background(), "583bc8dbca5e631017ed46c9", "invalidCompositeID")
	assert.NotNil(t, err, "GetPlans should return error.")
}

func TestGetPlansFailureInvalidMissingPlanLink(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/composites/583bca6dca5e631017ed46fb/plans").WithMethod("GET").WithBody(planJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(compositeJSONPayloadForPlanMissingPlanLinks).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetPlans(context.Background(), "583bc8dbca5e631017ed46c9", "583bca6dca5e631017ed46fb")
	assert.NotNil(t, err, "GetPlans should return error.")
	assert.Equal(t, "Resource for given ID not found.", err.Error())
}

func TestGetPlansFailureInvalidPlanResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/composites/583bca6dca5e631017ed46fb/plans").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetPlans(context.Background(), "583bc8dbca5e631017ed46c9", "583bca6dca5e631017ed46fb")
	assert.NotNil(t, err, "GetPlans should return error.")
}

func TestGetPlansFailureNoPlansFound(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/composites/583bca6dca5e631017ed46fb/plans").WithMethod("GET").WithBody(`[]`).WithStatus(http.StatusOK)
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetPlans(context.Background(), "583bc8dbca5e631017ed46c9", "583bca6dca5e631017ed46fb")
	assert.NotNil(t, err, "GetPlans should return error.")
	assert.Equal(t, "No plans found under team team ID 583bc8dbca5e631017ed46c9 and composite ID 583bca6dca5e631017ed46fb.", err.Error())
}

func TestGetPlanByIDSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/composites/583bca6dca5e631017ed46fb/plans").WithMethod("GET").WithBody(planJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetPlanByID(context.Background(), "583bc8dbca5e631017ed46c9", "583bca6dca5e631017ed46fb", "583bcab7ca5e631017ed470b")
	assert.Nil(t, err, "GetPlanByID shouldn't return error.")
}

func TestGetPlanByIDFailurePlanNotFound(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/composites/583bca6dca5e631017ed46fb/plans").WithMethod("GET").WithBody(planJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetPlanByID(context.Background(), "583bc8dbca5e631017ed46c9", "583bca6dca5e631017ed46fb", "invalidPlanID")
	assert.NotNil(t, err, "GetPlanByID should return error.")
	assert.Equal(t, "No plan with ID invalidPlanID found under team ID 583bc8dbca5e631017ed46c9 and composite ID 583bca6dca5e631017ed46fb.", err.Error())
}

func TestGetPlanByIDFailureInvalidPlanResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/composites/583bca6dca5e631017ed46fb/plans").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetPlanByID(context.Background(), "583bc8dbca5e631017ed46c9", "583bca6dca5e631017ed46fb", "invalidPlanID")
	assert.NotNil(t, err, "GetPlanByID should return error.")
}

func TestEnablePlanSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/583bcab7ca5e631017ed470b").WithMethod("PUT").WithBody(planJSONPayloadPlanEnabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/583bca6dca5e631017ed46fb/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayload, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.EnablePlan(context.Background(), "583bc8dbca5e631017ed46c9", "583bca6dca5e631017ed46fb", "583bcab7ca5e631017ed470b")
	assert.Nil(t, err, "EnablePlan shouldn't return error.")
}

func TestEnablePlanFailureInvalidPlanResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/583bcab7ca5e631017ed470b").WithMethod("PUT").WithBody(planJSONPayloadPlanEnabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/583bca6dca5e631017ed46fb/plans").WithMethod("GET").WithBody(fmt.Sprintf(`{}`, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.EnablePlan(context.Background(), "583bc8dbca5e631017ed46c9", "583bca6dca5e631017ed46fb", "583bcab7ca5e631017ed470b")
	assert.NotNil(t, err, "EnablePlan should return error.")
}

func TestEnablePlanFailureInvalidPlanID(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/583bcab7ca5e631017ed470b").WithMethod("PUT").WithBody(planJSONPayloadPlanEnabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/583bca6dca5e631017ed46fb/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayload, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.EnablePlan(context.Background(), "583bc8dbca5e631017ed46c9", "583bca6dca5e631017ed46fb", "invalidPlanID")
	assert.NotNil(t, err, "EnablePlan should return error.")
	assert.Equal(t, "Failed to enable plan ID invalidPlanID found under team ID 583bc8dbca5e631017ed46c9 and composite ID 583bca6dca5e631017ed46fb.", err.Error())
}

func TestEnablePlanFailureToUpdatePlan(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/583bcab7ca5e631017ed470b").WithMethod("PUT").WithBody(planJSONPayloadPlanDisabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/583bca6dca5e631017ed46fb/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayload, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.EnablePlan(context.Background(), "583bc8dbca5e631017ed46c9", "583bca6dca5e631017ed46fb", "invalidPlanID")
	assert.NotNil(t, err, "EnablePlan should return error.")
	assert.Equal(t, "Failed to enable plan ID invalidPlanID found under team ID 583bc8dbca5e631017ed46c9 and composite ID 583bca6dca5e631017ed46fb.", err.Error())
}

func TestEnablePlanFailureToUpdatePlanBadRequest(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/583bcab7ca5e631017ed470b").WithMethod("PUT").WithBody(planJSONPayloadPlanDisabled).WithStatus(http.StatusBadRequest)
	ts.Path("/api/composites/583bca6dca5e631017ed46fb/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayload, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.EnablePlan(context.Background(), "583bc8dbca5e631017ed46c9", "583bca6dca5e631017ed46fb", "583bcab7ca5e631017ed470b")
	assert.NotNil(t, err, "EnablePlan should return error.")
}


func TestEnablePlanFailureMissingSelfLinks(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/583bcab7ca5e631017ed470b").WithMethod("PUT").WithBody(planJSONPayloadPlanDisabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/583bca6dca5e631017ed46fb/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayloadMissingLink, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.EnablePlan(context.Background(), "583bc8dbca5e631017ed46c9", "583bca6dca5e631017ed46fb", "583bcab7ca5e631017ed470b")
	assert.NotNil(t, err, "EnablePlan should return error.")
	assert.Equal(t, "Resource for given ID not found.", err.Error())
}

func TestDisablePlanSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/583bcab7ca5e631017ed470b").WithMethod("PUT").WithBody(planJSONPayloadPlanDisabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/583bca6dca5e631017ed46fb/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayload, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.DisablePlan(context.Background(), "583bc8dbca5e631017ed46c9", "583bca6dca5e631017ed46fb", "583bcab7ca5e631017ed470b")
	assert.Nil(t, err, "DisablePlan shouldn't return error.")
}

func TestDisablePlanFailureInvalidPlanResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/583bcab7ca5e631017ed470b").WithMethod("PUT").WithBody(planJSONPayloadPlanDisabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/583bca6dca5e631017ed46fb/plans").WithMethod("GET").WithBody(fmt.Sprintf(`{}`, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.DisablePlan(context.Background(), "583bc8dbca5e631017ed46c9", "583bca6dca5e631017ed46fb", "583bcab7ca5e631017ed470b")
	assert.NotNil(t, err, "DisablePlan should return error.")
}

func TestDisablePlanFailureInvalidPlanID(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/583bcab7ca5e631017ed470b").WithMethod("PUT").WithBody(planJSONPayloadPlanDisabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/583bca6dca5e631017ed46fb/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayload, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.DisablePlan(context.Background(), "583bc8dbca5e631017ed46c9", "583bca6dca5e631017ed46fb", "invalidPlanID")
	assert.NotNil(t, err, "DisablePlan should return error.")
	assert.Equal(t, "Failed to disable plan ID invalidPlanID found under team ID 583bc8dbca5e631017ed46c9 and composite ID 583bca6dca5e631017ed46fb.", err.Error())
}

func TestDisablelanFailureToUpdatePlan(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/583bcab7ca5e631017ed470b").WithMethod("PUT").WithBody(planJSONPayloadPlanEnabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/583bca6dca5e631017ed46fb/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayload, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.DisablePlan(context.Background(), "583bc8dbca5e631017ed46c9", "583bca6dca5e631017ed46fb", "invalidPlanID")
	assert.NotNil(t, err, "DisablePlan should return error.")
	assert.Equal(t, "Failed to disable plan ID invalidPlanID found under team ID 583bc8dbca5e631017ed46c9 and composite ID 583bca6dca5e631017ed46fb.", err.Error())
}

func TestDisablePlanFailureToUpdatePlanBadRequest(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/583bcab7ca5e631017ed470b").WithMethod("PUT").WithBody(planJSONPayloadPlanEnabled).WithStatus(http.StatusBadRequest)
	ts.Path("/api/composites/583bca6dca5e631017ed46fb/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayload, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.DisablePlan(context.Background(), "583bc8dbca5e631017ed46c9", "583bca6dca5e631017ed46fb", "583bcab7ca5e631017ed470b")
	assert.NotNil(t, err, "DisablePlan should return error.")
}


func TestDisablePlanFailureMissingSelfLinks(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/plans/583bcab7ca5e631017ed470b").WithMethod("PUT").WithBody(planJSONPayloadPlanEnabled).WithStatus(http.StatusOK)
	ts.Path("/api/composites/583bca6dca5e631017ed46fb/plans").WithMethod("GET").WithBody(fmt.Sprintf(planJSONPayloadMissingLink, ts.URL, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/teams/583bc8dbca5e631017ed46c9/composites").WithMethod("GET").WithBody(fmt.Sprintf(compositeJSONPayloadForPlan, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCompositeJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.DisablePlan(context.Background(), "583bc8dbca5e631017ed46c9", "583bca6dca5e631017ed46fb", "583bcab7ca5e631017ed470b")
	assert.NotNil(t, err, "DisablePlan should return error.")
	assert.Equal(t, "Resource for given ID not found.", err.Error())
}



