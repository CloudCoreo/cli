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

// TokenShowCmd represents the based command for token subcommands
var TokenShowCmd = &cobra.Command{
	Use: content.CMD_TOKEN_SHOW_USE,
	Short: content.CMD_TOKEN_SHOW_SHORT,
	Long: content.CMD_TOKEN_SHOW_LONG,
	PreRun:func(cmd *cobra.Command, args []string) {
		if err := util.CheckTokenShowOrDeleteFlag(tokenID); err != nil {
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

		t, err := c.GetTokenByID(context.Background(), tokenID)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}

		if format == "json" {
			util.PrettyPrintJson(t)
		} else {
			table := util.NewTable()
			table.SetHeader([] string{"ID", "Name", "Description"})
			table.UseObj(t)
			fmt.Println(table.Render())
		}
	},
}

func init() {
	TokenCmd.AddCommand(TokenShowCmd)

	TokenShowCmd.Flags().StringVarP(&tokenID, content.CMD_FLAG_ID_LONG, content.CMD_FLAG_ID_SHORT, "",content.CMD_FLAG_TOKENID_DESCRIPTION )
}
