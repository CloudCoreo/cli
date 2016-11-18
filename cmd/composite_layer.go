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
	"path"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/spf13/cobra"
)

var cmdCompositeLayer = &cobra.Command{
	Use:   content.CmdLayerUse,
	Short: content.CmdCompositeLayerShort,
	Long:  content.CmdCompositeLayerLong,
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := util.CheckLayersFlags(name, gitRepoURL); err != nil {
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

		err := util.CreateFolder("stack-"+name, directory)

		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}

		directory = path.Join(directory, "stack-"+name)

		err = util.CreateGitSubmodule(directory, gitRepoURL)

		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}
		fmt.Println(content.CmdCompositeLayerSuccess)

		// generate override and service files
		genContent(directory)

		if serverDir {
			genServerContent(directory)
		}
	},
}

func init() {
	CompositeCmd.AddCommand(cmdCompositeLayer)

	cmdCompositeLayer.Flags().StringVarP(&directory, content.CmdFlagDirectoryLong, content.CmdFlagDirectoryShort, "", content.CmdFlagDirectoryDescription)
	cmdCompositeLayer.Flags().StringVarP(&gitRepoURL, content.CmdFlagGitRepoLong, content.CmdFlagGitRepoShort, "", content.CmdFlagGitRepoDescription)
	cmdCompositeLayer.Flags().StringVarP(&name, content.CmdFlagNameLong, content.CmdFlagNameShort, "", content.CmdFlagNameDescription)
	cmdCompositeLayer.Flags().BoolVarP(&serverDir, content.CmdFlagServerLong, content.CmdFlagServerShort, false, content.CmdFlagServerDescription)
}
