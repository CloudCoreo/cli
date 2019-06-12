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

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"

	"github.com/stretchr/testify/assert"
)

//TestSignRequest Test SignRequest method
func TestSignRequest(t *testing.T) {
	payLoad := fmt.Sprintf(`{"id":"a1","name":"Alice","gender":"f"}`)
	var jsonStr = []byte(payLoad)

	req, _ := http.NewRequest("GET", "", bytes.NewBuffer(jsonStr))

	auth := &Auth{
		RefreshToken: "asdf",
	}
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", cspURL+cspResource, httpmock.NewStringResponder(http.StatusOK, refreshTokenJSONPayload))

	auth.SignRequest(req)

	authToken := req.Header.Get("csp-auth-token")

	assert.Contains(t, authToken, "fake-access-token", "Request Authorization header doesn't contain csp-auth-token.")
}
