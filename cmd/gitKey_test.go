package main

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/CloudCoreo/cli/client"
	"github.com/pkg/errors"
)

func TestGitKeyListCmd(t *testing.T) {
	mockGitKey := func(gitKeyName, teamID, gitKeyID string) *client.GitKey {
		return &client.GitKey{
			ID:     gitKeyID,
			TeamID: teamID,
			Name:   gitKeyName,
		}
	}

	tests := []struct {
		cmds  string
		desc  string
		flags []string
		args  []string
		resp  []*client.GitKey
		json  bool
		err   bool
		xout  string
	}{
		{
			cmds:  "coreo git-key list",
			desc:  "get list of git key",
			flags: []string{""},
			resp: []*client.GitKey{
				mockGitKey("ID1", "Team1", "GitKeyName1"),
				mockGitKey("ID2", "Team2", "GitKeyName2"),
			},
			xout: "----------------  -----------------  ------------\n   " +
				"Git Key ID        Git Key Name       Team ID  \n" +
				"----------------  -----------------  ------------\n   " +
				"GitKeyName1           ID1             Team1   \n\n   " +
				"GitKeyName2           ID2             Team2   \n" +
				"----------------  -----------------  ------------\n\n",
		},
		{
			cmds: "coreo git-key lis",
			desc: "get list of git key",
			args: []string{""},
			err:  true,
		},
	}

	var buf bytes.Buffer
	for _, tt := range tests {
		frc := &fakeReleaseClient{gitKeys: tt.resp}
		if tt.err {
			frc.err = errors.New("Error")
		}

		cmd := newGitKeyListCmd(frc, &buf)
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
