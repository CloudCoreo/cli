// Copyright © 2016 Paul Allen <paul@cloudcoreo.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

// Basic imports
import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"

	"fmt"

	"github.com/CloudCoreo/cli/client/content"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/net/context"
)

// MakeClientTestSuite Define MakeClient test suite
type MakeClientTestSuite struct {
	suite.Suite
	jsonPayload string
}

// MakeClientTestSuite before each test
func (suite *MakeClientTestSuite) SetupTest() {
}

// TestMakeClientWithApiNoneError MakeClient error
func (suite *MakeClientTestSuite) TestMakeClientWithApiNoneError() {
	_, err := MakeClient("None", "endpoint")
	suite.testMakeClientError(err)
}

// TestMakeClientError MakeClient error
func (suite *MakeClientTestSuite) TestMakeClientWithApiEmptyError() {
	_, err := MakeClient("", "endpoint")
	suite.testMakeClientError(err)
}

// TestMakeClient MakeClient valid
func (suite *MakeClientTestSuite) TestMakeClient() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, err := MakeClient("APIkey", "endpoint")
	assert.Nil(suite.T(), err, "MakeClient should not return error for valid ApiKey, secretKey or endpoint")
	assert.NotNil(suite.T(), client.auth, "client.auth is nil")
	assert.NotNil(suite.T(), client.client, "client.client is nil")
	assert.NotNil(suite.T(), client.opts, "client.opts is nil")
	assert.NotNil(suite.T(), client.opts.interceptor, "client.opts.interceptor is nil")
	assert.NotNil(suite.T(), client.endpoint, "client.endpoint is nil")
}

func (suite *MakeClientTestSuite) testMakeClientError(err error) {
	assert.NotNil(suite.T(), err, "MakeClient should return error for invalid ApiKey or endpoint.")
	assert.Contains(suite.T(), err.Error(), content.ErrorMissingAPIOrSecretKey)

}

// TestMakeClientTestSuite Execute TestMakeClientTestSuite test suite
func TestMakeClientTestSuite(t *testing.T) {
	setupTester := new(MakeClientTestSuite)
	suite.Run(t, setupTester)
}

// DoTestSuite Define client Do test suite
type DoTestSuite struct {
	suite.Suite
}

// DoTestSuite before each test
func (suite *DoTestSuite) SetupTest() {
}

func (suite *DoTestSuite) TestDo() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", defaultAPIEndpoint+"/teams", httpmock.NewStringResponder(http.StatusOK, `{
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
		"href": "https://app.cloudcoreo.com/api/teams/teamID/cloudaccounts"
		}
	],
		"id": "teamID"
	}`))
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	client, _ := MakeClient("APIkey", defaultAPIEndpoint)
	team := &Team{}
	err := client.Do(context.Background(), "POST", defaultAPIEndpoint+"/teams", nil, &team)
	assert.Nil(suite.T(), err, "Do shouldn't return error.")
	assert.Equal(suite.T(), "teamID", team.ID)
}

// TestDoTestSuite Execute TestDoTestSuite test suite
func TestDoTestSuite(t *testing.T) {
	setupTester := new(DoTestSuite)
	suite.Run(t, setupTester)
}

func TestBuildRequest(t *testing.T) {

	i := Interceptor(func(req *http.Request) error { return fmt.Errorf("Return error") })
	c := newClient("http://test.com", WithInterceptor(i))
	_, err := c.buildRequest("GET", "http://test.com", nil)

	assert.NotNil(t, err, "buildRequest should return error.")
}
