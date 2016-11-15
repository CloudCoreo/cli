package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// Plan struct object
type Plan struct {
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
	DefaultRegion          string `json:"defaultRegion"`
	RefreshInterval        int    `json:"refreshInterval"`
	Revision               string `json:"revision"`
	Branch                 string `json:"branch"`
	Enabled                bool   `json:"enabled"`
	Links                  []Link `json:"links"`
	ID                     string `json:"id"`
}

// PlanConfig struct object
type PlanConfig struct {
	AppstackID string `json:"appstackId"`
	TeamID     string `json:"teamId"`
	Links      []Link `json:"links"`
	Variables  string `json:"variables"`
	PlanID     string `json:"planId"`
	ID         string `json:"id"`
}

// GetPlans method to get plans info array object
func (c *Client) GetPlans(ctx context.Context, teamID, compositeID string) ([]Plan, error) {
	plans := []Plan{}
	composite, err := c.GetCompositeByID(ctx, teamID, compositeID)

	if err != nil {
		return plans, err
	}

	plansLink, err := GetLinkByRef(composite.Links, "plans")
	if err != nil {
		return plans, err
	}

	err = c.Do(ctx, "GET", plansLink.Href, nil, &plans)
	if err != nil {
		return plans, err
	}

	return plans, nil
}

// GetPlanByID method to get plan info object
func (c *Client) GetPlanByID(ctx context.Context, teamID, compositeID, planID string) (Plan, error) {
	plan := Plan{}
	plans, err := c.GetPlans(ctx, teamID, compositeID)
	if err != nil {
		return plan, err
	}

	for _, p := range plans {
		if p.ID == planID {
			plan = p
			break
		}
	}

	return plan, nil
}

// DeletePlanByID method to delete cloud object
func (c *Client) DeletePlanByID(ctx context.Context, teamID, compositeID, planID string) error {
	plans, err := c.GetPlans(ctx, teamID, compositeID)

	if err != nil {
		return err
	}

	for _, plan := range plans {
		if plan.ID == planID {
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

	return nil
}

// EnablePlan method to enable a plan object
func (c *Client) EnablePlan(ctx context.Context, teamID, compositeID, planID string) (Plan, error) {
	plan := Plan{}
	plans, err := c.GetPlans(ctx, teamID, compositeID)

	if err != nil {
		return plan, err
	}

	for _, p := range plans {

		if p.ID == planID {
			p.Enabled = true
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

	return plan, nil
}

// DisablePlan method to disable a plan object
func (c *Client) DisablePlan(ctx context.Context, teamID, compositeID, planID string) (Plan, error) {
	plan := Plan{}
	plans, err := c.GetPlans(ctx, teamID, compositeID)

	if err != nil {
		return plan, err
	}

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

	return plan, nil
}

//InitPlan init plan TODO - UPDATE CODE WHEN READY
func (c *Client) InitPlan(ctx context.Context, branch, name, interval, region, teamID, cloudID, compositeID, revision string) (PlanConfig, error) {
	planConfig := PlanConfig{}

	composite, err := c.GetCompositeByID(ctx, teamID, compositeID)
	if err != nil {
		return planConfig, err
	}

	plansLink, err := GetLinkByRef(composite.Links, "plans")
	if err != nil {
		return planConfig, err
	}

	planPayLoad := fmt.Sprintf(
		`{"name":"%s","awsCredsId":"%s","region":"%s","branch":"%s","revision":"%s","refreshInterval":"%s","appStackId":"%s"}`,
		//{"branch":"master","refreshInterval":1,"revision":"HEAD","defaultRegion":"us-east-1","awsCredsId":"5824e64f4008a3cb41af144f","appStackId":"5824e6284008a3cb41af144d","name":"test-plan"}
		name,
		cloudID,
		region,
		branch,
		revision,
		interval,
		compositeID)

	var jsonStr = []byte(planPayLoad)
	err = c.Do(ctx, "POST", plansLink.Href, bytes.NewBuffer(jsonStr), &composite)
	if err != nil {
		return planConfig, err
	}

	for start := time.Now(); time.Since(start) < 300; time.Sleep(10 * time.Second) {
		fmt.Println("Loading planconfig...")
		plans, err := c.GetPlans(ctx, teamID, compositeID)
		if err != nil {
			return planConfig, err
		}

		for _, p := range plans {
			if p.Name == name {
				planConfig, err = c.getPlanConfig(ctx, p)
				if err != nil {
					return planConfig, err
				}
				break
			}
		}
	}

	return planConfig, nil
}

func (c *Client) getPlanConfig(ctx context.Context, plan Plan) (PlanConfig, error) {
	planConfig := PlanConfig{}
	planConfigLink, err := GetLinkByRef(plan.Links, "planconfig")
	if err != nil {
		return planConfig, err
	}

	err = c.Do(ctx, "GET", planConfigLink.Href, nil, &planConfig)
	if err != nil {
		return planConfig, err
	}

	return planConfig, nil

}
