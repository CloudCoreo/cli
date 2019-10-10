package util

import (
	"testing"

	"github.com/CloudCoreo/cli/cmd/content"

	"github.com/stretchr/testify/assert"
)

func TestCheckCloudShowOrDeleteFlagSuccess(t *testing.T) {
	err := CheckCloudShowOrDeleteFlag("cloudID", true)
	assert.Nil(t, err, "TestCheckCloudShowOrDeleteFlagSuccess shouldn't return error")
}

func TestCheckCloudShowOrDeleteFlagFailure(t *testing.T) {
	err := CheckCloudShowOrDeleteFlag("", true)
	assert.NotNil(t, err, "TestCheckCloudShowOrDeleteFlagFailure should return error")
	assert.Equal(t, content.ErrorCloudIDRequired, err.Error())
}

func TestCheckTokenShowOrDeleteFlagSuccess(t *testing.T) {
	err := CheckTokenShowOrDeleteFlag("tokenid", true)
	assert.Nil(t, err, "TestCheckTokenShowOrDeleteFlagSuccess shouldn't return error")
}

func TestCheckTokenShowOrDeleteFlagFailure(t *testing.T) {
	err := CheckTokenShowOrDeleteFlag("", true)
	assert.NotNil(t, err, "TestCheckTokenShowOrDeleteFlagFailure should return error")
	assert.Equal(t, content.ErrorTokenIDMissing, err.Error())
}

func TestCheckAPIKeyFlagSuccess(t *testing.T) {
	res, err := CheckAPIKeyFlag("api-key", "default")
	assert.Nil(t, err, "TestCheckAPIKeyFlagSuccess shouldn't return error")
	assert.Equal(t, "api-key", res)
}

func TestCheckAPIKeyFlagFailure(t *testing.T) {
	_, err := CheckAPIKeyFlag("None", "invalid")
	assert.NotNil(t, err, "TestCheckAPIKeyFlagFailure should return error")
	assert.Equal(t, content.ErrorAPIKeyMissing, err.Error())
}

func TestCheckCloudAddFlagsFailure(t *testing.T) {
	err := CheckCloudAddFlagsForAWS("", "", "", "")
	assert.NotNil(t, err, "TestCloudAddFlagsFailure should return error")
	assert.Equal(t, "Please either provide both externalID and roleArn or the name of the new role ", err.Error())
}
