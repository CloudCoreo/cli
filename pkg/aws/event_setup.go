package aws

import (
	"fmt"
	"time"

	"github.com/CloudCoreo/cli/client"
	"github.com/CloudCoreo/cli/pkg/command"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
)

//Setup is the struct implements CloudProvider interface for aws
type Setup struct {
	filename string
	profile  string
}

//NewAwsSetup returns a pointer to a setup struct object
func NewAwsSetup() *Setup {
	return &Setup{}
}

//SetupEventStream sets up event stream for aws account
func (a *Setup) SetupEventStream(input *command.SetupEventStreamInput) error {
	regions := input.Config.Regions

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

	for _, region := range regions {
		_, err := a.checkCloudTrailForRegion(sess, region)
		if err != nil {
			return err
		}
	}

	for _, region := range regions {
		// Set up event stream
		res, err := a.checkStack(sess, region, input.Config)
		if err != nil {
			return err
		}
		if res {
			err := a.updateStack(sess, region, input.Config)
			if err != nil {
				return client.NewError(err.Error() + " in region" + region)
			}
			fmt.Println("Successfully updated stack on region " + region)
		} else {
			err := a.installStack(sess, region, input.Config)
			if err != nil {
				return client.NewError(err.Error() + " in region" + region)
			}
			fmt.Println("Successfully installed stack on region " + region)
		}
	}

	return nil
}

func (a *Setup) checkCloudTrailForRegion(sess *session.Session, region string) (bool, error) {
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
	return false, client.NewError("CloudTrail is not enabled in region " + region)
}
func (a *Setup) newTag(key, value string) *cloudformation.Tag {
	tag := &cloudformation.Tag{}
	tag.SetKey(key)
	tag.SetValue(value)
	return tag
}

func (a *Setup) newTagList(config *client.EventStreamConfig) []*cloudformation.Tag {
	tags := make([]*cloudformation.Tag, 2)
	keys := []string{"Version", "LastUpdatedTime"}
	values := []string{config.Version, time.Now().Format(time.RFC850)}
	for i := range tags {
		tags[i] = a.newTag(keys[i], values[i])
	}
	return tags
}

func (a *Setup) newParameter(key, value string) *cloudformation.Parameter {
	parameter := &cloudformation.Parameter{}
	parameter.SetParameterKey(key)
	parameter.SetParameterValue(value)
	return parameter
}

func (a *Setup) newParameterList(config *client.EventStreamConfig) []*cloudformation.Parameter {
	parameters := make([]*cloudformation.Parameter, 3)
	keys := []string{"CloudCoreoDevTimeQueueArn", "CloudCoreoDevTimeTopicName", "CloudCoreoDevTimeMonitorRule"}
	values := []string{config.DevtimeQueueArn, config.TopicName, config.MonitorRule}

	for i := range parameters {
		parameters[i] = a.newParameter(keys[i], values[i])
	}
	return parameters
}

func (a *Setup) newUpdateStackInput(config *client.EventStreamConfig) *cloudformation.UpdateStackInput {
	input := &cloudformation.UpdateStackInput{}
	input.SetTemplateURL(config.TemplateURL)
	input.SetStackName(config.StackName)
	input.SetParameters(a.newParameterList(config))
	input.SetTags(a.newTagList(config))
	return input
}

func (a *Setup) updateStack(sess *session.Session, region string, config *client.EventStreamConfig) error {
	cloudFormation := cloudformation.New(sess, aws.NewConfig().WithRegion(region))
	_, err := cloudFormation.UpdateStack(a.newUpdateStackInput(config))
	return err
}

func (a *Setup) newCreateStackInput(config *client.EventStreamConfig) *cloudformation.CreateStackInput {
	input := &cloudformation.CreateStackInput{}
	input.SetStackName(config.StackName)
	input.SetTemplateURL(config.TemplateURL)
	input.SetParameters(a.newParameterList(config))
	input.SetTags(a.newTagList(config))
	input.SetOnFailure("DO_NOTHING")
	return input
}

func (a *Setup) installStack(sess *session.Session, region string, config *client.EventStreamConfig) error {
	cloudFormation := cloudformation.New(sess, aws.NewConfig().WithRegion(region))
	_, err := cloudFormation.CreateStack(a.newCreateStackInput(config))
	return err
}

func (a *Setup) checkStack(sess *session.Session, region string, config *client.EventStreamConfig) (bool, error) {
	cloudFormation := cloudformation.New(sess, aws.NewConfig().WithRegion(region))
	input := &cloudformation.DescribeStacksInput{StackName: &config.StackName}
	output, err := cloudFormation.DescribeStacks(input)

	if err != nil {
		return false, err
	}
	if len(output.Stacks) >= 1 {
		return true, nil
	}

	return false, nil
}
