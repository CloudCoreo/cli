package aws

import (
	"fmt"
	"time"

	"github.com/CloudCoreo/cli/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
)

//SetupService  is the struct implements CloudProvider interface for aws
type SetupService struct {
	awsProfilePath     string
	awsProfile         string
	ignoreMissingTrail bool
}

//NewSetupService returns a pointer to a setup struct object
func NewSetupService(input *NewServiceInput) *SetupService {
	return &SetupService{
		awsProfile:         input.AwsProfile,
		awsProfilePath:     input.AwsProfilePath,
		ignoreMissingTrail: input.IgnoreMissingTrails,
	}
}

func (a *SetupService) newSession() (*session.Session, error) {
	var sess *session.Session
	var err error
	if a.awsProfile != "" {
		if a.awsProfilePath != "" {
			sess, err = session.NewSessionWithOptions(session.Options{Profile: a.awsProfile, SharedConfigFiles: []string{a.awsProfilePath}, SharedConfigState: session.SharedConfigEnable})
		} else {
			sess, err = session.NewSessionWithOptions(session.Options{Profile: a.awsProfile, SharedConfigState: session.SharedConfigEnable})
		}
		if err != nil {
			return nil, err
		}
	} else {
		sess, err = session.NewSession()
		if err != nil {
			return nil, err
		}
	}
	return sess, nil
}

//SetupEventStream sets up event stream for aws account
func (a *SetupService) SetupEventStream(input *client.EventStreamConfig) error {
	regions := input.Regions

	sess, err := a.newSession()
	if err != nil {
		return err
	}

	for _, region := range regions {
		// Check CloudTrail
		_, err := a.checkCloudTrailForRegion(sess, region)
		if err != nil {
			if a.ignoreMissingTrail {
				fmt.Println("CloudTrail is not enabled in region " + region + ". Skip event stream setup for this region.")
				continue
			} else {
				return err
			}
		}

		// Set up event stream
		res := a.checkStack(sess, region, input)
		if res {
			fmt.Println("Updating stack in " + region)
			err := a.updateStack(sess, region, input)
			if err != nil {
				return client.NewError(err.Error() + " in region" + region)
			}
			fmt.Println("Successfully updated stack on region " + region)
		} else {
			fmt.Println("Installing stack in " + region)
			err := a.installStack(sess, region, input)
			if err != nil {
				return client.NewError(err.Error() + " in region" + region)
			}
			fmt.Println("Successfully installed stack on region " + region)
		}
	}

	return nil
}

func (a *SetupService) checkCloudTrailForRegion(sess *session.Session, region string) (bool, error) {
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
func (a *SetupService) newTag(key, value string) *cloudformation.Tag {
	tag := &cloudformation.Tag{}
	tag.SetKey(key)
	tag.SetValue(value)
	return tag
}

func (a *SetupService) newTagList(config *client.EventStreamConfig) []*cloudformation.Tag {
	tags := make([]*cloudformation.Tag, 2)
	keys := []string{"Version", "LastUpdatedTime"}
	//Do not put comma in tag values due to aws internal bug.
	values := []string{config.Version, time.Now().Format(time.RFC3339)}
	for i := range tags {
		tags[i] = a.newTag(keys[i], values[i])
	}
	return tags
}

func (a *SetupService) newParameter(key, value string) *cloudformation.Parameter {
	parameter := &cloudformation.Parameter{}
	parameter.SetParameterKey(key)
	parameter.SetParameterValue(value)
	return parameter
}

func (a *SetupService) newParameterList(config *client.EventStreamConfig) []*cloudformation.Parameter {
	parameters := make([]*cloudformation.Parameter, 3)
	keys := []string{"CloudCoreoDevTimeQueueArn", "CloudCoreoDevTimeTopicName", "CloudCoreoDevTimeMonitorRule"}
	values := []string{config.DevtimeQueueArn, config.TopicName, config.MonitorRule}

	for i := range parameters {
		parameters[i] = a.newParameter(keys[i], values[i])
	}
	return parameters
}

func (a *SetupService) newUpdateStackInput(config *client.EventStreamConfig) *cloudformation.UpdateStackInput {
	input := &cloudformation.UpdateStackInput{}
	input.SetTemplateURL(config.TemplateURL)
	input.SetStackName(config.StackName)
	input.SetParameters(a.newParameterList(config))
	input.SetTags(a.newTagList(config))
	return input
}

func (a *SetupService) updateStack(sess *session.Session, region string, config *client.EventStreamConfig) error {
	cloudFormation := cloudformation.New(sess, aws.NewConfig().WithRegion(region))
	_, err := cloudFormation.UpdateStack(a.newUpdateStackInput(config))
	return err
}

func (a *SetupService) newCreateStackInput(config *client.EventStreamConfig) *cloudformation.CreateStackInput {
	input := &cloudformation.CreateStackInput{}
	input.SetStackName(config.StackName)
	input.SetTemplateURL(config.TemplateURL)
	input.SetParameters(a.newParameterList(config))
	input.SetTags(a.newTagList(config))
	input.SetOnFailure("DO_NOTHING")
	return input
}

func (a *SetupService) installStack(sess *session.Session, region string, config *client.EventStreamConfig) error {
	cloudFormation := cloudformation.New(sess, aws.NewConfig().WithRegion(region))
	_, err := cloudFormation.CreateStack(a.newCreateStackInput(config))
	return err
}

func (a *SetupService) checkStack(sess *session.Session, region string, config *client.EventStreamConfig) bool {
	cloudFormation := cloudformation.New(sess, aws.NewConfig().WithRegion(region))
	input := &cloudformation.DescribeStacksInput{StackName: &config.StackName}
	output, err := cloudFormation.DescribeStacks(input)
	if err != nil {
		return false
	}
	if len(output.Stacks) >= 1 {
		return true
	}

	return false
}
