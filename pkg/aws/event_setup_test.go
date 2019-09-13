package aws

import (
	"testing"

	"github.com/CloudCoreo/cli/client"

	"github.com/stretchr/testify/assert"
)

func TestNewTagListSuccess(t *testing.T) {
	setup := NewSetupService(&NewServiceInput{})
	input := client.EventStreamConfig{
		AWSEventStreamConfig: client.AWSEventStreamConfig{
			TemplateURL:     "fake-url",
			TopicName:       "fake-topic",
			StackName:       "fake-stack",
			DevtimeQueueArn: "fake-devtime-queue-arn",
			Version:         "fake-version",
			MonitorRule:     "fake-monitor-rule",
		},
	}
	tag := setup.newTagList(&input)
	assert.Equal(t, "Version", *tag[0].Key)
	assert.Equal(t, "LastUpdatedTime", *tag[1].Key)
}

func TestNewParameterListSuccess(t *testing.T) {
	setup := NewSetupService(&NewServiceInput{})
	input := client.EventStreamConfig{
		AWSEventStreamConfig: client.AWSEventStreamConfig{
			TemplateURL:     "fake-url",
			TopicName:       "fake-topic",
			StackName:       "fake-stack",
			DevtimeQueueArn: "fake-devtime-queue-arn",
			Version:         "fake-version",
			MonitorRule:     "fake-monitor-rule",
		},
	}
	parameterList := setup.newParameterList(&input)
	assert.Equal(t, "CloudCoreoDevTimeQueueArn", *parameterList[0].ParameterKey)
	assert.Equal(t, input.DevtimeQueueArn, *parameterList[0].ParameterValue)
}

func TestNewCreateStackInputSuccess(t *testing.T) {
	setup := NewSetupService(&NewServiceInput{})
	input := client.EventStreamConfig{
		AWSEventStreamConfig: client.AWSEventStreamConfig{
			TemplateURL:     "fake-url",
			TopicName:       "fake-topic",
			StackName:       "fake-stack",
			DevtimeQueueArn: "fake-devtime-queue-arn",
			Version:         "fake-version",
			MonitorRule:     "fake-monitor-rule",
		},
	}
	createStackInput := setup.newCreateStackInput(&input)
	assert.Equal(t, input.StackName, *createStackInput.StackName)
	assert.Equal(t, input.TemplateURL, *createStackInput.TemplateURL)
	assert.Equal(t, "DO_NOTHING", *createStackInput.OnFailure)
}

func TestNewUpdateStackInputSuccess(t *testing.T) {
	setup := NewSetupService(&NewServiceInput{})
	input := client.EventStreamConfig{
		AWSEventStreamConfig: client.AWSEventStreamConfig{
			TemplateURL:     "fake-url",
			TopicName:       "fake-topic",
			StackName:       "fake-stack",
			DevtimeQueueArn: "fake-devtime-queue-arn",
			Version:         "fake-version",
			MonitorRule:     "fake-monitor-rule",
		},
	}

	updateStackInput := setup.newUpdateStackInput(&input)
	assert.Equal(t, input.StackName, *updateStackInput.StackName)
	assert.Equal(t, input.TemplateURL, *updateStackInput.TemplateURL)
}
