package client

import (
	"bytes"
	"context"
	"fmt"
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
		if p.ID == compositeID {
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
	plan, err := c.updatePlan(ctx, teamID, compositeID, planID, fmt.Sprintf(`{"id":"%s", "enabled":"%t"}`, planID, true))
	if err != nil {
		return plan, err
	}

	return plan, nil
}

// DisablePlan method to disable a plan object
func (c *Client) DisablePlan(ctx context.Context, teamID, compositeID, planID string) (Plan, error) {
	plan, err := c.updatePlan(ctx, teamID, compositeID, planID, fmt.Sprintf(`{"id":"%s", "enabled":"%t"}`, planID, false))
	if err != nil {
		return plan, err
	}

	return plan, nil
}

func (c *Client) updatePlan(ctx context.Context, teamID, compositeID, planID, payLoad string) (Plan, error) {
	plan := Plan{}
	plans, err := c.GetPlans(ctx, teamID, compositeID)

	if err != nil {
		return plan, err
	}

	for _, plan := range plans {

		if plan.ID == planID {
			var jsonStr = []byte(payLoad)

			planLink, err := GetLinkByRef(plan.Links, "plans")
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
