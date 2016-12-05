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

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/CloudCoreo/cli/client"
	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/spf13/cobra"
)

// PlanShowCmd represents the based command for plan subcommands
var PlanShowCmd = &cobra.Command{
	Use:   content.CmdShowUse,
	Short: content.CmdPlanShowShort,
	Long:  content.CmdPlanShowLong,
	PreRun: func(cmd *cobra.Command, args []string) {
		util.CheckArgsCount(args)

		SetupCoreoCredentials()
		SetupCoreoDefaultTeam()
		if err := util.CheckCompositeIDAndPlandIDFlag(compositeID, planID, verbose); err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.MakeClient(key, secret, apiEndpont)
		if err != nil {
			util.PrintError(err, json)
			os.Exit(-1)
		}

		t, err := c.GetPlanByID(context.Background(), teamID, compositeID, planID)
		if err != nil {
			util.PrintError(err, json)
			os.Exit(-1)
		}

		util.PrintResult(
			t,
			[]string{"ID", "Name", "Enabled", "Branch", "RefreshInterval"},
			map[string]string {
				"ID": "Plan ID",
				"Name": "Plan Name",
				"Enabled" : "Active",
				"Branch" : "Git Branch",
				"RefreshInterval": "Interval",
			},
			json,
			verbose)	},
}

func init() {
	PlanCmd.AddCommand(PlanShowCmd)

	PlanShowCmd.Flags().StringVarP(&planID, content.CmdFlagPlanIDLong, "", "", content.CmdFlagPlanIDDescription)
	PlanShowCmd.Flags().StringVarP(&compositeID, content.CmdFlagCompositeIDLong, "", "", content.CmdFlagCompositeIDDescription)
}
