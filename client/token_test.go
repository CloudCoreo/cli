package client

import (
	"net/http"
	"testing"

	"fmt"

	"github.com/jharlap/httpstub"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

const userJSONPayloadForToken = `{
	"username": "coolguru",
	"email": "nandesh@cloudcoreo.com",
	"gravatarIconUrl": "http://www.gravatar.com/avatar/ff813888fedf42ea2f849b2fba9e9de8",
	"createdAt": "2016-11-26T09:22:40.356Z",
	"defaultTeamId": "583bc8dbca5e631017ed46c9",
	"links": [
		{
			"ref": "tokens",
			"method": "GET",
			"href": "%s/api/users/583bc8dbca5e631017ed46c8/tokens"
		}
	],
	"id": "583bc8dbca5e631017ed46c8"
}`

const userJSONPayloadForTokenMissingTokenData = `{
	"username": "coolguru",
	"email": "nandesh@cloudcoreo.com",
	"gravatarIconUrl": "http://www.gravatar.com/avatar/ff813888fedf42ea2f849b2fba9e9de8",
	"createdAt": "2016-11-26T09:22:40.356Z",
	"defaultTeamId": "583bc8dbca5e631017ed46c9",
	"links": [
		{
			"ref": "self",
			"method": "GET",
			"href": "https://app.cloudcoreo.com/api/users/583bc8dbca5e631017ed46c8"
		},
		{
			"ref": "defaultTeam",
			"method": "GET",
			"href": "https://app.cloudcoreo.com/api/teams/583bc8dbca5e631017ed46c9"
		},
		{
			"ref": "teams",
			"method": "GET",
			"href": "https://app.cloudcoreo.com/api/users/583bc8dbca5e631017ed46c8/teams"
		}
	],
	"id": "583bc8dbca5e631017ed46c8"
}`

const tokenJSONPayload = `[{
		"name": "CLI",
		"description": "cli prod",
		"creationDate": "2016-11-28T22:51:47.81Z",
		"links": [
			{
				"ref": "self",
				"method": "GET",
				"href": "%s/api/tokens/583cb503ca5e631017ed6ac5"
			}
		],
		"id": "583cb503ca5e631017ed6ac5"
	}]`

const tokenJSONPayloadMissingSelfData = `[{
		"name": "CLI",
		"description": "cli prod",
		"creationDate": "2016-11-28T22:51:47.81Z",
		"id": "583cb503ca5e631017ed6ac5"
	}]`

const createdTokenJSONPayload = `{
		"id": "583cb503ca5e631017ed6ac5"
	}`

func TestGetTokensSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/tokens").WithMethod("GET").WithBody(tokenJSONPayload).WithStatus(http.StatusOK)

	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForToken, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetTokens(context.Background())
	assert.Nil(t, err, "GetTokens shouldn't return error.")
}

func TestGetTokensFailureInvalidTokenResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/tokens").WithMethod("GET").WithBody(`[]`).WithStatus(http.StatusOK)

	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForToken, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetTokens(context.Background())
	assert.NotNil(t, err, "GetTokens should return error.")
	assert.Equal(t, "No tokens found. To create a token use `coreo token add [flags]` command.", err.Error())
}

func TestGetTokensFailureInvalidUserResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/tokens").WithMethod("GET").WithBody(tokenJSONPayload).WithStatus(http.StatusOK)

	ts.Path("/me").WithMethod("GET").WithBody(``).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetTokens(context.Background())
	assert.NotNil(t, err, "GetTokens should return error.")
}

func TestGetTokensFailureMissingTokenLinkInResponse(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/tokens").WithMethod("GET").WithBody(``).WithStatus(http.StatusOK)

	ts.Path("/me").WithMethod("GET").WithBody(userJSONPayloadForTokenMissingTokenData).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetTokens(context.Background())
	assert.NotNil(t, err, "GetTokens should return error.")
}

func TestGetTokenByIDSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/tokens").WithMethod("GET").WithBody(tokenJSONPayload).WithStatus(http.StatusOK)

	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForToken, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetTokenByID(context.Background(), "583cb503ca5e631017ed6ac5")
	assert.Nil(t, err, "GetTokenByID shouldn't return error.")
}

func TestGetTokenByIDFailure(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/tokens").WithMethod("GET").WithBody(``).WithStatus(http.StatusOK)

	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForToken, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetTokenByID(context.Background(), "583cb503ca5e631017ed6ac5")
	assert.NotNil(t, err, "GetTokenByID should return error.")
}

func TestGetTokenByIDFailureTeamIDNotFound(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/tokens").WithMethod("GET").WithBody(tokenJSONPayload).WithStatus(http.StatusOK)

	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForToken, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.GetTokenByID(context.Background(), "583cb503ca5e631017ed6ac")
	assert.NotNil(t, err, "GetTokenByID should return error.")
}

func TestCreateTokenSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/tokens").WithMethod("POST").WithBody(createdTokenJSONPayload).WithStatus(http.StatusCreated)

	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForToken, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateToken(context.Background(), "tokenName", "tokenDescription")
	assert.Nil(t, err, "CreateToken shouldn't return error.")
}

func TestCreateTokenFailedToParseUser(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/tokens").WithMethod("POST").WithBody(createdTokenJSONPayload).WithStatus(http.StatusCreated)

	ts.Path("/me").WithMethod("GET").WithBody(``).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateToken(context.Background(), "tokenName", "tokenDescription")
	assert.NotNil(t, err, "CreateToken should return error.")
}

func TestCreateTokenFailedToParseTokenLink(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/tokens").WithMethod("POST").WithBody(createdTokenJSONPayload).WithStatus(http.StatusCreated)

	ts.Path("/me").WithMethod("GET").WithBody(userJSONPayloadForTokenMissingTokenData).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateToken(context.Background(), "tokenName", "tokenDescription")
	assert.NotNil(t, err, "CreateToken should return error.")
	assert.Equal(t, "Resource for given ID not found.", err.Error())
}

func TestCreateTokenFailedToCreateToken(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/tokens").WithMethod("POST").WithBody(`{}`).WithStatus(http.StatusCreated)

	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForToken, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateToken(context.Background(), "tokenName", "tokenDescription")
	assert.NotNil(t, err, "CreateToken should return error.")
	assert.Equal(t, "Failed to create token.", err.Error())
}

func TestCreateTokenFailedToCreateTokenBadRequest(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/tokens").WithMethod("POST").WithBody(`{}`).WithStatus(http.StatusBadRequest)

	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForToken, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.CreateToken(context.Background(), "tokenName", "tokenDescription")
	assert.NotNil(t, err, "CreateToken should return error.")
}

func TestDeleteTokenByIDSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/tokens/*").WithMethod("DELETE").WithStatus(http.StatusOK)

	ts.Path("/api/users/583bc8dbca5e631017ed46c8/tokens").WithMethod("GET").WithBody(fmt.Sprintf(tokenJSONPayload, ts.URL)).WithStatus(http.StatusOK)

	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForToken, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	err := client.DeleteTokenByID(context.Background(), "583cb503ca5e631017ed6ac5")
	assert.Nil(t, err, "DeleteTokenByID shouldn't return error.")
}

func TestDeleteTokenByIDFailedToParseUser(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/tokens/*").WithMethod("DELETE").WithStatus(http.StatusOK)

	ts.Path("/api/users/583bc8dbca5e631017ed46c8/tokens").WithMethod("GET").WithBody(fmt.Sprintf(tokenJSONPayload, ts.URL)).WithStatus(http.StatusOK)

	ts.Path("/me").WithMethod("GET").WithBody(``).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	err := client.DeleteTokenByID(context.Background(), "583cb503ca5e631017ed6ac5")
	assert.NotNil(t, err, "DeleteTokenByID should return error.")
}

func TestDeleteTokenByIDFailedToParseLink(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/tokens/*").WithMethod("DELETE").WithStatus(http.StatusOK)

	ts.Path("/api/users/583bc8dbca5e631017ed46c8/tokens").WithMethod("GET").WithBody(tokenJSONPayloadMissingSelfData).WithStatus(http.StatusOK)

	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForToken, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	err := client.DeleteTokenByID(context.Background(), "583cb503ca5e631017ed6ac5")
	assert.NotNil(t, err, "DeleteTokenByID should return error.")
}

func TestDeleteTokenByIDFailedToDeleteTokenBadRequest(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/tokens/*").WithMethod("DELETE").WithStatus(http.StatusBadRequest)

	ts.Path("/api/users/583bc8dbca5e631017ed46c8/tokens").WithMethod("GET").WithBody(tokenJSONPayload).WithStatus(http.StatusOK)

	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForToken, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	err := client.DeleteTokenByID(context.Background(), "583cb503ca5e631017ed6ac5")
	assert.NotNil(t, err, "DeleteTokenByID should return error.")
}

func TestDeleteTokenByIDFailedToDeleteTokenInvalidTokenID(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/api/users/583bc8dbca5e631017ed46c8/tokens/*").WithMethod("DELETE").WithStatus(http.StatusBadRequest)

	ts.Path("/api/users/583bc8dbca5e631017ed46c8/tokens").WithMethod("GET").WithBody(tokenJSONPayload).WithStatus(http.StatusOK)

	ts.Path("/me").WithMethod("GET").WithBody(fmt.Sprintf(userJSONPayloadForToken, ts.URL)).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	err := client.DeleteTokenByID(context.Background(), "583cb503ca5e631017ed6ac")
	assert.NotNil(t, err, "DeleteTokenByID should return error.")
}
