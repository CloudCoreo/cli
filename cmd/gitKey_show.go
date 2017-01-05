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
	"io"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type gitKeyShowCmd struct {
	out      io.Writer
	client   coreo.Interface
	teamID   string
	gitKeyID string
}

func newGitKeyShowCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	gitKeyShow := &gitKeyShowCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdShowUse,
		Short: content.CmdGitKeyShowShort,
		Long:  content.CmdGitKeyShowLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckGitKeyShowOrDeleteFlag(gitKeyShow.gitKeyID, verbose); err != nil {
				return err
			}

			if gitKeyShow.client == nil {
				gitKeyShow.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			gitKeyShow.teamID = teamID

			return gitKeyShow.run()
		},
	}
	f := cmd.Flags()

	f.StringVarP(&gitKeyShow.gitKeyID, content.CmdFlagGitKeyIDLong, "", "", content.CmdFlagGitKeyIDDescription)

	return cmd
}

func (t *gitKeyShowCmd) run() error {
	gitKey, err := t.client.ShowGitKeyByID(t.teamID, t.gitKeyID)
	if err != nil {
		return err
	}

	util.PrintResult(
		t.out,
		gitKey,
		[]string{"ID", "Name", "TeamID"},
		map[string]string{
			"ID":     "Git Key ID",
			"Name":   "Git Key Name",
			"TeamID": "Team ID",
		},
		json,
		verbose)

	return nil
}
