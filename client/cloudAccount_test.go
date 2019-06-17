package client

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"

	"github.com/CloudCoreo/cli/client/content"

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
		"href": "%s/teams/teamID/cloudaccounts"
		},
		{
		"ref": "defaultid",
		"method": "GET",
		"href":  "%s/teams/teamID/defaultid"
		}
	],
		"id": "teamID"
	}]`

const revalidateRoleJSONPayload = `{
	"isValid": true
	}`

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
				"href": "%s/cloudaccounts/cloudAccountID"
			},
			{
				"ref": "team",
				"method": "GET",
				"href": "https://app.cloudcoreo.com/api/teams/teamID"
			},
			{
				"ref": "setup",
				"method": "GET",
				"href": "%s/cloudaccounts/cloudAccountID/event/setup"
			},
			{
				"ref": "remove",
				"method": "GET",
				"href": "%s/cloudaccounts/cloudAccountID/event/remove"
			},
			{
				"ref": "update",
				"method": "POST",
				"href": "%s/cloudaccounts/cloudAccountID/update"
			},
			{
				"ref": "test",
				"method": "GET",
				"href": "%s/cloudaccounts/cloudAccountID/re-validate"
			}
		],
		"id": "cloudAccountID",
		"email": "testEmail"
	}]`

const RoleCreationInfoJSONPayload = `{
		"accountId": "Fake-aws-account-id",
		"externalId": "Fake-external-id",
		"domain": "fake domain"
	}`

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
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusOK, CloudAccountJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetCloudAccounts(context.Background(), "teamID")
	assert.Nil(t, err, "GetCloudAccounts shouldn't return error.")
}

func TestGetCloudAccountsFailureInvalidUserResponse(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusOK, CloudAccountJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, `{}`))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetCloudAccounts(context.Background(), "teamID")
	assert.NotNil(t, err, "GetCloudAccounts should return error.")
	assert.Equal(t, "resource for given ID not found", err.Error())
}

func TestGetCloudAccountsFailureInvalidTeamResponse(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusOK, CloudAccountJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, `{}`))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetCloudAccounts(context.Background(), "teamID")
	assert.NotNil(t, err, "GetCloudAccounts should return error.")
	assert.Equal(t, "json: cannot unmarshal object into Go value of type []*client.Team", err.Error())
}

func TestGetCloudAccountsFailureInvalidCloudAccountresponse(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusOK, `{}`))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetCloudAccounts(context.Background(), "teamID")
	assert.NotNil(t, err, "GetCloudAccounts should return error.")
	assert.Equal(t, "json: cannot unmarshal object into Go value of type []*client.CloudAccount", err.Error())
}

func TestGetCloudAccountsFailureMissingCloudAccountsLink(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusOK, CloudAccountJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, teamCloudAccountJSONPayloadMissingCloudAccountLink))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetCloudAccounts(context.Background(), "teamID")
	assert.NotNil(t, err, "GetCloudAccounts should return error.")
	assert.Equal(t, "resource for given ID not found", err.Error())
}

func TestGetCloudAccountsFailedNoCloudAccountsFound(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusOK, `[]`))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetCloudAccounts(context.Background(), "teamID")
	assert.NotNil(t, err, "GetCloudAccounts should return error.")
	assert.Equal(t, "No cloud accounts found under team ID teamID.", err.Error())

}

func TestGetCloudAccountByIDSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusOK, CloudAccountJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetCloudAccountByID(context.Background(), "teamID", "cloudAccountID")
	assert.Nil(t, err, "GetCloudAccountByID shouldn't return error.")
}

func TestGetCloudAccountByIDFailureInvalidTeamID(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusOK, CloudAccountJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetCloudAccountByID(context.Background(), "583bc8dbca5e631017ed46c", "cloudAccountID")
	assert.NotNil(t, err, "GetCloudAccountByID should return error.")
	assert.Equal(t, "No cloud accounts found under team ID 583bc8dbca5e631017ed46c.", err.Error())

}

func TestGetCloudAccountByIDFailureInvalidCloudID(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusOK, CloudAccountJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetCloudAccountByID(context.Background(), "teamID", "InvalidcloudAccountID")
	assert.NotNil(t, err, "GetCloudAccountByID should return error.")
	assert.Equal(t, "No cloud account with ID InvalidcloudAccountID found under team ID teamID.", err.Error())
}

func TestSendCloudCreateRequestSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusCreated, createdCloudAccountJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	link := Link{
		Href: defaultAPIEndpoint + "/teams/teamID/cloudaccounts",
	}
	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	input := &sendCloudCreateRequestInput{
		CloudInfo: CloudInfo{
			Name:       "cloudName",
			Arn:        "default",
			ExternalID: "default",
		},

		cloudLink: link,
	}
	_, err := client.sendCloudCreateRequest(context.Background(), input)
	assert.Nil(t, err, "CreateCloudAccount shouldn't return error.")
}

func TestCreateCloudAccountSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusCreated, createdCloudAccountJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.CreateCloudAccount(context.Background(), &CreateCloudAccountInput{
		TeamID:      "teamID",
		RoleArn:     "roleArn",
		Environment: "Product",
	})
	assert.Nil(t, err, "CreateCloudAccount shouldn't return error.")

}

func TestCreateCloudAccountFailureMissingRoleArn(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusCreated, createdCloudAccountJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.CreateCloudAccount(context.Background(), &CreateCloudAccountInput{
		TeamID:   "teamID",
		Provider: "AWS",
	})
	assert.NotNil(t, err, "CreateCloudAccount should return error.")
	assert.Equal(t, content.ErrorMissingRoleInformation, err.Error())

}

func TestCreateCloudAccountFailureBadRequest(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusBadRequest, createdCloudAccountJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	input := &CreateCloudAccountInput{
		TeamID:    "teamID",
		CloudName: "cloudName",
	}
	_, err := client.CreateCloudAccount(context.Background(), input)
	assert.NotNil(t, err, "CreateCloudAccount should return error.")
}

func TestCreateCloudAccountFailedToParseUser(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusCreated, createdCloudAccountJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, `{}`))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	input := &CreateCloudAccountInput{
		TeamID:    "teamID",
		CloudName: "cloudName",
	}
	_, err := client.CreateCloudAccount(context.Background(), input)
	assert.NotNil(t, err, "CreateCloudAccount should return error.")
}

func TestCreateCloudAccountFailedToParseCloudAccountLink(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusCreated, createdCloudAccountJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, `{}`))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	input := &CreateCloudAccountInput{
		TeamID:    "teamID",
		CloudName: "cloudName",
	}
	_, err := client.CreateCloudAccount(context.Background(), input)
	assert.NotNil(t, err, "CreateCloudAccount should return error.")
	assert.Equal(t, "json: cannot unmarshal object into Go value of type []*client.Team", err.Error())
}

func TestCreateCloudAccountsFailureMissingCloudAccountsLink(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusCreated, createdCloudAccountJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, teamCloudAccountJSONPayloadMissingCloudAccountLink))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	input := &CreateCloudAccountInput{
		TeamID:    "teamID",
		CloudName: "cloudName",
	}
	_, err := client.CreateCloudAccount(context.Background(), input)
	assert.NotNil(t, err, "CreateCloudAccount should return error.")
	assert.Equal(t, "resource for given ID not found", err.Error())
}

func TestCreateCloudAccountFailureCloudAccountNotCreated(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusCreated, `{}`))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	input := &CreateCloudAccountInput{
		TeamID:    "teamID",
		CloudName: "cloudName",
		Provider:  "AWS",
	}
	_, err := client.CreateCloudAccount(context.Background(), input)
	assert.NotNil(t, err, "CreateCloudAccount should return error.")
	assert.Equal(t, "Adding cloud account falied, you need to provide either rolearn and external id or new role name", err.Error())
}

func TestDeleteCloudAccountByIDSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", defaultAPIEndpoint+"/cloudaccounts/cloudAccountID", httpmock.NewStringResponder(http.StatusOK, CloudAccountJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(CloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	err := client.DeleteCloudAccountByID(context.Background(), "teamID", "cloudAccountID")
	assert.Nil(t, err, "DeleteCloudAccountByID shouldn't return error.")
}

func TestDeleteCloudAccountByIDFailedToParseUser(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", defaultAPIEndpoint+"/cloudaccounts/cloudAccountID", httpmock.NewStringResponder(http.StatusOK, CloudAccountJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(CloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, ``))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	err := client.DeleteCloudAccountByID(context.Background(), "tokenID", "cloudAccountID")
	assert.NotNil(t, err, "DeleteCloudAccountByID should return error.")
}

func TestDeleteCloudAccountByIDFailedToParseLink(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", defaultAPIEndpoint+"/cloudaccounts/cloudAccountID", httpmock.NewStringResponder(http.StatusOK, CloudAccountJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(CloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, teamCloudAccountJSONPayloadMissingCloudAccountLink))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	err := client.DeleteCloudAccountByID(context.Background(), "tokenID", "cloudAccountID")
	assert.NotNil(t, err, "DeleteCloudAccountByID should return error.")
	assert.Equal(t, "No cloud accounts found under team ID tokenID.", err.Error())
}

func TestDeleteCloudAccountByIDFailedToParseCloudAccountLink(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", defaultAPIEndpoint+"/cloudaccounts/cloudAccountID", httpmock.NewStringResponder(http.StatusOK, CloudAccountJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusOK, `{}`))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, teamCloudAccountJSONPayloadMissingCloudAccountLink))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	err := client.DeleteCloudAccountByID(context.Background(), "tokenID", "cloudAccountID")
	assert.NotNil(t, err, "DeleteCloudAccountByID should return error.")
	assert.Equal(t, "No cloud accounts found under team ID tokenID.", err.Error())
}

