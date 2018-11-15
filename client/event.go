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
)

//EventStreamConfig for event stream setup
type EventStreamConfig struct {
	TemplateURL     string   `json:"templateURL"`
	TopicName       string   `json:"topicName"`
	StackName       string   `json:"stackName"`
	DevtimeQueueArn string   `json:"devtimeQueueArn"`
	Version         string   `json:"version"`
	MonitorRule     string   `json:"monitorRule"`
	Regions         []string `json:"regions"`
}

//GetSetupConfig get the config for event stream setup from secure state
func (c *Client) GetSetupConfig(ctx context.Context, teamID, cloudID string) (*EventStreamConfig, error) {
	config := &EventStreamConfig{}

	accounts, err := c.GetCloudAccountByID(ctx, teamID, cloudID)
	link, err := GetLinkByRef(accounts.Links, "setup")
	if err != nil {
		return nil, err
	}

	err = c.Do(ctx, "GET", link.Href, nil, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
