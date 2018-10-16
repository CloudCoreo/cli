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

func TestTeamListCmd(t *testing.T) {
	mockTeam := func(teamName, teamID, teamDescription string) *client.Team {
		return &client.Team{
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
		resp  []*client.Team
		json  bool
		err   bool
		xout  string
	}{
		{
			cmds:  "coreo team list",
			desc:  "get list of teams",
			flags: []string{""},
			resp: []*client.Team{
				mockTeam("ID1", "TeamDescription1", "TeamName1"),
				mockTeam("ID2", "TeamDescription2", "TeamName2"),
			},
			xout: "------------  --------------  ---------------------\n   " +
				"Team ID       Team Name       Team Description  \n" +
				"------------  --------------  ---------------------\n     " +
				"ID1         TeamName1       TeamDescription1  \n\n     " +
				"ID1         TeamName1       TeamDescription1  " +
				"\n------------  --------------  ---------------------\n\n",
		},
		{
			cmds: "coreo team list",
			desc: "get list of teams",
			args: []string{""},
			err:  true,
		},
	}

	var buf bytes.Buffer
	for _, test := range tests {
		frc := &fakeReleaseClient{teams: test.resp}
		if test.err {
			frc.err = errors.New("Error")
		}

		cmd := newTeamListCmd(frc, &buf)
		cmd.ParseFlags(test.flags)
		err := cmd.RunE(cmd, test.args)

		if (err != nil) != test.err {
			t.Errorf("%q. expected error, got '%v'", test.desc, err)
		}

		re := regexp.MustCompile(test.xout)
		if !re.Match(buf.Bytes()) {
			t.Fatalf("%q\n\t%s:\nexpected\n\t%q\nactual\n\t%q", test.cmds, test.desc, test.xout, buf.String())
		}
		buf.Reset()
	}
}
