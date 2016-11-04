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



// GitKeyAddCmd represents the based command for gitkey subcommands
var GitKeyAddCmd = &cobra.Command{
	Use: content.CMD_GITKEY_ADD_USE,
	Short: content.CMD_GITKEY_ADD_SHORT,
	Long: content.CMD_GITKEY_ADD_LONG,
	PreRun:func(cmd *cobra.Command, args []string) {
		if err := util.CheckGitKeyAddFlags(resourceName, resourceSecret); err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}
		SetupCoreoCredentials()
	},
	Run:func(cmd *cobra.Command, args []string) {
		c, err := client.MakeClient(key, secret, content.ENDPOINT_ADDRESS)
		t, err := c.CreateGitKey(context.Background(), teamID, resourceSecret, resourceName)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}

		fmt.Printf("%#v", t)
	},
}

func init() {
	GitKeyCmd.AddCommand(GitKeyAddCmd)

	GitKeyAddCmd.Flags().StringVarP(&resourceSecret, content.CMD_FLAG_SECRET_LONG, content.CMD_FLAG_SECRET_SHORT, "",content.CMD_FLAG_SECRET_DESCRIPTION )
	GitKeyAddCmd.Flags().StringVarP(&resourceName, content.CMD_FLAG_NAME_LONG, content.CMD_FLAG_NAME_SHORT, "",content.CMD_FLAG_NAME_DESCRIPTION )
}