func TestDeleteCloudAccountByIDFailureMissingCloudAccountsLink(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", defaultAPIEndpoint+"/cloudaccounts/cloudAccountID", httpmock.NewStringResponder(http.StatusOK, CloudAccountJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusOK, CloudAccountJSONPayloadMissingSelfData))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	err := client.DeleteCloudAccountByID(context.Background(), "teamID", "cloudAccountID")
	assert.NotNil(t, err, "DeleteCloudAccountByID should return error.")
	assert.Equal(t, "resource for given ID not found", err.Error())
}

func TestDeleteCloudAccountByIDAccountFailureBadRequest(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", defaultAPIEndpoint+"/cloudaccounts/cloudAccountID", httpmock.NewStringResponder(http.StatusBadRequest, CloudAccountJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(CloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	err := client.DeleteCloudAccountByID(context.Background(), "teamID", "cloudAccountID")
	assert.NotNil(t, err, "DeleteCloudAccountByID should return error.")
}

func TestDeleteCloudAccountByIDFailureInvalidCloudID(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", defaultAPIEndpoint+"/cloudaccounts/cloudAccountID", httpmock.NewStringResponder(http.StatusBadRequest, CloudAccountJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(CloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	err := client.DeleteCloudAccountByID(context.Background(), "teamID", "InvalidCloudAccountID")
	assert.NotNil(t, err, "DeleteCloudAccountByID should return error.")
	assert.Equal(t, "Failed to delete cloud account with ID InvalidCloudAccountID under team ID teamID.", err.Error())
}

func TestGetRoleCreationInfoFailureNoUserInfo(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/teams/teamID/defaultid", httpmock.NewStringResponder(http.StatusOK, RoleCreationInfoJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, `{}`))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetRoleCreationInfo(context.Background(), &CreateCloudAccountInput{
		TeamID: "teamID",
	})
	assert.NotNil(t, err, "GetRoleCreationInfo should return error.")
}

func TestGetRoleCreationInfoSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/teams/teamID/defaultid", httpmock.NewStringResponder(http.StatusOK, RoleCreationInfoJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetRoleCreationInfo(context.Background(), &CreateCloudAccountInput{
		TeamID: "teamID",
	})
	assert.Nil(t, err, "GetRoleCreationInfo shouldn't return error.")
}

func TestGetRoleCreationInfoFailureNoTeamIDMatch(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/teams/teamID/defaultid", httpmock.NewStringResponder(http.StatusOK, RoleCreationInfoJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetRoleCreationInfo(context.Background(), &CreateCloudAccountInput{
		TeamID: "teamid",
	})
	assert.NotNil(t, err, "GetRoleCreationInfo should return error.")
	assert.Equal(t, "No team id match", err.Error())
}

func TestGetRoleCreationInfoFailureDefaultIdLinkMissing(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, teamCloudAccountJSONPayloadMissingCloudAccountLink))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetRoleCreationInfo(context.Background(), &CreateCloudAccountInput{
		TeamID: "teamID",
	})
	assert.NotNil(t, err, "GetRoleCreationInfo should return error.")
	assert.Equal(t, "resource for given ID not found", err.Error())
}

func TestUpdateCloudAccountSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", defaultAPIEndpoint+"/cloudaccounts/cloudAccountID/update", httpmock.NewStringResponder(http.StatusOK, createdCloudAccountJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(CloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.UpdateCloudAccount(context.Background(), &UpdateCloudAccountInput{CreateCloudAccountInput: CreateCloudAccountInput{TeamID: "teamID"}, CloudID: "cloudAccountID"})
	assert.Nil(t, err, "UpdateCloudAccount shouldn't return error.")
}

func TestReValidateRoleSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/cloudaccounts/cloudAccountID/re-validate", httpmock.NewStringResponder(http.StatusOK, revalidateRoleJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(CloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.ReValidateRole(context.Background(), "teamID", "cloudAccountID")
	assert.Nil(t, err, "ReValidateRole shouldn't return error.")
}

func TestReValidateRoleFaliureWithLinkMissing(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/cloudaccounts/cloudAccountID/re-validate", httpmock.NewStringResponder(http.StatusOK, revalidateRoleJSONPayload))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/teams/teamID/cloudaccounts", httpmock.NewStringResponder(http.StatusOK, CloudAccountJSONPayloadMissingSelfData))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/users/userID/teams", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(teamCloudAccountJSONPayload, defaultAPIEndpoint, defaultAPIEndpoint)))
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/me", httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(userJSONPayloadForTeam, defaultAPIEndpoint)))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.ReValidateRole(context.Background(), "teamID", "cloudAccountID")
	assert.NotNil(t, err, "ReValidateRole should return error.")
	assert.Equal(t, "resource for given ID not found", err.Error())
}
