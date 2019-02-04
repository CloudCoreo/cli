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
	"context"
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
	ID     string             `json:"id"`
	Info   Info               `json:"info"`
	TInfo  []TeamInfoWrapper  `json:"teamAndPlan"`
	CInfo  []CloudAccountInfo `json:"accounts"`
	Object int                `json:"objects"`
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
}

type ResultRuleWrapper struct {
	Rules []*ResultRule `json:"rules"`
}

//ShowResultObject shows violated objects. If the filter condition (teamID, cloudID in this case) is valid,
//objects will be filtered. Otherwise return all violation objects under this user account.
func (c *Client) ShowResultObject(ctx context.Context, teamID, cloudID, level string) (*ResultObjectWrapper, error) {
	result, err := c.getAllResultObject(ctx)
	res := new(ResultObjectWrapper)
	if err != nil {
		return nil, NewError(err.Error())
	}

	targetLevels := strings.Split(strings.Replace(level, " ", "", -1), "|")
	for i := range result.Objects {
		if (teamID == content.None || result.Objects[i].TInfo.ID == teamID) &&
			(cloudID == content.None || result.Objects[i].CInfo.ID == cloudID) &&
			(level == content.None || hasLevel(targetLevels, result.Objects[i].Info.Level)) {
			res.Objects = append(res.Objects, result.Objects[i])
		}
	}
	res.TotalItems = result.TotalItems
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
			(level == content.None || hasLevel(targetLevels, result[i].Info.Level)) {
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
func (c *Client) getAllResultObject(ctx context.Context) (*ResultObjectWrapper, error) {
	result := new(ResultObjectWrapper)

	link, err := c.getResultLinkByRef(ctx, "object")
	if err != nil {
		return nil, err
	}

	err = c.Do(ctx, "GET", link.Href, nil, result)
	if err != nil {
		return nil, NewError(err.Error())
	}

	if len(result.Objects) == 0 {
		return nil, NewError("No violated object")
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

func hasCloudID(cloudInfo []CloudAccountInfo, cloudID string) bool {
	for i := range cloudInfo {
		if cloudInfo[i].ID == cloudID {
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
