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
	"os"
	"strings"

	"github.com/CloudCoreo/cli/cmd/content"
)

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
	ID        string   `json:"id"`
	Info      Info     `json:"rule_report"`
	TInfo     TeamInfo `json:"team"`
	RiskScore int      `json:"riskScore"`
}

// ResultObjectWrapper contains an object array and number of total items
type ResultObjectWrapper struct {
	AccountName   string          `json:"accountName,omitempty"`
	AccountNumber string          `json:"accountNumber,omitempty"`
	TotalItems    int             `json:"totalItems"`
	Objects       []*ResultObject `json:"violations"`
	ScrollID      string          `json:"scrollId,omitempty"`
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
func (c *Client) ShowResultObject(ctx context.Context, teamID, cloudID, level, provider string, retry uint) ([]*ResultObjectWrapper, error) {
	var res []*ResultObjectWrapper
	if cloudID != content.None {
		result, err := c.getResultObjects(ctx, teamID, cloudID, level, provider, retry)
		if err != nil {
			return nil, err
		}
		return []*ResultObjectWrapper{result}, nil
	} else {
		var teams []*Team
		//If teamID is None, then get all teams, otherwise only get team with <teamID>
		teams, err := c.getTeamsForObjects(ctx, teamID)
		if err != nil {
			return nil, err
		}
		for _, team := range teams {
			accounts, err := c.GetCloudAccounts(ctx, team.ID)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				continue
			}
			for _, account := range accounts {
				result, err := c.getResultObjects(ctx, team.ID, account.ID, level, provider, retry)
				if err != nil {
					return nil, err
				}
				res = append(res, result)

			}
		}
		return res, nil
	}
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

//If teamID is None, then return all teams, otherwise return teamID passed
func (c *Client) getTeamsForObjects(ctx context.Context, teamID string) ([]*Team, error) {
	if teamID != content.None {
		return []*Team{{ID: teamID}}, nil
	}
	return c.GetTeams(ctx)
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

	var scrollId string
	var cur = 0
	var firstCall = true
	var totalItems = 0
	var request *resultObjectRequest

	for firstCall || cur < totalItems {
		if firstCall {
			request = c.buildGetResultObjectsRequest(teamID, cloudID, level, scrollId, provider, false)
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
		if firstCall {
			totalItems = tmp.TotalItems
			scrollId = tmp.ScrollID
			request = c.buildGetResultObjectsRequest(teamID, cloudID, level, scrollId, provider, false)
			firstCall = false
		}

		res = append(res, tmp.Objects...)
		if len(tmp.Objects) < 200 {
			break
		}

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
