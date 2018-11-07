package main

import (
	"bytes"
	"testing"

	"github.com/CloudCoreo/cli/_vendor-20180913135234/github.com/stretchr/testify/assert"
	"github.com/CloudCoreo/cli/cmd/content"
)

func TestNewEventCmd(t *testing.T) {
	var buf bytes.Buffer
	cmd := newEventCmd(&buf)
	assert.Equal(t, content.CmdEventUse, cmd.Use)
}
