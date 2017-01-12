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

type planPanelCmd struct {
	out         io.Writer
	client      coreo.Interface
	teamID      string
	compositeID string
	planID      string
}

func newPlanPanelCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	planPanel := &planPanelCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   "panel",
		Short: content.CmdPlanShowShort,
		Long:  content.CmdPlanShowLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckCompositeIDAndPlandIDFlag(planPanel.compositeID, planPanel.planID, verbose); err != nil {
				return err
			}

			if planPanel.client == nil {
				planPanel.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			planPanel.teamID = teamID

			return planPanel.run()
		},
	}

	f := cmd.Flags()

	f.StringVarP(&planPanel.compositeID, content.CmdFlagCompositeIDLong, "", "", content.CmdFlagCompositeIDDescription)
	f.StringVarP(&planPanel.planID, content.CmdFlagPlanIDLong, "", "", content.CmdFlagPlanIDDescription)

	return cmd
}

func (t *planPanelCmd) run() error {
	plan, err := t.client.GetPlanPanel(t.teamID, t.compositeID, t.planID)
	if err != nil {
		return err
	}

	util.PrettyPrintJSON(plan)

	return nil
}
