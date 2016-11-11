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
	Use:   content.CMD_SHOW_USE,
	Short: content.CMD_PLAN_SHOW_SHORT,
	Long:  content.CMD_PLAN_SHOW_LONG,
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := util.CheckCloudShowOrDeleteFlag(cloudID); err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}
		SetupCoreoCredentials()
		SetupCoreoDefaultTeam()

	},
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.MakeClient(key, secret, content.ENDPOINT_ADDRESS)

		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}

		t, err := c.GetPlanByID(context.Background(), teamID, cloudID, planID)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}

		if format == "json" {
			util.PrettyPrintJSON(t)
		} else {
			table := util.NewTable()
			table.SetHeader([]string{"ID", "Name", "TeamID"})
			table.UseObj(t)
			fmt.Println(table.Render())
		}
	},
}

func init() {
	PlanCmd.AddCommand(PlanShowCmd)

	PlanShowCmd.Flags().StringVarP(&planID, content.CMD_FLAG_PLAN_ID_LONG, content.CMD_FLAG_PLAN_ID_SHORT, "", content.CMD_FLAG_PLANID_DESCRIPTION)
	PlanShowCmd.Flags().StringVarP(&compositeID, content.CMD_FLAG_ID_LONG, content.CMD_FLAG_ID_SHORT, "", content.CMD_FLAG_COMPOSITE_DESCRIPTION)
}
