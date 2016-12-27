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
	for _, tt := range tests {
		frc := &fakeReleaseClient{teams: tt.resp}
		if tt.err {
			frc.err = errors.New("Error")
		}

		cmd := newTeamListCmd(frc, &buf)
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
