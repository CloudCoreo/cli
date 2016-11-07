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

// GitKeyListCmd represents the based command for gitkey subcommands
var GitKeyListCmd = &cobra.Command{
	Use: content.CMD_GITKEY_LIST_USE,
	Short: content.CMD_GITKEY_LIST_SHORT,
	Long: content.CMD_GITKEY_LIST_LONG,
	PreRun:func(cmd *cobra.Command, args []string) {
		SetupCoreoCredentials()
		SetupCoreoDefaultTeam()

	},
	Run:func(cmd *cobra.Command, args []string) {
		c, err := client.MakeClient(key, secret, content.ENDPOINT_ADDRESS)

		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}

		t, err := c.GetGitKeys(context.Background(), teamID)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}

		b := make([]interface{}, len(t))
		for i := range t {
			b[i] = t[i]
		}

		if format == "json" {
			util.PrettyPrintJson(t)
		} else {
			table := util.NewTable()
			table.SetHeader([] string{"ID", "Name", "TeamID"})
			table.UseObj(b)
			fmt.Println(table.Render())
		}
	},
}

func init() {
	GitKeyCmd.AddCommand(GitKeyListCmd)
}