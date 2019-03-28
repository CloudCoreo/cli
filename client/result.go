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
	"strings"

	"github.com/CloudCoreo/cli/cmd/content"
)

//CloudAccountInfo records the info of a cloud account
type CloudAccountInfo struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

//TeamInfo records the info of a team
type TeamInfo struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

//TeamInfoWrapper is a wrapper for team Info
type TeamInfoWrapper struct {
	TeamInfo *TeamInfo `json:"team"`
}

//Info is the struct for rule_report
type Info struct {
	SuggestedAction          string `json:"suggested_action"`
	Link                     string `json:"link"`
	Description              string `json:"description"`
	DisplayName              string `json:"display_name"`
	Level                    string `json:"level"`
	Service                  string `json:"service"`
	Name                     string `json:"name"`
	Region                   string `json:"region"`
	IncludeViolationsInCount bool   `json:"include_violations_in_count"`
	TimeStamp                string `json:"timestamp"`
}

// ResultRule struct decodes json file returned by webapp
type ResultRule struct {
	ID     string            `json:"id"`
	Info   Info              `json:"info"`
	TInfo  []TeamInfoWrapper `json:"teamAndPlan"`
	CInfo  []string          `json:"accounts"`
	Object int               `json:"objects"`
}

// The ResultObject struct decodes json file returned by webapp
type ResultObject struct {
	ID        string           `json:"id"`
	Info      Info             `json:"rule_report"`
	TInfo     TeamInfo         `json:"team"`
	CInfo     CloudAccountInfo `json:"cloud_account"`
	RunID     string           `json:"run_id"`
	RiskScore int              `json:"riskScore"`
}

// ResultObjectWrapper contains an object array and number of total items
type ResultObjectWrapper struct {
	Objects    []*ResultObject `json:"violations"`
	TotalItems *int            `json:"totalItems"`
	ScrollID   string          `json:"scrollId,omitempty"`
}

type ResultRuleWrapper struct {
	Rules []*ResultRule `json:"result"`
}

type resultObjectRequest struct {
	RemoveScrollID bool   `json:"removeScrollId"`
	ScrollID       string `json:"scrollId,omitempty"`
	Filter         filter `json:"filter,omitempty"`
}

type filter struct {
	CloudAccounts []string `json:"cloudAccounts,omitempty"`
	Levels        []string `json:"levels,omitempty"`
	Providers     []string `json:"providers,omitempty"`
	Teams         []string `json:"teams,omitempty"`
}

//ShowResultObject shows violated objects. If the filter condition (teamID, cloudID in this case) is valid,
//objects will be filtered. Otherwise return all violation objects under this user account.
func (c *Client) ShowResultObject(ctx context.Context, teamID, cloudID, level, provider string) (*ResultObjectWrapper, error) {

	result, err := c.getAllResultObjects(ctx, teamID, cloudID, level, provider)
	if err != nil {
		return nil, NewError(err.Error())
	}
	res := new(ResultObjectWrapper)
	res.Objects = result
	var num = len(result)
	res.TotalItems = &num

	return res, nil
}

//ShowResultRule show violated rules. If the filter condition (teamID, cloudID in this case) is valid,
//rules will be filtered. Otherwise return all violation rules under this user account.
func (c *Client) ShowResultRule(ctx context.Context, teamID, cloudID, level string) ([]*ResultRule, error) {
	result, err := c.getAllResultRule(ctx)
	res := []*ResultRule{}
	if err != nil {
		return nil, NewError(err.Error())
	}

	targetLevels := strings.Split(strings.Replace(level, " ", "", -1), "|")
	for i := range result {
		if (teamID == content.None || hasTeamID(result[i].TInfo, teamID)) &&
			(cloudID == content.None || hasCloudID(result[i].CInfo, cloudID)) &&
			(level == "" || hasLevel(targetLevels, result[i].Info.Level)) {
			res = append(res, result[i])
		}
	}
	return res, nil
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

func (c *Client) getAllResultObjects(ctx context.Context, teamID, cloudID, level, provider string) ([]*ResultObject, error) {
	link, err := c.getResultLinkByRef(ctx, "object")
	if err != nil {
		return nil, err
	}
	// buffer used to store objects
	buf := make([]*ResultObject, 0, 200)
	res := make([]*ResultObject, 0)

	var scrollId string
	var cur = 0
	var totalItems = 0
	var request *resultObjectRequest

	for cur == 0 || cur < totalItems {
		if cur == 0 {
			request = c.buildGetResultObjectsRequest(teamID, cloudID, level, scrollId, provider, false)
		}
		tmp, err := c.getResultObjects(ctx, request, link.Href, buf)
		if err != nil {
			return res, err
		}
		if cur == 0 {
			totalItems = *(tmp.TotalItems)
			scrollId = tmp.ScrollID
			request = c.buildGetResultObjectsRequest(teamID, cloudID, level, scrollId, provider, false)
		}

		res = append(res, tmp.Objects...)
		if len(tmp.Objects) < 200 {
			break
		}

		cur += len(tmp.Objects)
	}
	return res, nil
}

func (c *Client) buildGetResultObjectsRequest(teamID, cloudID, level, scrollId, provider string, removeScrollId bool) *resultObjectRequest {
	request := resultObjectRequest{
		RemoveScrollID: removeScrollId,
		ScrollID:       scrollId,
		Filter:         filter{},
	}
	if teamID != content.None {
		request.Filter.Teams = []string{teamID}
	}

	if cloudID != content.None {
		request.Filter.CloudAccounts = []string{cloudID}
	}

	if level != "" {
		request.Filter.Levels = strings.Split(strings.Replace(level, " ", "", -1), "|")
	}

	if provider != "" {
		request.Filter.Providers = []string{provider}
	}
	return &request
}

//getResultObject returns at most 200 objects, this is chunk design in webapp
func (c *Client) getResultObjects(ctx context.Context, request *resultObjectRequest, href string, buf []*ResultObject) (*ResultObjectWrapper, error) {
	result := new(ResultObjectWrapper)
	result.Objects = buf
	jsonStr, err := json.Marshal(*request)
	if err != nil {
		return nil, err
	}
	println(string(jsonStr))
	err = c.Do(ctx, "POST", href, bytes.NewBuffer(jsonStr), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) getAllResultRule(ctx context.Context) ([]*ResultRule, error) {
	result := new(ResultRuleWrapper)

	link, err := c.getResultLinkByRef(ctx, "rule")
	if err != nil {
		return nil, err
	}

	err = c.Do(ctx, "GET", link.Href, nil, result)
	if err != nil {
		return nil, NewError(err.Error())
	}
	if len(result.Rules) == 0 {
		return nil, NewError("No violated rule")
	}
	return result.Rules, nil
}

func hasTeamID(teamInfo []TeamInfoWrapper, teamID string) bool {
	for i := range teamInfo {
		if teamInfo[i].TeamInfo.ID == teamID {
			return true
		}
	}
	return false
}

func hasCloudID(cloudInfo []string, cloudID string) bool {
	for i := range cloudInfo {
		if cloudInfo[i] == cloudID {
			return true
		}
	}
	return false
}

func hasLevel(targetLevel []string, level string) bool {
	for i := range targetLevel {
		if targetLevel[i] == level {
			return true
		}
	}
	return false
}
