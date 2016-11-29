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

// TokenShowCmd represents the based command for token subcommands
var TokenShowCmd = &cobra.Command{
	Use:   content.CmdShowUse,
	Short: content.CmdTokenShowShort,
	Long:  content.CmdTokenShowLong,
	PreRun: func(cmd *cobra.Command, args []string) {
		util.CheckArgsCount(args)

		SetupCoreoCredentials()
		SetupCoreoDefaultTeam()
		if err := util.CheckTokenShowOrDeleteFlag(tokenID, verbose); err != nil {
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

		t, err := c.GetTokenByID(context.Background(), tokenID)
		if err != nil {
			util.PrintError(err, json)
			os.Exit(-1)
		}

		util.PrintResult(t, []string{"ID", "Name", "Description"}, json)
	},
}

func init() {
	TokenCmd.AddCommand(TokenShowCmd)

	TokenShowCmd.Flags().StringVarP(&tokenID, content.CmdFlagTokenIDLong, "", "", content.CmdFlagTokenIDDescription)
}
