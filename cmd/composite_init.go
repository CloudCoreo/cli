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

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/spf13/cobra"
	"github.com/CloudCoreo/cli/cmd/util"
	"fmt"
	"os"
	"path"
)

var cmdCompositeInit = &cobra.Command{
	Use: content.CMD_COMPOSITE_INIT_USE,
	Short: content.CMD_COMPOSITE_INIT_SHORT,
	Long: content.CMD_COMPOSITE_INIT_LONG,
	Run: func(cmd *cobra.Command, args []string) {

		if directory == "" {
			directory, _ = os.Getwd()
		}

		genContent(directory)

		if serverDir {
			genServerContent(directory)
		}
	},
}

func genContent (directory string) {
	if directory == "" {
		directory, _ = os.Getwd()
	}

	// config.yml file
	fmt.Println()
	util.CreateFile(content.DEFAULT_FILES_CONFIG_YAML_FILE, directory,content.DEFAULT_FILES_CONFIG_YAML, false)

	// override folder
	util.CreateFolder(content.DEFAULT_FILES_OVERRIDES_FOLDER, directory)

	overrideTree := fmt.Sprintf(content.DEFAULT_FILES_OVERRIDES_README_TREE, content.DEFAULT_FILES_README_CODE_TICKS, content.DEFAULT_FILES_README_CODE_TICKS)

	overrideReadmeContent := fmt.Sprintf("%s%s%s", content.DEFAULT_FILES_OVERRIDES_README_HEADER, overrideTree, content.DEFAULT_FILES_OVERRIDES_README_FOOTER)

	err := util.CreateFile(content.DEFAULT_FILES_README_FILE, path.Join(directory, content.DEFAULT_FILES_OVERRIDES_FOLDER), overrideReadmeContent, false)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)

	}

	// services folder
	util.CreateFolder(content.DEFAULT_FILES_SERVICES_FOLDER, directory)

	err = util.CreateFile(content.DEFAULT_FILES_CONFIG_RB_FILE, path.Join(directory, content.DEFAULT_FILES_SERVICES_FOLDER), content.DEFAULT_FILES_CONFIG_RB, false)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}

	servicesReadMeCode := fmt.Sprintf(content.DEFAULT_FILES_SERVICES_README_CODE, content.DEFAULT_FILES_README_CODE_TICKS, content.DEFAULT_FILES_README_CODE_TICKS)

	servicesReadMeContent := fmt.Sprintf("%s%s", content.DEFAULT_FILES_SERVICES_README_HEADER, servicesReadMeCode)

	err = util.CreateFile(content.DEFAULT_FILES_README_FILE, path.Join(directory + content.DEFAULT_FILES_SERVICES_FOLDER), servicesReadMeContent, false)

	if err != nil {
		fmt.Println(err.Error())
	}

	if err == nil {
		fmt.Println(content.CMD_COMPOSITE_INIT_SUCCESS)
	}
}

func genServerContent(directory string) {
	//operational scripts dir
	util.CreateFolder(content.DEFAULT_FILES_OPERATIONAL_SCRIPTS_FOLDER, directory)

	// generate operational readme file
	err := util.CreateFile(content.DEFAULT_FILES_README_FILE, path.Join(directory, content.DEFAULT_FILES_OPERATIONAL_SCRIPTS_FOLDER), content.DEFAULT_FILES_OPERATIONAL_README_CONTENT, false)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}

	//boot scripts dir
	util.CreateFolder(content.DEFAULT_FILES_BOOT_SCRIPTS_FOLDER, directory)

	//README.md
	err = util.CreateFile(content.DEFAULT_FILES_README_FILE, path.Join(directory, content.DEFAULT_FILES_BOOT_SCRIPTS_FOLDER), content.DEFAULT_FILES_BOOT_README_CONTENT, false)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}

	//order.yaml
	err = util.CreateFile(content.DEFAULT_FILES_ORDER_YAML_FILE, path.Join(directory, content.DEFAULT_FILES_BOOT_SCRIPTS_FOLDER), content.DEFAULT_FILES_BOOT_ORDER_YAML_CONTENT, false)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}


	//shutdown scripts dir
	util.CreateFolder(content.DEFAULT_FILES_SHUTDOWN_SCRIPTS_FOLDER, directory)


	//README.md
	err = util.CreateFile(content.DEFAULT_FILES_README_FILE, path.Join(directory, content.DEFAULT_FILES_SHUTDOWN_SCRIPTS_FOLDER), content.DEFAULT_FILES_SHUTDOWN_README_CONTENT, false)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}

	//order.yaml
	err = util.CreateFile(content.DEFAULT_FILES_ORDER_YAML_FILE, path.Join(directory, content.DEFAULT_FILES_SHUTDOWN_SCRIPTS_FOLDER), content.DEFAULT_FILES_SHUTDOWN_ORDER_YAML_CONTENT, false)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}
}

func init() {
	CompositeCmd.AddCommand(cmdCompositeInit)

	cmdCompositeInit.Flags().StringVarP(&directory, content.CMD_FLAG_DIRECTORY_LONG, content.CMD_FLAG_DIRECTORY_SHORT, "", content.CMD_FLAG_DIRECTORY_DESCRIPTION)
	cmdCompositeInit.Flags().BoolVarP(&serverDir, content.CMD_FLAG_SERVER_LONG, content.CMD_FLAG_SERVER_SHORT, false, content.CMD_FLAG_SERVER_DESCRIPTION )
}

