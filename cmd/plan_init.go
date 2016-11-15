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

var revision, region, interval, branch string

// PlanInitCmd represents the based command for plan subcommands
var PlanInitCmd = &cobra.Command{
	Use:   content.CMD_INIT_USE,
	Short: content.CMD_PLAN_INIT_SHORT,
	Long:  content.CMD_PLAN_INIT_LONG,
	PreRun: func(cmd *cobra.Command, args []string) {
		util.CheckCompositeShowOrDeleteFlag(compositeID)
		SetupCoreoCredentials()
		SetupCoreoDefaultTeam()

	},
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.MakeClient(key, secret, content.ENDPOINT_ADDRESS)
		t, err := c.InitPlan(context.Background(), branch, name, interval, region, teamID, cloudID, compositeID, revision)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}

		if format == "json" {
			util.PrettyPrintJSON(t)
		} else {
			table := util.NewTable()
			table.UseObj(t)
			fmt.Println(table.Render())
		}
	},
}

func init() {
	PlanCmd.AddCommand(PlanInitCmd)

	PlanInitCmd.Flags().StringVarP(&compositeID, content.CMD_FLAG_ID_LONG, content.CMD_FLAG_ID_SHORT, "", content.CMD_FLAG_COMPOSITE_DESCRIPTION)
	PlanInitCmd.Flags().StringVarP(&cloudID, "cloud-id", "", "", content.CMD_FLAG_COMPOSITE_DESCRIPTION)

	PlanInitCmd.Flags().StringVarP(&name, content.CMD_FLAG_NAME_LONG, content.CMD_FLAG_NAME_SHORT, "", content.CMD_FLAG_NAME_DESCRIPTION)
	PlanInitCmd.Flags().StringVarP(&revision, "gitcommit-id", "", "HEAD", content.CMD_FLAG_COMPOSITE_DESCRIPTION)
	PlanInitCmd.Flags().StringVarP(&region, "region", "", "us-east-1", content.CMD_FLAG_COMPOSITE_DESCRIPTION)
	PlanInitCmd.Flags().StringVarP(&interval, "interval", "", "1", content.CMD_FLAG_COMPOSITE_DESCRIPTION)
	PlanInitCmd.Flags().StringVarP(&branch, "branch", "", "master", content.CMD_FLAG_COMPOSITE_DESCRIPTION)
}
