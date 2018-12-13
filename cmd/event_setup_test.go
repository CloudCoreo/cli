package main

import (
	"bytes"
	"testing"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestNewEventSetupCmd(t *testing.T) {
	// frc := &fakeReleaseClient{}
	// cloud := &fakeCloudProvider{}
	var buf bytes.Buffer

	tests := []struct {
		flags    []string
		args     []string
		err      string
		cloudErr string
		desc     string
	}{
		{
			flags: []string{
				"--aws-profile", "default",
				"--aws-profile-path", "$HOME",
			},
			desc: "event stream setup without cloud-id",
			err:  content.ErrorCloudIDRequired,
		},
		{
			flags: []string{
				"--cloud-id", "cloud-id",
				"--aws-profile", "default",
				"--aws-profile-path", "$HOME",
			},
			desc: "event stream setup with cloud-id",
		},
		{
			flags: []string{
				"--cloud-id", "cloud-id",
				"--aws-profile", "default",
				"--aws-profile-path", "$HOME",
			},
			desc: "event stream setup with cloud-id",
			err:  "error",
		},
		{
			flags: []string{
				"--cloud-id", "cloud-id",
				"--aws-profile", "default",
				"--aws-profile-path", "$HOME",
			},
			desc:     "event stream setup with cloud-id",
			cloudErr: "error",
		},
	}

	for _, tt := range tests {
		frc := &fakeReleaseClient{}
		if tt.err != "" {
			frc.err = errors.New(tt.err)
		}
		cloud := &fakeCloudProvider{}
		if tt.cloudErr != "" {
			cloud.err = errors.New(tt.cloudErr)
		}
		cmd := newEventSetupCmd(frc, cloud, &buf)
		err := cmd.ParseFlags(tt.flags)
		assert.Nil(t, err)
		err = cmd.RunE(cmd, []string{})
		if (err == nil) && (tt.err != "") {
			t.Errorf("%q. expected error, got '%v'", tt.desc, err.Error())
		}
		if err != nil {
			if tt.err != "" {
				assert.Equal(t, tt.err, err.Error())
			}
			if tt.cloudErr != "" {
				assert.Equal(t, tt.cloudErr, err.Error())
			}
		}
		buf.Reset()
	}
	/*
		cmd := newEventSetupCmd(frc, cloud, &buf)
		flags := []string{
			"--aws-profile", "default",
			"--aws-profile-path", "$HOME",
		}


		cmd.ParseFlags(flags)
		err := cmd.RunE(cmd, []string{})
		assert.NotNil(t, err)
		assert.Equal(t, content.ErrorCloudIDRequired, err.Error())
		assert.Equal(t, content.CmdEventSetupUse, cmd.Use)
		assert.Equal(t, content.CmdEventSetupShort, cmd.Short)
	*/
}
