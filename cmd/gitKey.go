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

func newGitKeyCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:               content.CmdGitKeyUse,
		Short:             content.CmdGitKeyShort,
		Long:              content.CmdGitKeyLong,
		PersistentPreRunE: setupCoreoConfig,
	}

	cmd.AddCommand(newGitKeyListCmd(nil, out))
	cmd.AddCommand(newGitKeyShowCmd(nil, out))
	cmd.AddCommand(newGitKeyCreateCmd(nil, out))
	cmd.AddCommand(newGitKeyDeleteCmd(nil, out))

	return cmd
}

type gitKeyListCmd struct {
	out    io.Writer
	client coreo.Interface
	teamID string
}

func newGitKeyListCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	gitKeyList := &gitKeyListCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdListUse,
		Short: content.CmdGitKeyListShort,
		Long:  content.CmdGitKeyListLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if gitKeyList.client == nil {
				gitKeyList.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			gitKeyList.teamID = teamID

			return gitKeyList.run()
		},
	}

	return cmd
}

func (t *gitKeyListCmd) run() error {
	gitKeys, err := t.client.ListGitKeys(t.teamID)
	if err != nil {
		return err
	}

	b := make([]interface{}, len(gitKeys))
	for i := range gitKeys {
		b[i] = gitKeys[i]
	}

	util.PrintResult(
		t.out,
		b,
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
