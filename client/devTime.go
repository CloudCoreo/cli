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
	"fmt"

	"github.com/CloudCoreo/cli/client/content"
)

// DevTime struct for api payload
type DevTime struct {
	Context    string `json:"context"`
	Task       string `json:"task"`
	DevTimeURL string `json:"devTimeUrl"`
	Links      []Link `json:"links"`
	DevTimeID  string `json:"devTimeId"`
}

// DevTimeStatus struct for status
type DevTimeStatus struct {
	Status struct {
		RunningState string `json:"runningState"`
		EngineState  string `json:"engineState"`
	} `json:"status"`
}

// DevTimeResults struct for results
type DevTimeResults struct {
	Results interface{} `json:"results"`
}

// CreateDevTime method to create a ProxyTask object
func (c *Client) CreateDevTime(ctx context.Context, teamID, context, task string) (*DevTime, error) {
	devTime := &DevTime{}
	teams, err := c.GetTeams(ctx)

	if err != nil {
		return nil, NewError(err.Error())
	}

	for _, team := range teams {
		if team.ID == teamID {
			proxyTaskLink, err := GetLinkByRef(team.Links, "devtime")
			if err != nil {
				return nil, NewError(err.Error())
			}
			proxyTaskPayLoad := fmt.Sprintf(`{"context":"%s","task":"%s","teamId":"%s"}`, context, task, teamID)
			var jsonStr = []byte(proxyTaskPayLoad)
			err = c.Do(ctx, "POST", proxyTaskLink.Href, bytes.NewBuffer(jsonStr), &devTime)
			if err != nil {
				return nil, NewError(err.Error())
			}

			break
		}
	}

	if devTime.DevTimeID == "" {
		return nil, NewError(fmt.Sprintf(content.ErrorFailedToCreateDevTime, teamID))
	}

	return devTime, nil
}

// GetDevTimes method to devTimes info array object
func (c *Client) GetDevTimes(ctx context.Context, teamID string) ([]*DevTime, error) {
	devTimes := []*DevTime{}
	teams, err := c.GetTeams(ctx)

	if err != nil {
		return nil, NewError(err.Error())
	}

	for _, team := range teams {
		if team.ID == teamID {
			devTimeLink, e := GetLinkByRef(team.Links, "devtime")

			if e != nil {
				return nil, NewError(e.Error())
			}

			e = c.Do(ctx, "GET", devTimeLink.Href, nil, &devTimes)
			if e != nil {
				return nil, NewError(e.Error())
			}
		}
	}

	if len(devTimes) == 0 {
		return nil, NewError(fmt.Sprintf(content.ErrorNoDevTimesFound, teamID))
	}

	return devTimes, nil
}

// GetDevTimeResults method to get devTimes results
func (c *Client) GetDevTimeResults(ctx context.Context, teamID, devTimeID string) (*DevTimeResults, error) {
	devTimes, err := c.GetDevTimes(ctx, teamID)
	if err != nil {
		return nil, err
	}

	devTime := &DevTime{}
	for _, d := range devTimes {
		if d.DevTimeID == devTimeID {
			devTime = d
			break
		}
	}

	if devTime.DevTimeID == "" {
		return nil, NewError(fmt.Sprintf(content.ErrorNoDevTimeWithIDFound, devTimeID, teamID))
	}

	resultsLink, err := GetLinkByRef(devTime.Links, "results")
	if err != nil {
		return nil, err
	}

	devTimeResults := &DevTimeResults{}
	err = c.Do(ctx, "GET", resultsLink.Href, nil, &devTimeResults)
	if err != nil {
		return nil, err
	}

	return devTimeResults, nil
}

// GetDevTimeStatus method to get devTimes status
func (c *Client) GetDevTimeStatus(ctx context.Context, teamID, devTimeID string) (*DevTimeStatus, error) {
	devTimes, err := c.GetDevTimes(ctx, teamID)
	if err != nil {
		return nil, err
	}

	devTime := &DevTime{}
	for _, d := range devTimes {
		if d.DevTimeID == devTimeID {
			devTime = d
			break
		}
	}

	if devTime.DevTimeID == "" {
		return nil, NewError(fmt.Sprintf(content.ErrorNoDevTimeWithIDFound, devTimeID, teamID))
	}

	resultsLink, err := GetLinkByRef(devTime.Links, "status")
	if err != nil {
		return nil, err
	}
	devTimeStatus := &DevTimeStatus{}
	err = c.Do(ctx, "GET", resultsLink.Href, nil, &devTimeStatus)
	if err != nil {
		return nil, err
	}

	return devTimeStatus, nil
}

// actionDevTime method to devTimes info array object
func (c *Client) actionDevTime(ctx context.Context, teamID, devTimeID, action string) error {
	devTimes, err := c.GetDevTimes(ctx, teamID)
	if err != nil {
		return err
	}

	devTime := &DevTime{}
	for _, d := range devTimes {
		if d.DevTimeID == devTimeID {
			devTime = d
			break
		}
	}

	if devTime.DevTimeID == "" {
		return NewError(fmt.Sprintf(content.ErrorNoDevTimeWithIDFound, devTimeID, teamID))
	}

	resultsLink, err := GetLinkByRef(devTime.Links, action)
	if err != nil {
		return err
	}

	err = c.Do(ctx, "GET", resultsLink.Href, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

// StartDevTime start devtime
func (c *Client) StartDevTime(ctx context.Context, teamID, devTimeID string) error {
	return c.actionDevTime(ctx, teamID, devTimeID, "start")
}

// StopDevTime stop devtime
func (c *Client) StopDevTime(ctx context.Context, teamID, devTimeID string) error {
	return c.actionDevTime(ctx, teamID, devTimeID, "stop")
}
