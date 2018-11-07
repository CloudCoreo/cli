package client

import (
	"testing"

	"github.com/jharlap/httpstub"
	"github.com/stretchr/testify/assert"
)

func TestNewTagListSuccess(t *testing.T) {
	ts := httpstub.New()
	defer ts.Close()
	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	tag := client.newTagList()
	assert.Equal(t, "Version", *tag[0].Key)
	assert.Equal(t, "LastUpdatedTime", *tag[0].Value)
}

func TestNewParameterListSuccess(t *testing.T) {
	ts := httpstub.New()
	defer ts.Close()
	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	parameterList := client.newParameterList()
	assert.Equal(t, "CloudCoreoDevTimeQueueArn", *parameterList[0].ParameterKey)
	assert.Equal(t, cloudCoreoDevTimeQueueArn, *parameterList[0].ParameterValue)
}

func TestNewCreateStackInput(t *testing.T) {
	ts := httpstub.New()
	defer ts.Close()
	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	input := client.newCreateStackInput()
	assert.Equal(t, stackName, *input.StackName)
	assert.Equal(t, templateURL, *input.TemplateURL)
	assert.Equal(t, "DO_NOTHING", *input.OnFailure)
}

func TestNewUpdateStackInput(t *testing.T) {
	ts := httpstub.New()
	defer ts.Close()
	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	input := client.newUpdateStackInput()
	assert.Equal(t, stackName, *input.StackName)
	assert.Equal(t, templateURL, *input.TemplateURL)
}
