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

func newTeamCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:               content.CmdTeamUse,
		Short:             content.CmdTeamShort,
		Long:              content.CmdTeamLong,
		PersistentPreRunE: setupCoreoConfig,
	}

	cmd.AddCommand(newTeamListCmd(nil, out))
	cmd.AddCommand(newTeamShowCmd(nil, out))

	return cmd
}

type teamListCmd struct {
	out    io.Writer
	client coreo.Interface
}

func newTeamListCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	teamList := &teamListCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:               content.CmdListUse,
		Short:             content.CmdTeamListShort,
		Long:              content.CmdTeamListLong,
		PersistentPreRunE: setupCoreoCredentials,
		RunE: func(cmd *cobra.Command, args []string) error {

			if teamList.client == nil {
				teamList.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			return teamList.run()
		},
	}

	return cmd
}

func (t *teamListCmd) run() error {
	teams, err := t.client.ListTeams()
	if err != nil {
		return err
	}

	b := make([]interface{}, len(teams))
	for i := range teams {
		b[i] = teams[i]
	}

	util.PrintResult(t.out, b,
		[]string{"ID", "TeamName", "TeamDescription"},
		map[string]string{
			"ID":              "Team ID",
			"TeamName":        "Team Name",
			"TeamDescription": "Team Description",
		},
		json,
		verbose)

	return nil
}
