// Copyright © 2018 Zechen Jiang <zechen@cloudcoreo.com>
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
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/CloudCoreo/cli/cmd/content"
)

//TeamInfo records the info of a team
type TeamInfo struct {
	Name string `json:"name"`
	ID   string `json:"teamId"`
}

//TeamInfoWrapper is a wrapper for team Info
type TeamInfoWrapper struct {
	TeamInfo *TeamInfo `json:"team"`
}

//Info is the struct for rule_report
type Info struct {
	SuggestedAction          string `json:"suggestedAction"`
	Link                     string `json:"link"`
	Description              string `json:"description"`
	DisplayName              string `json:"displayName"`
	Level                    string `json:"level"`
	Service                  string `json:"service"`
	Name                     string `json:"name"`
	IncludeViolationsInCount bool   `json:"include_violations_in_count"`
	TimeStamp                string `json:"lastUpdateTime,omitempty"`
}

// ResultRule struct decodes json file returned by webapp
type ResultRule struct {
	ID      string            `json:"id"`
	Info    Info              `json:"info"`
	TInfo   []TeamInfoWrapper `json:"teamAndPlan"`
	CInfo   []string          `json:"accounts"`
	Object  int               `json:"objects"`
	Regions []string          `json:"regions"`
}

// The ResultObject struct decodes json file returned by webapp
type ResultObject struct {
	ID        string   `json:"id"`
	Info      Info     `json:"ruleInfo"`
	TInfo     TeamInfo `json:"team"`
	RiskScore int      `json:"riskScore"`
	Region    string   `json:"region"`
}

// ResultObjectWrapper contains an object array and number of total items
type ResultObjectWrapper struct {
	AccountName   string          `json:"accountName,omitempty"`
	AccountNumber string          `json:"accountNumber,omitempty"`
	TotalItems    int             `json:"totalCount"`
	Objects       []*ResultObject `json:"violations"`
	ScrollID      string          `json:"continuationToken,omitempty"`
}

//ResultRuleWrapper is wrapper result for violation by rule
type ResultRuleWrapper struct {
	ViolatingRules ViolatingRules `json:"result"`
}

//ViolatingRules contains rules
type ViolatingRules struct {
	Rules []*ResultRule `json:"violatingRules"`
}

type resultObjectRequest struct {
	RemoveScrollID bool   `json:"removeScrollId"`
	ScrollID       string `json:"scrollId,omitempty"`
	Filter         filter `json:"filter"`
}

type resultRuleRequest struct {
	Filter filter `json:"filter"`
}

type filter struct {
	CloudAccounts []string `json:"cloudAccounts,omitempty"`
	Levels        []string `json:"levels,omitempty"`
	Providers     []string `json:"providers,omitempty"`
	Teams         []string `json:"teams,omitempty"`
}

