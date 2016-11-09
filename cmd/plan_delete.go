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

	"github.com/cloudcoreo/cli/client"
	"github.com/cloudcoreo/cli/cmd/content"
	"github.com/cloudcoreo/cli/cmd/util"
	"github.com/spf13/cobra"
)

// PlandDeleteCmd represents the based command for plan subcommands
var PlandDeleteCmd = &cobra.Command{
	Use:   content.CMD_DELETE_USE,
	Short: content.CMD_PLAN_DELETE_SHORT,
	Long:  content.CMD_PLAN_DELETE_LONG,
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := util.CheckCompositeIdAndPlandIdFlag(compositeID, planID); err != nil {
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

		err = c.DeletePlanByID(context.Background(), teamID, compositeID, planID)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}

		fmt.Println("Plan was deleted")
	},
}

func init() {
	PlanCmd.AddCommand(PlandDeleteCmd)

	PlandDeleteCmd.Flags().StringVarP(&planID, content.CMD_FLAG_ID_LONG, content.CMD_FLAG_ID_SHORT, "", content.CMD_FLAG_CLOUDID_DESCRIPTION)
	PlandDeleteCmd.Flags().StringVarP(&comositeID, content.CMD_FLAG_ID_LONG, content.CMD_FLAG_ID_SHORT, "", content.CMD_FLAG_COMPOSITE_DESCRIPTION)

}
