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
	Use:   content.CMD_COMPOSITE_LAYER_USE,
	Short: content.CMD_COMPOSITE_LAYER_SHORT,
	Long:  content.CMD_COMPOSITE_LAYER_LONG,
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
		fmt.Println(content.CMD_COMPOSITE_LAYER_SUCCESS)

		// generate override and service files
		genContent(directory)

		if serverDir {
			genServerContent(directory)
		}
	},
}

func init() {
	CompositeCmd.AddCommand(cmdCompositeLayer)

	cmdCompositeLayer.Flags().StringVarP(&directory, content.CMD_FLAG_DIRECTORY_LONG, content.CMD_FLAG_DIRECTORY_SHORT, "", content.CMD_FLAG_DIRECTORY_DESCRIPTION)
	cmdCompositeLayer.Flags().StringVarP(&gitRepoURL, content.CMD_FLAG_GIT_REPO_LONG, content.CMD_FLAG_GIT_REPO_SHORT, "", content.CMD_FLAG_GIT_REPO_DESCRIPTION)
	cmdCompositeLayer.Flags().StringVarP(&name, content.CMD_FLAG_NAME_LONG, content.CMD_FLAG_NAME_SHORT, "", content.CMD_FLAG_NAME_DESCRIPTION)
	cmdCompositeLayer.Flags().BoolVarP(&serverDir, content.CMD_FLAG_SERVER_LONG, content.CMD_FLAG_SERVER_SHORT, false, content.CMD_FLAG_SERVER_DESCRIPTION)
}
