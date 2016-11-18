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
	Use:   content.CmdInitUse,
	Short: content.CmdPlanInitShort,
	Long:  content.CmdPlanInitLong,
	PreRun: func(cmd *cobra.Command, args []string) {
		SetupCoreoCredentials()
		SetupCoreoDefaultTeam()
		if err := util.CheckCompositeShowOrDeleteFlag(compositeID, verbose); err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.MakeClient(key, secret, content.EndpointAddress)
		if err != nil {
			util.PrintError(err, json)
			os.Exit(-1)
		}

		t, err := c.InitPlan(context.Background(), branch, name, interval, region, teamID, cloudID, compositeID, revision)
		if err != nil {
			util.PrintError(err, json)
			os.Exit(-1)
		}

		util.PrintResult(t, []string{"ID", "Name", "Enabled", "Branch", "RefreshInterval"}, json)
	},
}

func init() {
	PlanCmd.AddCommand(PlanInitCmd)

	PlanInitCmd.Flags().StringVarP(&compositeID, content.CmdFlagCompositeIDLong, "", "", content.CmdFlagCompositeIDDescription)
	PlanInitCmd.Flags().StringVarP(&cloudID, content.CmdFlagCloudIDLong, "", "", content.CmdFlagCloudIDDescripton)
	PlanInitCmd.Flags().StringVarP(&name, content.CmdFlagNameLong, content.CmdFlagNameShort, "", content.CmdFlagNameDescription)
	PlanInitCmd.Flags().StringVarP(&revision, content.CmdFlagGitCommitIDLong, "", "HEAD", content.CmdFlagGitCommitIDDescription)
	PlanInitCmd.Flags().StringVarP(&region, content.CmdFlagCloudRegionLong, "", "us-east-1", content.CmdFlagCloudRegionDescription)
	PlanInitCmd.Flags().StringVarP(&interval, content.CmdFlagIntervalLong, "", "1", content.CmdFlagIntervalDescription)
	PlanInitCmd.Flags().StringVarP(&branch, content.CmdFlagBranchLong, "", "master", content.CmdFlagBranchDescription)
}
