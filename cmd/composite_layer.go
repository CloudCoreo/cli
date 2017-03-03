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

package main

import (
	"io"

	"fmt"
	"os"
	"path"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/spf13/cobra"
)

type compositeLayerCmd struct {
	out        io.Writer
	directory  string
	name       string
	gitRepoURL string
	serverDir  bool
}

func newCompositeLayerCmd(out io.Writer) *cobra.Command {
	compositeLayer := &compositeLayerCmd{
		out: out,
	}

	cmd := &cobra.Command{
		Use:   content.CmdLayerUse,
		Short: content.CmdCompositeLayerShort,
		Long:  content.CmdCompositeLayerLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckLayersFlags(compositeLayer.name, compositeLayer.gitRepoURL); err != nil {
				return err
			}

			return compositeLayer.run()
		},
	}

	f := cmd.Flags()

	f.StringVarP(&compositeLayer.directory, content.CmdFlagDirectoryLong, content.CmdFlagDirectoryShort, "", content.CmdFlagDirectoryDescription)
	f.StringVarP(&compositeLayer.name, content.CmdFlagNameLong, content.CmdFlagNameShort, "", content.CmdFlagNameDescription)
	f.StringVarP(&compositeLayer.gitRepoURL, content.CmdFlagGitRepoLong, content.CmdFlagGitRepoShort, "", content.CmdFlagGitRepoDescription)
	f.BoolVarP(&compositeLayer.serverDir, content.CmdFlagServerLong, content.CmdFlagServerShort, false, content.CmdFlagServerDescription)

	return cmd
}

func (t *compositeLayerCmd) run() error {

	if err := util.CheckGitInstall(); err != nil {
		return err
	}

	if t.directory == "" {
		t.directory, _ = os.Getwd()
	}

	if t.name == "" {
		t.name = util.GetRepoNameFromGitURL(t.gitRepoURL)
	}

	stackName := "stack-" + t.name

	_, err := os.Stat(path.Join(t.directory, stackName))

	if !os.IsNotExist(err) {
		return fmt.Errorf(content.ErrorStackNameExist, stackName)
	}

	err = util.CreateFolder(stackName, t.directory)

	if err != nil {
		return err
	}

	t.directory = path.Join(t.directory, stackName)

	err = util.CreateGitSubmodule(t.directory, t.gitRepoURL)

	if err != nil {
		return err
	}
	fmt.Println(content.CmdCompositeLayerSuccess)

	// generate override and service files
	genContent(t.directory)

	if t.serverDir {
		genServerContent(t.directory)
	}

	return nil
}
