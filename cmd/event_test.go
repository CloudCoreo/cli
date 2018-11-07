package main

import (
	"bytes"
	"testing"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/stretchr/testify/assert"
)

func TestNewEventCmd(t *testing.T) {
	var buf bytes.Buffer
	cmd := newEventCmd(&buf)
	assert.Equal(t, content.CmdEventUse, cmd.Use)
}
