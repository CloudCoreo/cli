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
	ResourcesArray             interface{} `json:"resourcesArray"`
	NumberOfResources          int         `json:"numberOfResources"`
	PlanRefreshIntervalInHours float32     `json:"planRefreshIntervalInHours"`
	LastExecutionTime          int64       `json:"lastExecutionTime"`
	EngineState                string      `json:"engineState"`
	IsEnabled                  bool        `json:"isEnabled"`
}

// Plan struct object
type Plan struct {
	EngineRunInfos struct {
		EngineState        string    `json:"engineState"`
		EngineStatus       string    `json:"engineStatus"`
		RunID              string    `json:"runId"`
		CreatedAt          time.Time `json:"createdAt"`
		NumberOfResources  int       `json:"numberOfResources"`
		EngineStateMessage struct {
			ResourceError string `json:"resource_error"`
			ErrorType     string `json:"error_type"`
			ErrorMessage  string `json:"error_message"`
		} `json:"engineStateMessage"`
	} `json:"engineRunInfos"`
	DefaultPanelRepo       string `json:"defaultPanelRepo"`
	DefaultPanelDirectory  string `json:"defaultPanelDirectory"`
	DefaultPanelBranch     string `json:"defaultPanelBranch"`
	Name                   string `json:"name"`
	IamUserAccessKeyID     string `json:"iamUserAccessKeyId"`
	IamUserID              string `json:"iamUserId"`
	IamUserSecretAccessKey string `json:"iamUserSecretAccessKey"`
	SnsSubscriptionArn     string `json:"snsSubscriptionArn"`
	SqsArn                 string `json:"sqsArn"`
	SqsURL                 string `json:"sqsUrl"`
	TopicArn               string `json:"topicArn"`
	EnginePrefix           string `json:"enginePrefix"`
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
	IntervalInMinutes      int    `json:"intervalInMinutes"`
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

	planLink, err := GetLinkByRef(plan.Links, "self")
	if err != nil {
		return nil, err
	}

	err = c.Do(ctx, "GET", planLink.Href, nil, &plan)
	if err != nil {
		return nil, err
	}

	return plan, nil
}

// RunNowPlanByID method to run-now plan info object
func (c *Client) RunNowPlanByID(ctx context.Context, teamID, compositeID, planID string, block bool) (*Plan, error) {
	plan, err := c.GetPlanByID(ctx, teamID, compositeID, planID)
	if err != nil {
		return nil, err
	}

	runNowLink, err := GetLinkByRef(plan.Links, "runnow")
	if err != nil {
		return nil, err
	}

	err = c.Do(ctx, "GET", runNowLink.Href, nil, &plan)
	if err != nil {
		return nil, err
	}

	planLink, err := GetLinkByRef(plan.Links, "self")
	if err != nil {
		return nil, err
	}

	prevRunID := plan.EngineRunInfos.RunID

	if block {
		fmt.Print(content.InfoPlanRunNowMessageBlock)
		for {
			fmt.Print(".")
			_ = c.Do(ctx, "GET", planLink.Href, nil, &plan)
			if plan.EngineRunInfos.RunID != prevRunID && plan.EngineRunInfos.EngineState == "COMPLETED" {
				break
			}

			time.Sleep(5 * time.Second)
		}

		fmt.Println()

		if plan.EngineRunInfos.EngineStatus == "OK" {
			return plan, nil
		}
		return nil, fmt.Errorf(content.ErrorPlanRunNow, plan.EngineRunInfos.EngineStateMessage.ErrorMessage)
	}

	fmt.Print(content.InfoPlanRunNowMessage)
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
func (c *Client) InitPlan(ctx context.Context, branch, name, region, teamID, cloudID, compositeID, revision string, intervalInMinutes int) (*PlanConfig, error) {

	if intervalInMinutes < 2 || intervalInMinutes > 525600 {
		return nil, fmt.Errorf(content.ErrorPlanIntervalMintuesInvalid)
	}

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
		intervalInMinutes,
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

	for {
		fmt.Print(".")
		_ = c.Do(ctx, "GET", planLink.Href, nil, &plan)

		if plan.EngineRunInfos.EngineState == "INITIALIZED" {
			break
		}

		time.Sleep(5 * time.Second)
	}

	fmt.Println()

	planConfig := &PlanConfig{}
	if plan.EngineRunInfos.EngineStatus == "OK" {
		planConfig, err = c.getPlanConfig(ctx, plan)
		if err == nil {
			planConfig.Variables = setPlanConfigRequiredValues(planConfig.Variables)
			return planConfig, nil
		}
	}

	return nil, fmt.Errorf(content.ErrorPlanCreation, plan.EngineRunInfos.EngineStateMessage.ErrorMessage)
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

	err := verifyPlanConfigRequiredValues(planConfig.Variables)

	if err != nil {
		return nil, err
	}

	planConfigLink, err := GetLinkByRef(planConfig.Links, "self")
	if err != nil {
		return nil, err
	}

	jsonStr, err := json.Marshal(planConfig)
	err = c.Do(ctx, "PUT", planConfigLink.Href, bytes.NewBuffer(jsonStr), nil)
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

	fmt.Print(content.InfoPlanCompilingMessage)
	compilePlanLink, err := GetLinkByRef(plan.Links, "compile_now")
	if err != nil {
		return nil, err
	}

	err = c.Do(ctx, "GET", compilePlanLink.Href, nil, &plan)
	if err != nil {
		return nil, err
	}

	for {
		fmt.Print(".")
		_ = c.Do(ctx, "GET", planLink.Href, nil, &plan)
		if plan.EngineRunInfos.EngineState == "COMPILED" || plan.EngineRunInfos.EngineState == "COMPLETED" {
			break
		}

		time.Sleep(5 * time.Second)
	}

	fmt.Println()

	if plan.EngineRunInfos.EngineStatus != "OK" {
		return nil, fmt.Errorf(content.ErrorPlanCompileNow, plan.EngineRunInfos.EngineStateMessage.ErrorMessage)
	}

	planLink, err = GetLinkByRef(plan.Links, "self")
	if err != nil {
		return nil, err
	}

	plan.IsDraft = false
	jsonStr, err = json.Marshal(plan)

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

func verifyPlanConfigRequiredValues(variables map[string]PlanAttribute) error {

	errorFound := false
	for k, v := range variables {
		if v.Required && (v.Value == nil || v.Value == "" || v.Value == "INPUT REQUIRED") {
			errorFound = true
			fmt.Printf(content.ErrorPlanConfigRequiredVariableMissing, k)
		}
	}

	if errorFound {
		return fmt.Errorf(content.ErrorPlanConfigVaribaleMissing)
	}

	return nil
}

func setPlanConfigRequiredValues(variables map[string]PlanAttribute) map[string]PlanAttribute {
	// Add value property
	for i, v := range variables {
		if v.Required {
			if v.Default != nil {
				v.Value = v.Default
			} else {
				v.Value = "INPUT REQUIRED"
			}
			variables[i] = v
		}
	}

	return variables
}
