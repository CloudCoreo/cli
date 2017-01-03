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

	"github.com/CloudCoreo/cli/client"
	"github.com/pkg/errors"
)

func TestCompositeCreateCmd(t *testing.T) {
	mockComposite := func(compositeName, teamID, compositeID string) *client.Composite {
		return &client.Composite{
			ID:     compositeID,
			TeamID: teamID,
			Name:   compositeName,
		}
	}

	tests := []struct {
		cmds  string
		desc  string
		flags []string
		args  []string
		resp  []*client.Composite
		json  bool
		err   bool
		xout  string
	}{
		{
			cmds: "coreo composite create",
			desc: "create a composite",
			flags: []string{
				"--name", "compositeName",
				"-g", "git@github.com:CloudCoreo/audit-aws.git",
			},
			resp: []*client.Composite{
				mockComposite("ID1", "Team1", "CompositeName1"),
				mockComposite("ID2", "Team2", "CompositeName2"),
			},
			xout: "-------------------  -------------------  ------------\n" +
				"    Composite ID        Composite Name       Team ID  \n" +
				"-------------------  -------------------  ------------\n" +
				"   CompositeName1            ID1              Team1   \n" +
				"-------------------  -------------------  ------------\n\n",
		},
		{
			cmds: "coreo composite create",
			desc: "create a composite",
			args: []string{""},
			err:  true,
		},
	}

	var buf bytes.Buffer
	for _, tt := range tests {
		frc := &fakeReleaseClient{composites: tt.resp}
		if tt.err {
			frc.err = errors.New("Error")
		}

		cmd := newCompositeCreateCmd(frc, &buf)
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
