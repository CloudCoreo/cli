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
	"github.com/CloudCoreo/cli/cmd/content"
)

type CloudAccountInfo struct {
	Name string `json:"name"`
	ID string `json:"id"`
}

type TeamInfo struct {
	Name string `json:"name"`
	ID string `json:"id"`
}


type Info struct {
	SuggestedAction string `json:"suggested_action"`
	Link string `json:"link"`
	Description string `json:"description"`
	DisplayName string `json:"display_name"`
	Level string `json:"level"`
	Service string `json:"service"`
	Name string `json:"name"`
	Region string `json:"region"`
	IncludeViolationsInCount bool `json:"include_violations_in_count"`
	TimeStamp string `json:"timestamp"`
}

// The RusultRule struct decodes json file returned by webapp
type ResultRule struct {
	ID string `json:"id"`
	Info Info `json:"info"`
	TInfo []TeamInfo `json:"teams"`
	CInfo []CloudAccountInfo `json:"accounts"`
	Object int `json:"objects"`
}

// The RusultObject struct decodes json file returned by webapp
type ResultObject struct {
	ID string `json:"id"`
	Info Info `json:"rule_report"`
	TInfo []TeamInfo `json:"team"`
	CInfo []CloudAccountInfo `json:"cloud_account"`
	RunId string `json:"run_id"`
}

//Show violated objects. If the filter condition (teamID, cloudID in this case) is valid,
//objects will be filtered. Otherwise return all violation objects under this user account.
func (c *Client) ShowResultObject(ctx context.Context, teamID, cloudID string) ([]* ResultObject, error) {
	result, err := c.getAllResultObject(ctx)
	res := []*ResultObject{}
	if err != nil {
		return nil, NewError(err.Error())
	}

	if teamID != content.None && cloudID != content.None{
		for i := range result {
			if hasTeamID(result[i].TInfo, teamID) && hasCloudID(result[i].CInfo, cloudID) {
				res = append(res, result[i])
			}
		}
	} else if teamID != content.None{
		for i := range result {
			if hasTeamID(result[i].TInfo, teamID){
				res = append(res, result[i])
			}
		}

	} else if cloudID != content.None{
		for i := range result {
			if hasCloudID(result[i].CInfo, cloudID) {
				res = append(res, result[i])
			}
		}
	} else {
		copy(res, result)
	}

	if len(res) == 0 {
		return nil, NewError("No violated object")
	}
	return result, nil
}

//Show violated rules. If the filter condition (teamID, cloudID in this case) is valid,
//rules will be filtered. Otherwise return all violation rules under this user account.
func (c *Client) ShowResultRule(ctx context.Context, teamID, cloudID string) ([]* ResultRule, error) {
	result, err := c.getAllResultRule(ctx)
	res := []*ResultRule{}
	if err != nil {
		return nil, NewError(err.Error())
	}

	if teamID != content.None && cloudID != content.None{
		for i := range result {
			if hasTeamID(result[i].TInfo, teamID) && hasCloudID(result[i].CInfo, cloudID) {
				res = append(res, result[i])
			}
		}
	} else if teamID != content.None{
		for i := range result {
			if hasTeamID(result[i].TInfo, teamID){
			res = append(res, result[i])
			}
		}

	} else if cloudID != content.None{
		for i := range result {
			if hasCloudID(result[i].CInfo, cloudID) {
			res = append(res, result[i])
			}
		}
	} else {
		res = result
	}

	if len(res) == 0 {
		return nil, NewError("No violated rule")
	}
	return res, nil
}

func (c *Client) getAllResultObject(ctx context.Context) ([]* ResultObject, error) {
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

func (c *Client) getAllResultRule(ctx context.Context) ([]* ResultRule, error) {
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
	for i := range teamInfo  {
		if teamInfo[i].ID == teamID{
			return true
		}
	}
	return false
}

func hasCloudID(cloudInfo []CloudAccountInfo, cloudID string) bool {
	for i := range cloudInfo {
		if cloudInfo[i].ID == cloudID{
			return true
		}
	}
	return false
}
