package main

import (
	"bytes"
	"testing"

	"github.com/CloudCoreo/cli/_vendor-20180913135234/github.com/stretchr/testify/assert"
	"github.com/CloudCoreo/cli/cmd/content"
)

func TestNewEventSetupCmd(t *testing.T) {
	frc := &fakeReleaseClient{}
	var buf bytes.Buffer
	cmd := newEventSetupCmd(frc, &buf)
	flags := []string{
		"--aws-profile", "default",
		"--aws-profile-path", "$HOME",
	}
	cmd.ParseFlags(flags)
	assert.Equal(t, content.CmdEventSetupUse, cmd.Use)
	assert.Equal(t, content.CmdEventSetupShort, cmd.Short)
}
