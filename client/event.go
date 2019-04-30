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
	AWSEventStreamConfig
	AzureEventStreamConfig
	Provider string `json:"provider"`
}

type AWSEventStreamConfig struct {
	TemplateURL     string   `json:"templateURL"`
	TopicName       string   `json:"topicName"`
	StackName       string   `json:"stackName"`
	DevtimeQueueArn string   `json:"devtimeQueueArn"`
	Version         string   `json:"version"`
	MonitorRule     string   `json:"monitorRule"`
	Regions         []string `json:"regions"`
}
type AzureEventStreamConfig struct {
	SubscriptionID       string `json:"subscriptionId"`
	ActionDeployFile     string `json:"actionDeployFile"`
	AlertDeployFile      string `json:"alertDeployFile"`
	WebhookServiceUri    string `json:"webhookServiceUri"`
	ResourceGroup        string `json:"resourceGroup"`
	ActionDeploymentName string `json:"actionDeploymentName"`
	AlertDeploymentName  string `json:"alertDeploymentName"`
	ActionGroup          string `json:"actionGroup"`
	ActionGroupShort     string `json:"actionGroupShort"`
	WebhookReceiverName  string `json:"webhookReceiverName"`
	AlertName            string `json:"alertName"`
}

//EventRemoveConfig for event stream removal
type EventRemoveConfig struct {
	AWSEventRemoveConfig
	AzureEventRemoveConfig
	Provider string `json:"provider"`
}

type AWSEventRemoveConfig struct {
	StackName      string   `json:"stackName"`
	TopicName      string   `json:"topicName"`
	Regions        []string `json:"regions"`
	ArnType        string   `json:"arnType"`
	CloudAccountId string   `json:"cloudAccountId"`
}

type AzureEventRemoveConfig struct {
	SubscriptionID    string `json:"subscriptionId"`
	ResourceGroup     string `json:"resourceGroup"`
	WebhookServiceUri string `json:"webhookServiceUri"`
}

//GetSetupConfig get the config for event stream setup from secure state
func (c *Client) GetSetupConfig(ctx context.Context, teamID, cloudID string) (*EventStreamConfig, error) {
	config := &EventStreamConfig{}

	accounts, err := c.GetCloudAccountByID(ctx, teamID, cloudID)
	if err != nil {
		return nil, err
	}

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

//GetRemoveConfig get the config for event stream removal from secure state
func (c *Client) GetRemoveConfig(ctx context.Context, teamID, cloudID string) (*EventRemoveConfig, error) {
	config := &EventRemoveConfig{}
	accounts, err := c.GetCloudAccountByID(ctx, teamID, cloudID)
	if err != nil {
		return nil, err
	}

	link, err := GetLinkByRef(accounts.Links, "remove")
	if err != nil {
		return nil, err
	}

	err = c.Do(ctx, "GET", link.Href, nil, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
