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

type planRunNowCmd struct {
	out         io.Writer
	client      coreo.Interface
	teamID      string
	compositeID string
	planID      string
	block       bool
}

func newPlanRunNowCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	planRunNow := &planRunNowCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdRunNowUse,
		Short: content.CmdPlanRunShort,
		Long:  content.CmdPlanRunLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckCompositeIDAndPlandIDFlag(planRunNow.compositeID, planRunNow.planID, verbose); err != nil {
				return err
			}

			if planRunNow.client == nil {
				planRunNow.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			planRunNow.teamID = teamID

			return planRunNow.run()
		},
	}

	f := cmd.Flags()

	f.StringVarP(&planRunNow.compositeID, content.CmdFlagCompositeIDLong, "", "", content.CmdFlagCompositeIDDescription)
	f.StringVarP(&planRunNow.planID, content.CmdFlagPlanIDLong, "", "", content.CmdFlagPlanIDDescription)
	f.BoolVarP(&planRunNow.block, content.CmdFlagBlockedLong, content.CmdFlagServerShort, false, content.CmdFlagBlockedDescription)

	return cmd
}

func (t *planRunNowCmd) run() error {
	_, err := t.client.RunNowPlanByID(t.teamID, t.compositeID, t.planID)
	if err != nil {
		return err
	}

	if t.block {
		cmd := newPlanPanelCmd(t.client, t.out)
		cmd.ParseFlags([]string{"--composite-id", t.compositeID, "--plan-id", t.planID})
		cmd.RunE(cmd, nil)
	}

	return nil
}
