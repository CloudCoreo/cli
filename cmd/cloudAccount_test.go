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

func TestCloudAccountListCmd(t *testing.T) {
	mockCloudAccount := func(cloudName, teamID, cloudID, accountID string) *client.CloudAccount {
		return &client.CloudAccount{
			ID:        cloudID,
			CloudInfo: client.CloudInfo{Name: cloudName},
			AccountID: accountID,
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
				mockCloudAccount("ID1", "Team1", "CloudName1", "AccountID1"),
				mockCloudAccount("ID2", "Team2", "CloudName2", "AccountID2"),
			},
			xout: "---------------  -----------------------  ---------------------  ------------  ---------  -------------\n       " +
				"ID           Cloud Account Name       Cloud account ID       IsDraft       Tags       Provider  \n" +
				"---------------  -----------------------  ---------------------  ------------  ---------  -------------\n" +
				`   CloudName1              ID1                  AccountID1           false         \[\]                  \n\n` +
				`   CloudName2              ID2                  AccountID2           false         \[\]                  \n` +
				"---------------  -----------------------  ---------------------  ------------  ---------  -------------\n\n",
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
