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
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/CloudCoreo/cli/client/content"
)

//Panel panel struct
type Panel struct {
	ResourcesArray []struct {
		StackName       string    `json:"stackName"`
		RunID           string    `json:"runId"`
		EngineStatus    string    `json:"engineStatus"`
		Namespace       string    `json:"namespace"`
		ExecutionNumber int       `json:"executionNumber"`
		Timestamp       time.Time `json:"timestamp"`
		ExecutionTime   float64   `json:"executionTime"`
		ResourceType    string    `json:"resourceType"`
		ResourceName    string    `json:"resourceName"`
		DataType        string    `json:"dataType"`
		Inputs          []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
			Name  string `json:"name"`
		} `json:"inputs"`
		Outputs []struct {
			Value string `json:"value"`
			Name  string `json:"name"`
		} `json:"outputs"`
		ResourceID string `json:"resourceId"`
		ID         string `json:"_id"`
	} `json:"resourcesArray"`
	NumberOfResources          int       `json:"numberOfResources"`
	PlanRefreshIntervalInHours int       `json:"planRefreshIntervalInHours"`
	LastExecutionTime          time.Time `json:"lastExecutionTime"`
	EngineState                string    `json:"engineState"`
	IsEnabled                  bool      `json:"isEnabled"`
}

// Plan struct object
type Plan struct {
	DefaultPanelRepo       string `json:"defaultPanelRepo"`
	DefaultPanelDirectory  string `json:"defaultPanelDirectory"`
	DefaultPanelBranch     string `json:"defaultPanelBranch"`
	Name                   string `json:"name"`
	EnginePrefix           string `json:"enginePrefix"`
	IamUserAccessKeyID     string `json:"iamUserAccessKeyId"`
	IamUserID              string `json:"iamUserId"`
	IamUserSecretAccessKey string `json:"iamUserSecretAccessKey"`
	SnsSubscriptionArn     string `json:"snsSubscriptionArn"`
	SqsArn                 string `json:"sqsArn"`
	SqsURL                 string `json:"sqsUrl"`
	TopicArn               string `json:"topicArn"`
	IsSynchronizing        bool   `json:"isSynchronizing"`
	IsDraft                bool   `json:"isDraft"`
	DefaultRegion          string `json:"defaultRegion"`
	RefreshInterval        int    `json:"refreshInterval"`
	RunCounter             int    `json:"runCounter"`
	Revision               string `json:"revision"`
	Branch                 string `json:"branch"`
	Enabled                bool   `json:"enabled"`
	Links                  []Link `json:"links"`
	ID                     string `json:"id"`
}

// PlanConfig struct object
type PlanConfig struct {
	GitRevision string                   `json:"gitRevision"`
	GitBranch   string                   `json:"gitBranch"`
	TeamID      string                   `json:"teamId"`
	Links       []Link                   `json:"links"`
	Variables   map[string]PlanAttribute `json:"variables"`
	PlanID      string                   `json:"planId"`
	ID          string                   `json:"id"`
}

// PlanAttribute struct object
type PlanAttribute struct {
	Namespace   string      `json:"namespace"`
	Default     interface{} `json:"default"`
	Description string      `json:"description"`
	Required    bool        `json:"required"`
	Type        string      `json:"type"`
	Class       string      `json:"class"`
	Value       interface{} `json:"value"`
}

// GetPlans method to get plans info array object
func (c *Client) GetPlans(ctx context.Context, teamID, compositeID string) ([]*Plan, error) {
	composite, err := c.GetCompositeByID(ctx, teamID, compositeID)

	if err != nil {
		return nil, err
	}

	plansLink, err := GetLinkByRef(composite.Links, "plans")
	if err != nil {
		return nil, err
	}

	plans := []*Plan{}
	err = c.Do(ctx, "GET", plansLink.Href, nil, &plans)
	if err != nil {
		return nil, err
	}

	if len(plans) == 0 {
		return nil, NewError(fmt.Sprintf(content.ErrorNoPlansFound, teamID, compositeID))
	}

	return plans, nil
}

// GetPlanByID method to get plan info object
func (c *Client) GetPlanByID(ctx context.Context, teamID, compositeID, planID string) (*Plan, error) {
	plans, err := c.GetPlans(ctx, teamID, compositeID)
	if err != nil {
		return nil, err
	}

	plan := &Plan{}
	for _, p := range plans {
		if p.ID == planID {
			plan = p
			break
		}
	}

	if plan.ID == "" {
		return nil, NewError(fmt.Sprintf(content.ErrorNoPlanWithIDFound, planID, teamID, compositeID))
	}

	return plan, nil
}

// DeletePlanByID method to delete cloud object
func (c *Client) DeletePlanByID(ctx context.Context, teamID, compositeID, planID string) error {
	plans, err := c.GetPlans(ctx, teamID, compositeID)

	if err != nil {
		return err
	}

	planFound := false
	for _, plan := range plans {
		if plan.ID == planID {
			planFound = true
			planLink, err := GetLinkByRef(plan.Links, "self")
			if err != nil {
				return err
			}

			err = c.Do(ctx, "DELETE", planLink.Href, nil, nil)
			if err != nil {
				return err
			}
			break
		}
	}

	if !planFound {
		return NewError(fmt.Sprintf(content.ErrorFailedToDeletePlan, planID, teamID, compositeID))
	}

	return nil
}