//ShowResultObject shows violated objects. If the filter condition (teamID, cloudID in this case) is valid,
//objects will be filtered. Otherwise return all violation objects under this user account.
func (c *Client) ShowResultObject(ctx context.Context, teamID, cloudID, level, provider string, retry uint) ([]*ResultObjectWrapper, error) {
	var res []*ResultObjectWrapper
	if cloudID != content.None {
		result, err := c.getResultObjects(ctx, teamID, cloudID, level, provider, retry)
		if err != nil {
			return nil, err
		}
		return []*ResultObjectWrapper{result}, nil
	}

	accounts, err := c.GetCloudAccounts(ctx, teamID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	for _, account := range accounts {
		if account.IsDraft {
			continue
		}
		result, err := c.getResultObjects(ctx, teamID, account.ID, level, provider, retry)
		if err != nil {
			return nil, err
		}
		res = append(res, result)
	}

	return res, nil
}

//ShowResultRule show violated rules. If the filter condition (teamID, cloudID in this case) is valid,
//rules will be filtered. Otherwise return all violation rules under this user account.
func (c *Client) ShowResultRule(ctx context.Context, teamID, cloudID, level, provider string) ([]*ResultRule, error) {
	var request = &resultRuleRequest{}
	request.Filter = c.buildFilter(teamID, cloudID, level, provider)
	result, err := c.getAllResultRule(ctx, request)
	if err != nil {
		return nil, NewError(err.Error())
	}
	return result, nil
}

func (c *Client) getResultLinks(ctx context.Context) ([]Link, error) {
	u, err := c.GetUser(ctx)
	if err != nil {
		return nil, err
	}

	resultLink, err := GetLinkByRef(u.Links, "result")
	if err != nil {
		return nil, err
	}

	link := []Link{}
	err = c.Do(ctx, "GET", resultLink.Href, nil, &link)
	if err != nil {
		return nil, err
	}

	return link, nil
}

func (c *Client) getResultLinkByRef(ctx context.Context, ref string) (*Link, error) {
	links, err := c.getResultLinks(ctx)
	if err != nil {
		return nil, err
	}

	link, err := GetLinkByRef(links, ref)
	if err != nil {
		return nil, err
	}

	return &link, err
}

func (c *Client) getResultObjects(ctx context.Context, teamID, cloudID, level, provider string, retry uint) (*ResultObjectWrapper, error) {
	link, err := c.getResultLinkByRef(ctx, "object")
	if err != nil {
		return nil, err
	}
	// buffer used to store objects
	buf := make([]*ResultObject, 0, 200)
	res := make([]*ResultObject, 0)

	var scrollID string
	var cur = 0
	var firstCall = true
	var totalItems = 0
	var request *resultObjectRequest

	for firstCall || cur < totalItems {
		if firstCall {
			request = c.buildGetResultObjectsRequest(teamID, cloudID, level, scrollID, provider, true)
		}
		var err error
		var tmp *ResultObjectWrapper
		for try := uint(0); try <= retry; try++ {
			tmp, err = c.getResultObjectsPaginated(ctx, request, link.Href, buf)
			if err == nil {
				break
			}
			if err != nil && try == retry {
				return nil, err
			}
		}
		if tmp == nil {
			return nil, NewError("No violation object")
		}
		if firstCall {
			totalItems = tmp.TotalItems
			firstCall = false
		}

		res = append(res, tmp.Objects...)
		if tmp.ScrollID == "" {
			break
		}
		scrollID = tmp.ScrollID
		request = c.buildGetResultObjectsRequest(teamID, cloudID, level, scrollID, provider, false)
		cur += len(tmp.Objects)
	}
	wrapper := &ResultObjectWrapper{Objects: res, TotalItems: len(res)}
	if teamID != content.None {
		account, err := c.GetCloudAccountByID(ctx, teamID, cloudID)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
		if account != nil {
			wrapper.AccountNumber = account.AccountID
			wrapper.AccountName = account.Name
		}
	}
	return wrapper, nil
}

func (c *Client) buildGetResultObjectsRequest(teamID, cloudID, level, scrollID, provider string, removeScrollID bool) *resultObjectRequest {
	request := resultObjectRequest{
		RemoveScrollID: removeScrollID,
		ScrollID:       scrollID,
		Filter:         c.buildFilter(teamID, cloudID, level, provider),
	}
	return &request
}

func (c *Client) buildFilter(teamID, cloudID, level, provider string) filter {
	f := filter{}
	if teamID != content.None {
		f.Teams = []string{teamID}
	}

	if cloudID != content.None {
		f.CloudAccounts = []string{cloudID}
	}

	if level != "" {
		f.Levels = strings.Split(strings.Replace(level, " ", "", -1), "|")
	}

	if provider != "" {
		f.Providers = []string{provider}
	}

	return f
}

//getResultObject returns at most 200 objects, this is chunk design in webapp
func (c *Client) getResultObjectsPaginated(ctx context.Context, request *resultObjectRequest, href string, buf []*ResultObject) (*ResultObjectWrapper, error) {
	result := new(ResultObjectWrapper)
	result.Objects = buf
	jsonStr, err := json.Marshal(*request)
	if err != nil {
		return nil, err
	}
	err = c.Do(ctx, "POST", href, bytes.NewBuffer(jsonStr), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) getAllResultRule(ctx context.Context, request *resultRuleRequest) ([]*ResultRule, error) {
	rules := make([]*ResultRule, 0)
	var result *ResultRuleWrapper
	jsonStr, err := json.Marshal(*request)
	if err != nil {
		return nil, err
	}

	link, err := c.getResultLinkByRef(ctx, "rule")
	if err != nil {
		return nil, err
	}

	// Getting rules using a streaming mechanism
	resp, err := c.makeRequest(ctx, "POST", link.Href, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, NewError(err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		message := new(bytes.Buffer)
		message.ReadFrom(resp.Body)
		msg := fmt.Sprintf("%s", message.String())
		return nil, NewError(msg)
	}

	decoder := json.NewDecoder(resp.Body)
	for {
		if err = decoder.Decode(&result); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		rules = append(rules, result.ViolatingRules.Rules...)
	}

	if len(rules) == 0 {
		return nil, NewError("No violated rule")
	}
	return rules, nil
}
