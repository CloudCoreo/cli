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

	"fmt"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type planDeleteCmd struct {
	out         io.Writer
	client      coreo.Interface
	teamID      string
	compositeID string
	planID      string
}

func newPlanDeleteCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	planDelete := &planDeleteCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdDeleteUse,
		Short: content.CmdPlanDeleteShort,
		Long:  content.CmdPlanDeleteLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckCompositeIDAndPlandIDFlag(planDelete.compositeID, planDelete.planID, verbose); err != nil {
				return err
			}

			if planDelete.client == nil {
				planDelete.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			planDelete.teamID = teamID

			return planDelete.run()
		},
	}

	f := cmd.Flags()

	f.StringVarP(&planDelete.compositeID, content.CmdFlagCompositeIDLong, "", "", content.CmdFlagCompositeIDDescription)
	f.StringVarP(&planDelete.planID, content.CmdFlagPlanIDLong, "", "", content.CmdFlagPlanIDDescription)

	return cmd
}

func (t *planDeleteCmd) run() error {
	err := t.client.DeletePlanByID(t.teamID, t.compositeID, t.planID)
	if err != nil {
		return err
	}

	if verbose {
		fmt.Println(content.InfoPlanDeleted)
	}
	return nil
}
