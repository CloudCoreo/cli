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

// CompositeShowCmd represents the based command for composite subcommands
var CompositeShowCmd = &cobra.Command{
	Use:   content.CmdShowUse,
	Short: content.CmdCompositeShowShort,
	Long:  content.CmdCompositeShowLong,
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

		t, err := c.GetCompositeByID(context.Background(), teamID, compositeID)
		if err != nil {
			util.PrintError(err, json)
			os.Exit(-1)
		}

		util.PrintResult(t, []string{"ID", "Name", "CreatedAt", "GitKeyID", "GitURL"}, json)
	},
}

func init() {
	CompositeCmd.AddCommand(CompositeShowCmd)

	CompositeShowCmd.Flags().StringVarP(&compositeID, content.CmdFlagCompositeIDLong, "", "", content.CmdFlagCompositeIDDescription)
}
