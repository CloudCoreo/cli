package aws

import (
	"testing"

	"github.com/CloudCoreo/cli/client"

	"github.com/CloudCoreo/cli/pkg/command"

	"github.com/stretchr/testify/assert"
)

func TestNewTagListSuccess(t *testing.T) {
	setup := NewSetupService(&NewServiceInput{})
	input := &command.SetupEventStreamInput{
		Config: &client.EventStreamConfig{
			AWSEventStreamConfig: client.AWSEventStreamConfig{
				TemplateURL:     "fake-url",
				TopicName:       "fake-topic",
				StackName:       "fake-stack",
				DevtimeQueueArn: "fake-devtime-queue-arn",
				Version:         "fake-version",
				MonitorRule:     "fake-monitor-rule",
			},
		},
	}
	tag := setup.newTagList(input.Config)
	assert.Equal(t, "Version", *tag[0].Key)
	assert.Equal(t, "LastUpdatedTime", *tag[1].Key)
}

func TestNewParameterListSuccess(t *testing.T) {
	setup := NewSetupService(&NewServiceInput{})
	input := &command.SetupEventStreamInput{
		Config: &client.EventStreamConfig{
			AWSEventStreamConfig: client.AWSEventStreamConfig{
				TemplateURL:     "fake-url",
				TopicName:       "fake-topic",
				StackName:       "fake-stack",
				DevtimeQueueArn: "fake-devtime-queue-arn",
				Version:         "fake-version",
				MonitorRule:     "fake-monitor-rule",
			},
		},
	}
	parameterList := setup.newParameterList(input.Config)
	assert.Equal(t, "CloudCoreoDevTimeQueueArn", *parameterList[0].ParameterKey)
	assert.Equal(t, input.Config.DevtimeQueueArn, *parameterList[0].ParameterValue)
}

func TestNewCreateStackInputSuccess(t *testing.T) {
	setup := NewSetupService(&NewServiceInput{})
	input := &command.SetupEventStreamInput{
		Config: &client.EventStreamConfig{
			AWSEventStreamConfig: client.AWSEventStreamConfig{
				TemplateURL:     "fake-url",
				TopicName:       "fake-topic",
				StackName:       "fake-stack",
				DevtimeQueueArn: "fake-devtime-queue-arn",
				Version:         "fake-version",
				MonitorRule:     "fake-monitor-rule",
			},
		},
	}
	createStackInput := setup.newCreateStackInput(input.Config)
	assert.Equal(t, input.Config.StackName, *createStackInput.StackName)
	assert.Equal(t, input.Config.TemplateURL, *createStackInput.TemplateURL)
	assert.Equal(t, "DO_NOTHING", *createStackInput.OnFailure)
}

func TestNewUpdateStackInputSuccess(t *testing.T) {
	setup := NewSetupService(&NewServiceInput{})
	input := &command.SetupEventStreamInput{
		Config: &client.EventStreamConfig{
			AWSEventStreamConfig: client.AWSEventStreamConfig{
				TemplateURL:     "fake-url",
				TopicName:       "fake-topic",
				StackName:       "fake-stack",
				DevtimeQueueArn: "fake-devtime-queue-arn",
				Version:         "fake-version",
				MonitorRule:     "fake-monitor-rule",
			},
		},
	}

	updateStackInput := setup.newUpdateStackInput(input.Config)
	assert.Equal(t, input.Config.StackName, *updateStackInput.StackName)
	assert.Equal(t, input.Config.TemplateURL, *updateStackInput.TemplateURL)
}
