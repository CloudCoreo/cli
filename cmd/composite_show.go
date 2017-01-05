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

type compositeShowCmd struct {
	out         io.Writer
	client      coreo.Interface
	teamID      string
	compositeID string
}

func newCompositeShowCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	compositeShow := &compositeShowCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdShowUse,
		Short: content.CmdCompositeShowShort,
		Long:  content.CmdCompositeShowLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckCompositeShowOrDeleteFlag(compositeShow.compositeID, verbose); err != nil {
				return err
			}

			if compositeShow.client == nil {
				compositeShow.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			compositeShow.teamID = teamID

			return compositeShow.run()
		},
	}

	f := cmd.Flags()

	f.StringVarP(&compositeShow.compositeID, content.CmdFlagCompositeIDLong, "", "", content.CmdFlagCompositeIDDescription)

	return cmd
}

func (t *compositeShowCmd) run() error {
	composite, err := t.client.ShowCompositeByID(t.teamID, t.compositeID)
	if err != nil {
		return err
	}

	util.PrintResult(
		t.out,
		composite,
		[]string{"ID", "Name", "TeamID"},
		map[string]string{
			"ID":     "Composite ID",
			"Name":   "Composite Name",
			"TeamID": "Team ID",
		},
		json,
		verbose)

	return nil
}
