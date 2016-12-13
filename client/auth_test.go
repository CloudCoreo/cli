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

	"github.com/stretchr/testify/assert"
)

//TestSignRequest Test SignRequest method
func TestSignRequest(t *testing.T) {
	payLoad := fmt.Sprintf(`{"id":"a1","name":"Alice","gender":"f"}`)
	var jsonStr = []byte(payLoad)

	req, _ := http.NewRequest("GET", "", bytes.NewBuffer(jsonStr))

	auth := &Auth{
		APIKey:    "asdf",
		SecretKey: "123",
	}

	auth.SignRequest(req)

	headerDate := req.Header.Get("date")
	headerAuth := req.Header.Get("Authorization")

	assert.Contains(t, headerAuth, "Hmac", "Request Authorization header doesn't contain Hmac keyword.")
	assert.Condition(t, func() bool { return len(headerDate) > 0 }, "Request date header is empty.")
}
