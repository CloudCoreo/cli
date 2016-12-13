package client

import (
	"net/http"
	"testing"

	"fmt"

	"github.com/jharlap/httpstub"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

const teamCloudAccountJSONPayload = `[{
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
		"href": "https://app.cloudcoreo.com/api/teams/teamID/composites"
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
		},
		{
		"ref": "cloudAccounts",
		"method": "GET",
		"href": "%s/api/teams/teamID/cloudaccounts"
		}
	],
		"id": "teamID"
	}]`

const teamCloudAccountJSONPayloadMissingCloudAccountLink = `[{
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
		"href": "https://app.cloudcoreo.com/api/teams/teamID/composites"
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

const CloudAccountJSONPayload = `[
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
	}
]`

const CloudAccountJSONPayloadMissingSelfData = `[{
		"teamId": "teamID",
		"name": "aws cloud account",
		"roleId": "asdf",
		"roleName": "CloudCoreoAssumedRole",
		"id": "cloudAccountID"
	}]`

const createdCloudAccountJSONPayload = `{
		"id": "cloudAccountID"
	}`

func TestGetCloudAccountsSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/cloudaccounts").WithMethod("GET").WithBody(CloudAccountJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCloudAccountJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetCloudAccounts(context.Background(), "teamID")
	assert.Nil(t, err, "GetCloudAccounts shouldn't return error.")
}

func TestGetCloudAccountsFailureInvalidUserResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/cloudaccounts").WithMethod("GET").WithBody(CloudAccountJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCloudAccountJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetCloudAccounts(context.Background(), "teamID")
	assert.NotNil(t, err, "GetCloudAccounts should return error.")
	assert.Equal(t, "Resource for given ID not found.", err.Error())
}

func TestGetCloudAccountsFailureInvalidTeamResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/cloudaccounts").WithMethod("GET").WithBody(CloudAccountJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetCloudAccounts(context.Background(), "teamID")
	assert.NotNil(t, err, "GetCloudAccounts should return error.")
	assert.Equal(t, "json: cannot unmarshal object into Go value of type []*client.Team", err.Error())
}

func TestGetCloudAccountsFailureInvalidCloudAccountresponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/cloudaccounts").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCloudAccountJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetCloudAccounts(context.Background(), "teamID")
	assert.NotNil(t, err, "GetCloudAccounts should return error.")
	assert.Equal(t, "json: cannot unmarshal object into Go value of type []*client.CloudAccount", err.Error())
}

func TestGetCloudAccountsFailureMissingCloudAccountsLink(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/cloudaccounts").WithMethod("GET").WithBody(CloudAccountJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(teamCloudAccountJSONPayloadMissingCloudAccountLink).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetCloudAccounts(context.Background(), "teamID")
	assert.NotNil(t, err, "GetCloudAccounts should return error.")
	assert.Equal(t, "Resource for given ID not found.", err.Error())
}

func TestGetCloudAccountsFailedNoCloudAccountsFound(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/cloudaccounts").WithMethod("GET").WithBody(`[]`).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCloudAccountJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetCloudAccounts(context.Background(), "teamID")
	assert.NotNil(t, err, "GetCloudAccounts should return error.")
	assert.Equal(t, "No cloud accounts found under team ID teamID.", err.Error())

}

func TestGetCloudAccountByIDSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/cloudaccounts").WithMethod("GET").WithBody(CloudAccountJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCloudAccountJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetCloudAccountByID(context.Background(), "teamID", "cloudAccountID")
	assert.Nil(t, err, "GetCloudAccountByID shouldn't return error.")
}

func TestGetCloudAccountByIDFailureInvalidTeamID(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/cloudaccounts").WithMethod("GET").WithBody(CloudAccountJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCloudAccountJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetCloudAccountByID(context.Background(), "583bc8dbca5e631017ed46c", "cloudAccountID")
	assert.NotNil(t, err, "GetCloudAccountByID should return error.")
	assert.Equal(t, "No cloud accounts found under team ID 583bc8dbca5e631017ed46c.", err.Error())

}

func TestGetCloudAccountByIDFailureInvalidCloudID(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/cloudaccounts").WithMethod("GET").WithBody(CloudAccountJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCloudAccountJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetCloudAccountByID(context.Background(), "teamID", "InvalidcloudAccountID")
	assert.NotNil(t, err, "GetCloudAccountByID should return error.")
	assert.Equal(t, "No cloud account with ID InvalidcloudAccountID found under team ID teamID.", err.Error())
}

func TestCreateCloudAccountSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/cloudaccounts").WithMethod("POST").WithBody(createdCloudAccountJSONPayload).WithStatus(http.StatusCreated)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCloudAccountJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateCloudAccount(context.Background(), "teamID", "accessKeyID", "secretAccessKey", "cloudName")
	assert.Nil(t, err, "CreateCloudAccount shouldn't return error.")
}

func TestCreateCloudAccountFailureBadRequest(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/cloudaccounts").WithMethod("POST").WithBody(createdCloudAccountJSONPayload).WithStatus(http.StatusBadRequest)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCloudAccountJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateCloudAccount(context.Background(), "teamID", "accessKeyID", "secretAccessKey", "cloudName")
	assert.NotNil(t, err, "CreateCloudAccount should return error.")
}

func TestCreateCloudAccountFailedToParseUser(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/cloudaccounts").WithMethod("POST").WithBody(createdCloudAccountJSONPayload).WithStatus(http.StatusCreated)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCloudAccountJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateCloudAccount(context.Background(), "teamID", "accessKeyID", "secretAccessKey", "cloudName")
	assert.NotNil(t, err, "CreateCloudAccount should return error.")
}

func TestCreateCloudAccountFailedToParseCloudAccountLink(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/cloudaccounts").WithMethod("POST").WithBody(createdCloudAccountJSONPayload).WithStatus(http.StatusCreated)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateCloudAccount(context.Background(), "teamID", "accessKeyID", "secretAccessKey", "cloudName")
	assert.NotNil(t, err, "CreateCloudAccount should return error.")
	assert.Equal(t, "json: cannot unmarshal object into Go value of type []*client.Team", err.Error())
}

func TestCreateCloudAccountsFailureMissingCloudAccountsLink(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/cloudaccounts").WithMethod("POST").WithBody(createdCloudAccountJSONPayload).WithStatus(http.StatusCreated)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(teamCloudAccountJSONPayloadMissingCloudAccountLink).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateCloudAccount(context.Background(), "teamID", "accessKeyID", "secretAccessKey", "cloudName")
	assert.NotNil(t, err, "CreateCloudAccount should return error.")
	assert.Equal(t, "Resource for given ID not found.", err.Error())
}

func TestCreateCloudAccountFailureCloudAccountNotCreated(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/teams/teamID/cloudaccounts").WithMethod("POST").WithBody(`{}`).WithStatus(http.StatusCreated)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCloudAccountJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateCloudAccount(context.Background(), "teamID", "accessKeyID", "secretAccessKey", "cloudName")
	assert.NotNil(t, err, "CreateCloudAccount should return error.")
	assert.Equal(t, "Failed to create cloud account under team ID teamID.", err.Error())
}

func TestDeleteCloudAccountByIDSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/cloudaccounts/cloudAccountID").WithMethod("DELETE").WithBody(CloudAccountJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/cloudaccounts").WithMethod("GET").WithBody(fmt.Sprintf(CloudAccountJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCloudAccountJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	err := client.DeleteCloudAccountByID(context.Background(), "teamID", "cloudAccountID")
	assert.Nil(t, err, "DeleteCloudAccountByID shouldn't return error.")
}

func TestDeleteCloudAccountByIDFailedToParseUser(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/cloudaccounts/cloudAccountID").WithMethod("DELETE").WithBody(CloudAccountJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/cloudaccounts").WithMethod("GET").WithBody(fmt.Sprintf(CloudAccountJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCloudAccountJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(``).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	err := client.DeleteCloudAccountByID(context.Background(), "tokenID", "cloudAccountID")
	assert.NotNil(t, err, "DeleteCloudAccountByID should return error.")
}

func TestDeleteCloudAccountByIDFailedToParseLink(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/cloudaccounts/cloudAccountID").WithMethod("DELETE").WithBody(CloudAccountJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/cloudaccounts").WithMethod("GET").WithBody(fmt.Sprintf(CloudAccountJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(teamCloudAccountJSONPayloadMissingCloudAccountLink).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	err := client.DeleteCloudAccountByID(context.Background(), "tokenID", "cloudAccountID")
	assert.NotNil(t, err, "DeleteCloudAccountByID should return error.")
	assert.Equal(t, "No cloud accounts found under team ID tokenID.", err.Error())
}

func TestDeleteCloudAccountByIDFailedToParseCloudAccountLink(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/cloudaccounts/cloudAccountID").WithMethod("DELETE").WithBody(CloudAccountJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/cloudaccounts").WithMethod("GET").WithBody(`{}`).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(teamCloudAccountJSONPayloadMissingCloudAccountLink).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	err := client.DeleteCloudAccountByID(context.Background(), "tokenID", "cloudAccountID")
	assert.NotNil(t, err, "DeleteCloudAccountByID should return error.")
	assert.Equal(t, "No cloud accounts found under team ID tokenID.", err.Error())
}

func TestDeleteCloudAccountByIDFailureMissingCloudAccountsLink(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/cloudaccounts/cloudAccountID").WithMethod("DELETE").WithBody(CloudAccountJSONPayload).WithStatus(http.StatusOK)
	ts.Path("/api/teams/teamID/cloudaccounts").WithMethod("GET").WithBody(CloudAccountJSONPayloadMissingSelfData).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCloudAccountJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	err := client.DeleteCloudAccountByID(context.Background(), "teamID", "cloudAccountID")
	assert.NotNil(t, err, "DeleteCloudAccountByID should return error.")
	assert.Equal(t, "Resource for given ID not found.", err.Error())
}

func TestDeleteCloudAccountByIDAccountFailureBadRequest(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/cloudaccounts/cloudAccountID").WithMethod("DELETE").WithBody(CloudAccountJSONPayload).WithStatus(http.StatusBadRequest)
	ts.Path("/api/teams/teamID/cloudaccounts").WithMethod("GET").WithBody(fmt.Sprintf(CloudAccountJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCloudAccountJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	err := client.DeleteCloudAccountByID(context.Background(), "teamID", "cloudAccountID")
	assert.NotNil(t, err, "DeleteCloudAccountByID should return error.")
}

func TestDeleteCloudAccountByIDFailureInvalidCloudID(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/cloudaccounts/cloudAccountID").WithMethod("DELETE").WithBody(CloudAccountJSONPayload).WithStatus(http.StatusBadRequest)
	ts.Path("/api/teams/teamID/cloudaccounts").WithMethod("GET").WithBody(fmt.Sprintf(CloudAccountJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/api/users/userID/teams").WithMethod("GET").WithBody(fmt.Sprintf(teamCloudAccountJSONPayload, ts.URL)).WithStatus(http.StatusOK)
	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForTeam, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	err := client.DeleteCloudAccountByID(context.Background(), "teamID", "InvalidCloudAccountID")
	assert.NotNil(t, err, "DeleteCloudAccountByID should return error.")
	assert.Equal(t, "Failed to delete cloud account with ID InvalidCloudAccountID under team ID teamID.", err.Error())
}
