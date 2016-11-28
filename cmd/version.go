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
	"fmt"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/spf13/cobra"
)

//Version build flag
var Version = "No Version Provided"

//Githash build flag
var Githash = "Unknown githash"

//BuildID build flag
var BuildID = "Unknown buildID"

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   content.CmdVersionUse,
	Short: content.CmdVersionShort,
	Long:  content.CmdVersionLong,
	Run: func(cmd *cobra.Command, args []string) {
		util.CheckArgsCount(args)

		fmt.Printf("Version: %s\n", Version)
		fmt.Printf("Git hash: %s\n", Githash)
		fmt.Printf("BuildID: %s\n", BuildID)
	},
}
