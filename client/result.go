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

	"github.com/CloudCoreo/cli/pkg/command"

	"github.com/CloudCoreo/cli/cmd/content"
)

//ShowResultObject shows violated objects. If the filter condition (teamID, cloudID in this case) is valid,
//objects will be filtered. Otherwise return all violation objects under this user account.
func (c *Client) ShowResultObject(ctx context.Context, teamID, cloudID, level string) ([]*command.ResultObject, error) {
	result, err := c.getAllResultObject(ctx)
	res := []*command.ResultObject{}
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
	return res, nil
}

//ShowResultRule show violated rules. If the filter condition (teamID, cloudID in this case) is valid,
//rules will be filtered. Otherwise return all violation rules under this user account.
func (c *Client) ShowResultRule(ctx context.Context, teamID, cloudID, level string) ([]*command.ResultRule, error) {
	result, err := c.getAllResultRule(ctx)
	res := []*command.ResultRule{}
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

func (c *Client) getAllResultObject(ctx context.Context) ([]*command.ResultObject, error) {
	result := []*command.ResultObject{}
	err := c.Do(ctx, "GET", c.endpoint+"/result/object", nil, &result)
	if err != nil {
		return nil, NewError(err.Error())
	}

	if len(result) == 0 {
		return nil, NewError("No violated object")
	}
	return result, nil
}

func (c *Client) getAllResultRule(ctx context.Context) ([]*command.ResultRule, error) {
	result := []*command.ResultRule{}
	err := c.Do(ctx, "GET", c.endpoint+"/result/rule", nil, &result)
	if err != nil {
		return nil, NewError(err.Error())
	}
	if len(result) == 0 {
		return nil, NewError("No violated rule")
	}
	return result, nil
}

func hasTeamID(teamInfo []command.TeamInfo, teamID string) bool {
	for i := range teamInfo {
		if teamInfo[i].ID == teamID {
			return true
		}
	}
	return false
}

func hasCloudID(cloudInfo []command.CloudAccountInfo, cloudID string) bool {
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
