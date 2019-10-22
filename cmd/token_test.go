// Copyright © 2016 Paul Allen <paul@cloudcoreo.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/stretchr/testify/assert"

	"github.com/pkg/errors"
)

func TestTokenListCmd(t *testing.T) {

	tests := []struct {
		cmds  string
		desc  string
		flags []string
		args  []string
		json  bool
		err   bool
		xout  string
	}{
		{
			cmds:  "coreo team list",
			desc:  "get list of teams",
			flags: []string{""},
			xout:  "Tokens are deprecated, only csp token is required` \n",
		},
	}

	var buf bytes.Buffer
	for _, tt := range tests {
		frc := &fakeReleaseClient{}
		if tt.err {
			frc.err = errors.New("Error")
		}

		cmd := newTokenListCmd(frc, &buf)
		cmd.ParseFlags(tt.flags)
		err := cmd.RunE(cmd, tt.args)

		if (err != nil) != tt.err {
			t.Errorf("%q. expected error, got '%v'", tt.desc, err)
		}

		re := regexp.MustCompile(tt.xout)
		if !re.Match(buf.Bytes()) {
			t.Fatalf("%q\n\t%s:\nexpected\n\t%q\nactual\n\t%q", tt.cmds, tt.desc, tt.xout, buf.String())
		}
		buf.Reset()
	}
}

func TestNewTokenCmd(t *testing.T) {
	var buf bytes.Buffer
	cmd := newTokenCmd(&buf)
	assert.Equal(t, content.CmdTokenUse, cmd.Use)
}
