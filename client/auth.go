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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const cspURL = "https://console.cloud.vmware.com"
const cspResource = "/csp/gateway/am/api/auth/api-tokens/authorize"

// Auth struct for API and secret key
type Auth struct {
	RefreshToken string
}

type cspToken struct {
	AccessToken string `json:"access_token"`
}

// SignRequest method to sign all requests
func (a *Auth) SignRequest(req *http.Request) error {
	body := ""
	if req.Body != nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(req.Body)
		body = buf.String()
		req.Body = ioutil.NopCloser(strings.NewReader(body))
		req.ContentLength = int64(len(body))
	}

	cspToken, err := a.getCspAuthToken()
	if err != nil {
		return err
	}
	req.Header.Add("csp-auth-token", cspToken.AccessToken)
	return nil
}

func (a *Auth) getCspAuthToken() (*cspToken, error) {
	cspToken := new(cspToken)

	data := url.Values{}
	data.Set("refresh_token", a.RefreshToken)

	url, _ := url.ParseRequestURI(cspURL)
	url.Path = cspResource

	req, err := http.NewRequest("POST", url.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", strconv.Itoa(len(data.Encode())))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		message := new(bytes.Buffer)
		message.ReadFrom(resp.Body)
		msg := fmt.Sprintf("%s", message.String())
		return nil, NewError(msg)
	}

	err = json.NewDecoder(resp.Body).Decode(cspToken)

	if err != nil {
		return nil, err
	}
	return cspToken, err
}
