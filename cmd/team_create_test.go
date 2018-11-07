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

	"github.com/CloudCoreo/cli/pkg/command"
	"github.com/pkg/errors"
)

func TestTeamCreateCmd(t *testing.T) {
	mockTeam := func(teamName, teamID, teamDescription string) *command.Team {
		return &command.Team{
			ID:              "ID1",
			TeamDescription: "TeamDescription1",
			TeamName:        "TeamName1",
		}
	}

	tests := []struct {
		cmds  string
		desc  string
		flags []string
		args  []string
		resp  []*command.Team
		json  bool
		err   bool
		xout  string
	}{
		{
			cmds:  "coreo team show",
			desc:  "get a particular team",
			flags: []string{"--name", "testName", "--description", "teamDescription"},
			resp: []*command.Team{
				mockTeam("ID1", "TeamDescription1", "TeamName1"),
				mockTeam("ID2", "TeamDescription2", "TeamName2"),
			},
			xout: "------------  --------------  ---------------------\n   " +
				"Team ID       Team Name       Team Description  \n------------  " +
				"--------------  ---------------------\n     " +
				"ID1         TeamName1       TeamDescription1 " +
				" \n------------  --------------  ---------------------\n\n",
		},
		{
			cmds: "coreo team create",
			desc: "get a particular team, returns error",
			args: []string{""},
			err:  true,
		},
	}

	var buf bytes.Buffer
	for _, tt := range tests {
		frc := &fakeReleaseClient{teams: tt.resp}
		if tt.err {
			frc.err = errors.New("Error")
		}

		cmd := newTeamCreateCmd(frc, &buf)
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
