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
	TInfo  []TeamInfo         `json:"teams"`
	CInfo  []CloudAccountInfo `json:"accounts"`
	Object int                `json:"objects"`
}

// The ResultObject struct decodes json file returned by webapp
type ResultObject struct {
	ID    string           `json:"id"`
	Info  Info             `json:"rule_report"`
	TInfo TeamInfo         `json:"team"`
	CInfo CloudAccountInfo `json:"cloud_account"`
	RunID string           `json:"run_id"`
}

//ShowResultObject shows violated objects. If the filter condition (teamID, cloudID in this case) is valid,
//objects will be filtered. Otherwise return all violation objects under this user account.
func (c *Client) ShowResultObject(ctx context.Context, teamID, cloudID, level string) ([]*ResultObject, error) {
	result, err := c.getAllResultObject(ctx)
	res := []*ResultObject{}
	if err != nil {
		return nil, NewError(err.Error())
	}

	targetLevels := strings.Split(strings.Replace(level, " ", "", -1), "|")
	for i := range result {
		if (teamID == content.None || result[i].TInfo.ID == teamID) &&
			(cloudID == content.None || result[i].CInfo.ID == cloudID) &&
			(level == content.None || hasLevel(targetLevels, result[i].Info.Level)) {
			res = append(res, result[i])
		}
	}

	if len(res) == 0 {
		return nil, NewError("No violated object")
	}
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

	targetLevels := strings.Split(strings.Replace(level, " ", "", -1), ",")
	for i := range result {
		if (teamID == content.None || hasTeamID(result[i].TInfo, teamID)) &&
			(cloudID == content.None || hasCloudID(result[i].CInfo, cloudID)) &&
			(level == content.None || hasLevel(targetLevels, result[i].Info.Level)) {
			res = append(res, result[i])
		}
	}

	if len(res) == 0 {
		return nil, NewError("No violated rule")
	}
	return res, nil
}

func (c *Client) getAllResultObject(ctx context.Context) ([]*ResultObject, error) {
	result := []*ResultObject{}
	err := c.Do(ctx, "GET", c.endpoint+"/result/object", nil, &result)
	if err != nil {
		return nil, NewError(err.Error())
	}

	if len(result) == 0 {
		return nil, NewError("No violated object")
	}
	return result, nil
}

func (c *Client) getAllResultRule(ctx context.Context) ([]*ResultRule, error) {
	result := []*ResultRule{}
	err := c.Do(ctx, "GET", c.endpoint+"/result/rule", nil, &result)
	if err != nil {
		return nil, NewError(err.Error())
	}
	if len(result) == 0 {
		return nil, NewError("No violated rule")
	}
	return result, nil
}

func hasTeamID(teamInfo []TeamInfo, teamID string) bool {
	for i := range teamInfo {
		if teamInfo[i].ID == teamID {
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
