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
	"os"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/spf13/cobra"
)

var cmdCompositeExtends = &cobra.Command{
	Use:   content.CmdExtendsUse,
	Short: content.CmdCompositeExtendsShort,
	Long:  content.CmdCompositeExtendsLong,
	PreRun: func(cmd *cobra.Command, args []string) {
		util.CheckArgsCount(args)
		if err := util.CheckExtendFlags(gitRepoURL); err != nil {
			fmt.Println("A composite name is required: -n")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

		if err := util.CheckGitInstall(); err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}

		if directory == "" {
			directory, _ = os.Getwd()
		}

		err := util.CreateGitSubmodule(directory, gitRepoURL)

		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}

		fmt.Println(content.CmdCompositeExtendsSuccess)

		// generate override and service files
		genContent(directory)

		if serverDir {
			genServerContent(directory)
		}
	},
}

func init() {
	CompositeCmd.AddCommand(cmdCompositeExtends)

	cmdCompositeExtends.Flags().StringVarP(&directory, content.CmdFlagDirectoryLong, content.CmdFlagDirectoryShort, "", content.CmdFlagDirectoryDescription)
	cmdCompositeExtends.Flags().StringVarP(&gitRepoURL, content.CmdFlagGitRepoLong, content.CmdFlagGitRepoShort, "", content.CmdFlagGitRepoDescription)
	cmdCompositeExtends.Flags().BoolVarP(&serverDir, content.CmdFlagServerLong, content.CmdFlagServerShort, false, content.CmdFlagServerDescription)
}
