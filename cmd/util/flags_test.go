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

func TestTeamIDFlagSuccess(t *testing.T) {
	res, err := CheckTeamIDFlag("teamID", "default", true)
	assert.Nil(t, err, "TestTeamIDFlagSuccess shouldn't return error")
	assert.Equal(t, "teamID", res)
}

func TestTeamIDFlagFailure(t *testing.T) {
	_, err := CheckTeamIDFlag(content.None, "invalid", true)
	assert.NotNil(t, err, "TestTeamIDFlagFailure should return error")
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

func TestCheckTeamAddFlagsSuccess(t *testing.T) {
	err := CheckTeamAddFlags("teamName", "teamDescription")
	assert.Nil(t, err, "TestCheckTeamAddFlagsSuccess shouldn't return error")
}

func TestCheckTeamAddFlagsFailureNoTeamName(t *testing.T) {
	err := CheckTeamAddFlags("", "teamDescription")
	assert.NotNil(t, err, "TestCheckTeamAddFlagsFailureNoTeamName should return error")
	assert.Equal(t, content.ErrorTeamNameRequired, err.Error())
}

func TestCheckTeamAddFlagsFailureNoTeamDescription(t *testing.T) {
	err := CheckTeamAddFlags("teamName", "")
	assert.NotNil(t, err, "TestCheckTeamAddFlagsFailureNoTeamDescription should return error")
	assert.Equal(t, content.ErrorTeamDescriptionRequired, err.Error())
}

func TestCheckCloudAddFlagsFailure(t *testing.T) {
	err := CheckCloudAddFlagsForAWS("", "", "", "")
	assert.NotNil(t, err, "TestCloudAddFlagsFailure should return error")
	assert.Equal(t, "Please either provide both externalID and roleArn or the name of the new role ", err.Error())
}
