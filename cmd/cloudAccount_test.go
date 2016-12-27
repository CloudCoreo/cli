package main

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/CloudCoreo/cli/client"
	"github.com/pkg/errors"
)

func TestCloudAccountListCmd(t *testing.T) {
	mockCloudAccount := func(cloudName, teamID, cloudID string) *client.CloudAccount {
		return &client.CloudAccount{
			ID:     cloudID,
			TeamID: teamID,
			Name:   cloudName,
		}
	}

	tests := []struct {
		cmds  string
		desc  string
		flags []string
		args  []string
		resp  []*client.CloudAccount
		json  bool
		err   bool
		xout  string
	}{
		{
			cmds:  "coreo cloud list, success",
			desc:  "get list of cloud accounts",
			flags: []string{""},
			resp: []*client.CloudAccount{
				mockCloudAccount("ID1", "Team1", "CloudName1"),
				mockCloudAccount("ID2", "Team2", "CloudName2"),
			},
			xout: "---------------------  -----------------------  ------------\n   " +
				"Cloud Account ID       Cloud Account Name       Team ID  \n" +
				"---------------------  -----------------------  ------------\n" +
				"      CloudName1                 ID1                Team1   \n\n" +
				"      CloudName2                 ID2                Team2   \n" +
				"---------------------  -----------------------  ------------\n\n",
		},
		{
			cmds: "coreo cloud list, failure",
			desc: "get list of cloud accounts",
			args: []string{""},
			err:  true,
		},
	}

	var buf bytes.Buffer
	for _, tt := range tests {
		frc := &fakeReleaseClient{cloudAccounts: tt.resp}
		if tt.err {
			frc.err = errors.New("Error")
		}

		cmd := newCloudListCmd(frc, &buf)
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
