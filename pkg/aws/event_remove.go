package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/cloudformation"

	"github.com/aws/aws-sdk-go/service/sns"

	"github.com/CloudCoreo/cli/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

//RemoveService contains info needed for AWS event stream removal
type RemoveService struct {
	awsProfilePath string
	awsProfile     string
}

// NewRemoveService returns an instance of RemoveService
func NewRemoveService(input *NewServiceInput) *RemoveService {
	return &RemoveService{
		awsProfile:     input.AwsProfile,
		awsProfilePath: input.AwsProfilePath,
	}
}

func (a *RemoveService) newSession() (*session.Session, error) {
	var sess *session.Session
	var err error

	if a.awsProfile != "" {
		sess, err = session.NewSessionWithOptions(session.Options{Profile: a.awsProfile, SharedConfigFiles: []string{a.awsProfilePath}})
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

func (a *RemoveService) snsPublish(sess *session.Session, arnType, region, cloudAccountID, topicName string) error {
	svc := sns.New(sess, aws.NewConfig().WithRegion(region))
	topicArn := fmt.Sprintf("arn:%s:sns:%s:%s:%s", arnType, region, cloudAccountID, topicName)
	publishInput := &sns.PublishInput{
		Message:  aws.String("UnsubscribeConfirmation"),
		TopicArn: aws.String(topicArn),
	}
	_, err := svc.Publish(publishInput)
	return err
}

//RemoveEventStream perform the same function as event stream removal script
func (a *RemoveService) RemoveEventStream(input *client.EventRemoveConfig) error {
	regions := input.Regions
	sess, err := a.newSession()
	if err != nil {
		return err
	}
	fmt.Println("Deactivating devTime for cloud account", input.CloudAccountID)
	for _, region := range regions {
		err := a.snsPublish(sess, input.ArnType, region, input.CloudAccountID, input.TopicName)
		if err != nil {
			fmt.Println(err.Error())
		}

		// Delete stack
		err = a.deleteStack(sess, region, input.StackName)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return nil
}

func (a *RemoveService) deleteStack(sess *session.Session, region, stackName string) error {
	fmt.Println("Deleting", stackName, "on", region)
	cloudFormation := cloudformation.New(sess, aws.NewConfig().WithRegion(region))
	deleteStackInput := &cloudformation.DeleteStackInput{
		StackName: aws.String(stackName),
	}
	_, err := cloudFormation.DeleteStack(deleteStackInput)
	return err
}
