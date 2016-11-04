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

	"github.com/cloudcoreo/cli/cmd/content"
	"github.com/cloudcoreo/cli/cmd/util"
	"github.com/cloudcoreo/cli/client"
	"github.com/spf13/cobra"
)

// CloudShowCmd represents the based command for cloud subcommands
var CloudShowCmd = &cobra.Command{
	Use: content.CMD_CLOUD_SHOW_USE,
	Short: content.CMD_CLOUD_SHOW_SHORT,
	Long: content.CMD_CLOUD_SHOW_LONG,
	PreRun:func(cmd *cobra.Command, args []string) {
		if err := util.CheckCloudShowOrDeleteFlag(cloudID); err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}
		SetupCoreoCredentials()
		SetupCoreoDefaultTeam()

	},
	Run:func(cmd *cobra.Command, args []string) {
		c, err := client.MakeClient(key, secret, content.ENDPOINT_ADDRESS)

		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}

		t, err := c.GetCloudAccountByID(context.Background(), teamID, cloudID)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}

		if format == "json" {
			util.PrettyPrintJson(t)
		} else {
			table := util.NewTable()
			table.SetHeader([] string{"ID", "Name", "TeamID"})
			table.UseObj(t)
			fmt.Println(table.Render())
		}
	},
}

func init() {
	CloudCmd.AddCommand(CloudShowCmd)

	CloudShowCmd.Flags().StringVarP(&cloudID, content.CMD_FLAG_ID_LONG, content.CMD_FLAG_ID_SHORT, "",content.CMD_FLAG_CLOUDID_DESCRIPTION )
}
