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
	"fmt"
	"time"

	"github.com/CloudCoreo/cli/pkg/command"
	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
)

const templateURL = "https://s3.amazonaws.com/cloudcoreo-files/devtime/devtime_cfn.yml"
const stackName = "cloudcoreo-events"
const version = "1"
const cloudCoreoDevTimeQueueArn = "arn:aws:sqs:us-west-2:910887748405:cloudcoreo-events-queue"
const cloudCoreoDevTimeTopicName = "cloudcoreo-events"
const cloudCoreoDevTimeMonitorRule = "cloudcoreo-events"

//SetupEventStream ...
func (c *Client) SetupEventStream(input *command.SetupEventStreamInput) error {
	regions := []string{"ap-south-1",
		"eu-west-3",
		"eu-west-2",
		"eu-west-1",
		"ap-northeast-2",
		"ap-northeast-1",
		"sa-east-1",
		"ca-central-1",
		"ap-southeast-1",
		"ap-southeast-2",
		"eu-central-1",
		"us-east-1",
		"us-east-2",
		"us-west-1",
		"us-west-2"}

	var sess *session.Session
	var err error

	if input.AwsProfile != "" {
		sess, err = session.NewSession(&aws.Config{Credentials: credentials.NewSharedCredentials(input.AwsProfilePath, input.AwsProfile)})
		if err != nil {
			return err
		}
	} else {
		sess, err = session.NewSession()
		if err != nil {
			return err
		}
	}

	if err != nil {
		return err
	}
	for _, region := range regions {
		_, err := c.checkCloudTrailForRegion(sess, region)
		if err != nil {
			return err
		}
	}

	for _, region := range regions {
		// Set up event stream
		res, err := c.checkStack(sess, region)
		if err != nil {
			return err
		}
		if res {
			err := c.updateStack(sess, region)
			if err != nil {
				return NewError(err.Error() + " in region" + region)
			}
			fmt.Println("Successfully updated stack on region " + region)
		} else {
			err := c.installStack(sess, region)
			if err != nil {
				return NewError(err.Error() + " in region" + region)
			}
			fmt.Println("Successfully installed stack on region " + region)
		}
	}

	return nil
}

func (c *Client) checkCloudTrailForRegion(sess *session.Session, region string) (bool, error) {
	// Set the Region to fetch CloudTrail information to region
	// WithRegion returns a new Config pointer that can be chained with builder
	// methods to set multiple configuration values inline without using pointers
	fmt.Println("Verifying that cloudtrail is enabled for region ", region)
	cloudTrail := cloudtrail.New(sess, aws.NewConfig().WithRegion(region))
	input := &cloudtrail.DescribeTrailsInput{}
	output, err := cloudTrail.DescribeTrails(input)
	if err != nil {
		return false, err
	}
	// Check whether IsMultiRegionTrail field is true.
	// Return true if one or more Trails qualify.
	for i := range output.TrailList {
		if *output.TrailList[i].IsMultiRegionTrail {
			return true, nil
		}
	}

	// If none, check whether there is a trail whose HomeRegion field is region.
	// Return true if one or more Trails qualify. Otherwise return false
	for i := range output.TrailList {
		if *output.TrailList[i].HomeRegion == region {
			return true, nil
		}
	}
	return false, NewError("CloudTrail is not enabled in region " + region)
}
func (c *Client) newTag(key, value string) *cloudformation.Tag {
	tag := &cloudformation.Tag{}
	tag.SetKey(key)
	tag.SetValue(value)
	return tag
}

func (c *Client) newTagList() []*cloudformation.Tag {
	tags := make([]*cloudformation.Tag, 2)
	keys := []string{"Version", "LastUpdatedTime"}
	values := []string{version, time.Now().Format(time.RFC850)}
	for i := range tags {
		tags[i] = c.newTag(keys[i], values[i])
	}
	return tags
}

func (c *Client) newParameter(key, value string) *cloudformation.Parameter {
	parameter := &cloudformation.Parameter{}
	parameter.SetParameterKey(key)
	parameter.SetParameterValue(value)
	return parameter
}

func (c *Client) newParameterList() []*cloudformation.Parameter {
	parameters := make([]*cloudformation.Parameter, 3)
	keys := []string{"CloudCoreoDevTimeQueueArn", "CloudCoreoDevTimeTopicName", "CloudCoreoDevTimeMonitorRule"}
	values := []string{cloudCoreoDevTimeQueueArn, cloudCoreoDevTimeTopicName, cloudCoreoDevTimeMonitorRule}

	for i := range parameters {
		parameters[i] = c.newParameter(keys[i], values[i])
	}
	return parameters
}

func (c *Client) newUpdateStackInput() *cloudformation.UpdateStackInput {
	input := &cloudformation.UpdateStackInput{}
	input.SetTemplateURL(templateURL)
	input.SetStackName(stackName)
	input.SetParameters(c.newParameterList())
	input.SetTags(c.newTagList())
	return input
}

func (c *Client) updateStack(sess *session.Session, region string) error {
	cloudFormation := cloudformation.New(sess, aws.NewConfig().WithRegion(region))
	_, err := cloudFormation.UpdateStack(c.newUpdateStackInput())
	return err
}

func (c *Client) newCreateStackInput() *cloudformation.CreateStackInput {
	input := &cloudformation.CreateStackInput{}
	input.SetStackName(stackName)
	input.SetTemplateURL(templateURL)
	input.SetParameters(c.newParameterList())
	input.SetTags(c.newTagList())
	input.SetOnFailure("DO_NOTHING")
	return input
}

func (c *Client) installStack(sess *session.Session, region string) error {
	cloudFormation := cloudformation.New(sess, aws.NewConfig().WithRegion(region))
	_, err := cloudFormation.CreateStack(c.newCreateStackInput())
	return err
}

func (c *Client) checkStack(sess *session.Session, region string) (bool, error) {
	stackName := "cloudcoreo-events"
	cloudFormation := cloudformation.New(sess, aws.NewConfig().WithRegion(region))
	input := &cloudformation.DescribeStacksInput{StackName: &stackName}
	output, err := cloudFormation.DescribeStacks(input)

	if err != nil {
		return false, err
	}
	if len(output.Stacks) >= 1 {
		return true, nil
	} else {
		return false, nil
	}
}
