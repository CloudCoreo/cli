package client

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"

	"github.com/CloudCoreo/cli/client/content"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

const defaultAPIEndpoint = "https://app.hack.securestate.vmware.com/api"

const refreshTokenJSONPayload = `{
    "id_token": "fake-id",
    "token_type": "bearer",
    "expires_in": 1799,
    "scope": "ALL_PERMISSIONS customer_number openid",
    "access_token": "fake-access-token",
    "refresh_token": "fake-refresh-token"
}`

const revalidateRoleJSONPayload = `{
	"isValid": true
	}`

const CloudAccountsJSONPayload = `[
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
				"method": "PUT",
				"href": "%s/cloudaccounts/cloudAccountID"
			},
			{
				"ref": "test",
				"method": "GET",
				"href": "%s/cloudaccounts/cloudAccountID/re-validate"
			}
		],
		"_id": "cloudAccountID",
		"email": "testEmail"
	}]`

const RoleCreationInfoJSONPayload = `{
		"accountId": "Fake-aws-account-id",
		"externalId": "Fake-external-id",
		"domain": "fake domain"	
	}`

const CloudAccountJSONPayload = `{
		"teamId": "teamID",
		"name": "aws cloud account",
		"roleId": "asdf",
		"roleName": "CloudCoreoAssumedRole",
		"_id": "cloudAccountID"
	}`

const createdCloudAccountJSONPayload = `{
		"_id": "cloudAccountID"
	}`

func TestGetCloudAccountsSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/cloudaccounts", httpmock.NewStringResponder(http.StatusOK, CloudAccountsJSONPayload))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetCloudAccounts(context.Background())
	assert.Nil(t, err, "GetCloudAccounts shouldn't return error.")
}

func TestGetCloudAccountsFailureInvalidCloudAccountresponse(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/cloudaccounts", httpmock.NewStringResponder(http.StatusOK, `{}`))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetCloudAccounts(context.Background())
	assert.NotNil(t, err, "GetCloudAccounts should return error.")
	assert.Equal(t, "json: cannot unmarshal object into Go value of type []*client.CloudAccount", err.Error())
}

func TestGetCloudAccountsFailedNoCloudAccountsFound(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/cloudaccounts", httpmock.NewStringResponder(http.StatusOK, `[]`))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetCloudAccounts(context.Background())
	assert.NotNil(t, err, "GetCloudAccounts should return error.")
	assert.Equal(t, "No cloud accounts found.", err.Error())

}

func TestGetCloudAccountByIDSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/cloudaccounts/cloudAccountID", httpmock.NewStringResponder(http.StatusOK, CloudAccountJSONPayload))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetCloudAccountByID(context.Background(), "cloudAccountID")
	assert.Nil(t, err, "GetCloudAccountByID shouldn't return error.")
}

func TestGetCloudAccountByIDFailureInvalidCloudID(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/cloudaccounts/InvalidcloudAccountID", httpmock.NewStringResponder(http.StatusOK, "{}"))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetCloudAccountByID(context.Background(), "InvalidcloudAccountID")
	assert.NotNil(t, err, "GetCloudAccountByID should return error.")
	assert.Equal(t, "No cloud account with ID InvalidcloudAccountID found.", err.Error())
}

func TestSendCloudCreateRequestSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", defaultAPIEndpoint+"/cloudaccounts", httpmock.NewStringResponder(http.StatusCreated, createdCloudAccountJSONPayload))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	input := &CloudInfo{
		Name:       "cloudName",
		Arn:        "default",
		ExternalID: "default",
	}

	_, err := client.sendCloudCreateRequest(context.Background(), input)
	assert.Nil(t, err, "CreateCloudAccount shouldn't return error.")
}

func TestCreateCloudAccountSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", defaultAPIEndpoint+"/cloudaccounts", httpmock.NewStringResponder(http.StatusCreated, createdCloudAccountJSONPayload))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.CreateCloudAccount(context.Background(), &CreateCloudAccountInput{
		RoleArn:     "roleArn",
		Environment: "Product",
		Provider:    "AWS",
	})
	assert.Nil(t, err, "CreateCloudAccount shouldn't return error.")

}

func TestCreateCloudAccountFailureMissingRoleArn(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", defaultAPIEndpoint+"/cloudaccounts", httpmock.NewStringResponder(http.StatusCreated, createdCloudAccountJSONPayload))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.CreateCloudAccount(context.Background(), &CreateCloudAccountInput{
		Provider: "AWS",
	})
	assert.NotNil(t, err, "CreateCloudAccount should return error.")
	assert.Equal(t, content.ErrorMissingRoleInformation, err.Error())

}

func TestCreateCloudAccountFailureBadRequest(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", defaultAPIEndpoint+"/cloudaccounts", httpmock.NewStringResponder(http.StatusBadRequest, createdCloudAccountJSONPayload))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	input := &CreateCloudAccountInput{
		CloudName: "cloudName",
	}
	_, err := client.CreateCloudAccount(context.Background(), input)
	assert.NotNil(t, err, "CreateCloudAccount should return error.")
}

func TestCreateCloudAccountFailureCloudAccountNotCreated(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", defaultAPIEndpoint+"cloudaccounts", httpmock.NewStringResponder(http.StatusCreated, `{}`))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	input := &CreateCloudAccountInput{
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
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	err := client.DeleteCloudAccountByID(context.Background(), "cloudAccountID")
	assert.Nil(t, err, "DeleteCloudAccountByID shouldn't return error.")
}

func TestDeleteCloudAccountByIDAccountFailureBadRequest(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", defaultAPIEndpoint+"/cloudaccounts/cloudAccountID", httpmock.NewStringResponder(http.StatusBadRequest, CloudAccountsJSONPayload))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	err := client.DeleteCloudAccountByID(context.Background(), "cloudAccountID")
	assert.NotNil(t, err, "DeleteCloudAccountByID should return error.")
}

func TestGetRoleCreationInfoSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/.well-known/vss-configuration", httpmock.NewStringResponder(http.StatusOK, RoleCreationInfoJSONPayload))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetRoleCreationInfo(context.Background(), &CreateCloudAccountInput{})
	assert.Nil(t, err, "GetRoleCreationInfo shouldn't return error.")
}

func TestGetRoleCreationInfoFailure(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/.well-known/vss-configuration", httpmock.NewStringResponder(http.StatusBadRequest, ""))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.GetRoleCreationInfo(context.Background(), &CreateCloudAccountInput{})
	assert.NotNil(t, err, "GetRoleCreationInfo should return error.")
}

func TestUpdateCloudAccountSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/cloudaccounts/cloudAccountID", httpmock.NewStringResponder(http.StatusOK, createdCloudAccountJSONPayload))
	httpmock.RegisterResponder("POST", defaultAPIEndpoint+"/cloudaccounts/cloudAccountID/update", httpmock.NewStringResponder(http.StatusOK, createdCloudAccountJSONPayload))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.UpdateCloudAccount(context.Background(), &UpdateCloudAccountInput{CreateCloudAccountInput: CreateCloudAccountInput{}, CloudID: "cloudAccountID"})
	assert.Nil(t, err, "UpdateCloudAccount shouldn't return error.")
}

func TestReValidateRoleSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", defaultAPIEndpoint+"/cloudaccounts/cloudAccountID/re-validate", httpmock.NewStringResponder(http.StatusOK, revalidateRoleJSONPayload))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("ApiKey", defaultAPIEndpoint)
	_, err := client.ReValidateRole(context.Background(), "cloudAccountID")
	assert.Nil(t, err, "ReValidateRole shouldn't return error.")
}
