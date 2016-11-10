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

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/client"
	"github.com/spf13/cobra"
)

// CompositeCreateCmd represents the based command for Composite subcommands
var CompositeCreateCmd = &cobra.Command{
	Use: content.CMD_COMPOSITE_CREATE_USE,
	Short: content.CMD_COMPOSITE_CREATE_SHORT,
	Long: content.CMD_COMPOSITE_CREATE_LONG,
	PreRun:func(cmd *cobra.Command, args []string) {
		if err := util.CheckCompositeCreateFlags(name, gitRepoUrl); err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}
		SetupCoreoCredentials()
		SetupCoreoDefaultTeam()

	},
	Run:func(cmd *cobra.Command, args []string) {
		c, err := client.MakeClient(key, secret, content.ENDPOINT_ADDRESS)
		t, err := c.CreateComposite(context.Background(), gitRepoUrl, name, teamID)
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
	CompositeCmd.AddCommand(CompositeCreateCmd)

	CompositeCreateCmd.Flags().StringVarP(&name, content.CMD_FLAG_NAME_LONG, content.CMD_FLAG_NAME_SHORT, "",content.CMD_FLAG_NAME_DESCRIPTION )
	CompositeCreateCmd.Flags().StringVarP(&gitRepoUrl, content.CMD_FLAG_GIT_REPO_LONG, content.CMD_FLAG_GIT_REPO_SHORT, "",content.CMD_FLAG_GIT_REPO_DESCRIPTION )
}