// EnablePlan method to enable a plan object
func (c *Client) EnablePlan(ctx context.Context, teamID, compositeID, planID string) (*Plan, error) {
	plans, err := c.GetPlans(ctx, teamID, compositeID)

	if err != nil {
		return nil, err
	}
	plan := &Plan{}

	for _, p := range plans {

		if p.ID == planID {
			p.Enabled = true
			jsonStr, err := json.Marshal(p)
			if err != nil {
				return nil, err
			}

			planLink, err := GetLinkByRef(p.Links, "self")
			if err != nil {
				return nil, err
			}

			err = c.Do(ctx, "PUT", planLink.Href, bytes.NewBuffer(jsonStr), &plan)
			if err != nil {
				return nil, err
			}
			break
		}
	}

	if plan.ID == "" && !plan.Enabled {
		return nil, NewError(fmt.Sprintf(content.ErrorFailedToEnablePlan, planID, teamID, compositeID))
	}

	return plan, nil
}

// DisablePlan method to disable a plan object
func (c *Client) DisablePlan(ctx context.Context, teamID, compositeID, planID string) (*Plan, error) {
	plans, err := c.GetPlans(ctx, teamID, compositeID)

	if err != nil {
		return nil, err
	}

	plan := &Plan{}

	for _, p := range plans {

		if p.ID == planID {
			p.Enabled = false
			jsonStr, err := json.Marshal(p)
			if err != nil {
				return plan, err
			}

			planLink, err := GetLinkByRef(p.Links, "self")
			if err != nil {
				return plan, err
			}

			err = c.Do(ctx, "PUT", planLink.Href, bytes.NewBuffer(jsonStr), &plan)
			if err != nil {
				return plan, err
			}
			break
		}
	}

	if plan.ID == "" {
		return nil, NewError(fmt.Sprintf(content.ErrorFailedToDisblePlan, planID, teamID, compositeID))
	}

	return plan, nil
}

//InitPlan init plan
func (c *Client) InitPlan(ctx context.Context, branch, name, region, teamID, cloudID, compositeID, revision string, interval int) (*PlanConfig, error) {
	composite, err := c.GetCompositeByID(ctx, teamID, compositeID)
	if err != nil {
		return nil, err
	}

	plansLink, err := GetLinkByRef(composite.Links, "plans")
	if err != nil {
		return nil, err
	}

	planPayLoad := fmt.Sprintf(
		`{"name":"%s","awsCredsId":"%s","region":"%s","branch":"%s","revision":"%s","refreshInterval":"%d","appStackId":"%s"}`,

		name,
		cloudID,
		region,
		branch,
		revision,
		interval,
		compositeID)

	var jsonStr = []byte(planPayLoad)
	plan := &Plan{}
	err = c.Do(ctx, "POST", plansLink.Href, bytes.NewBuffer(jsonStr), &plan)
	if err != nil {

		fmt.Println(err.Error())
		return nil, err
	}

	planLink, err := GetLinkByRef(plan.Links, "self")
	if err != nil {
		return nil, err
	}

	jsonStr, err = json.Marshal(plan)
	if err != nil {
		return nil, err
	}

	err = c.Do(ctx, "PUT", planLink.Href, bytes.NewBuffer(jsonStr), &plan)
	if err != nil {

		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Print(content.InfoPlanCreationMessage)

	planConfig := &PlanConfig{}

	for {
		fmt.Print(".")
		planConfig, err = c.getPlanConfig(ctx, plan)
		if err == nil {
			fmt.Println()
			return planConfig, nil
		}
		time.Sleep(5 * time.Second)
	}
}

func (c *Client) getPlanConfig(ctx context.Context, plan *Plan) (*PlanConfig, error) {
	planConfig := &PlanConfig{}
	planConfigLink, err := GetLinkByRef(plan.Links, "planconfig")
	if err != nil {
		return nil, err
	}

	err = c.Do(ctx, "GET", planConfigLink.Href, nil, &planConfig)
	if err != nil {
		return nil, err
	}

	return planConfig, nil

}

//CreatePlan Create plan
func (c *Client) CreatePlan(ctx context.Context, planConfigContent []byte) (*Plan, error) {

	planConfig := &PlanConfig{}
	if err := json.Unmarshal(planConfigContent, &planConfig); err != nil {
		return nil, err
	}

	planConfigLink, err := GetLinkByRef(planConfig.Links, "self")
	if err != nil {
		return nil, err
	}

	err = c.Do(ctx, "PUT", planConfigLink.Href, bytes.NewBuffer(planConfigContent), nil)
	if err != nil {
		return nil, err
	}

	planLink, err := GetLinkByRef(planConfig.Links, "plan")
	if err != nil {
		return nil, err
	}

	plan := &Plan{}
	err = c.Do(ctx, "GET", planLink.Href, nil, &plan)
	if err != nil {
		return nil, err
	}

	planLink, err = GetLinkByRef(plan.Links, "self")
	if err != nil {
		return nil, err
	}

	plan.IsDraft = false
	jsonStr, err := json.Marshal(plan)

	err = c.Do(ctx, "PUT", planLink.Href, bytes.NewBuffer(jsonStr), &plan)
	if err != nil {
		return nil, err
	}

	return plan, nil

}

//GetPanelInfo get panel info
func (c *Client) GetPanelInfo(ctx context.Context, teamID, compositeID, planID string) (*Panel, error) {
	panel := &Panel{}
	plan, err := c.GetPlanByID(ctx, teamID, compositeID, planID)
	if err != nil {
		return nil, err
	}

	panelLink, err := GetLinkByRef(plan.Links, "panel")
	if err != nil {
		return nil, err
	}

	err = c.Do(ctx, "GET", panelLink.Href, nil, &panel)
	if err != nil {
		return nil, err
	}

	return panel, nil
}
