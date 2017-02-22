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

	"io/ioutil"
	"path"
	"path/filepath"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/spf13/cobra"
)

type compositeExtendsCmd struct {
	out        io.Writer
	directory  string
	gitRepoURL string
	serverDir  bool
}

func newCompositeExtendsCmd(out io.Writer) *cobra.Command {
	compositeExtends := &compositeExtendsCmd{
		out: out,
	}

	cmd := &cobra.Command{
		Use:   content.CmdExtendsUse,
		Short: content.CmdCompositeExtendsShort,
		Long:  content.CmdCompositeExtendsLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckExtendFlags(compositeExtends.gitRepoURL); err != nil {
				return err
			}

			return compositeExtends.run()
		},
	}

	f := cmd.Flags()

	f.StringVarP(&compositeExtends.directory, content.CmdFlagDirectoryLong, content.CmdFlagDirectoryShort, "", content.CmdFlagDirectoryDescription)
	f.BoolVarP(&compositeExtends.serverDir, content.CmdFlagServerLong, content.CmdFlagServerShort, false, content.CmdFlagServerDescription)
	f.StringVarP(&compositeExtends.gitRepoURL, content.CmdFlagGitRepoLong, content.CmdFlagGitRepoShort, "", content.CmdFlagGitRepoDescription)

	return cmd
}

func tryReplaceRootConfigYamlFile(directory string) error {
	fp := filepath.Join(directory, content.DefaultFilesConfigYAMLName)
	fi, err := os.Stat(fp)

	if err != nil {
		return err
	}

	if fi.Size() == 0 {
		fileContent, err := ioutil.ReadFile(path.Join(directory, "extends", content.DefaultFilesConfigYAMLName))
		if err != nil {
			return err
		}

		f, err := os.OpenFile(path.Join(directory, content.DefaultFilesConfigYAMLName), os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			return err
		}

		defer f.Close()

		if _, err = f.WriteString(string(fileContent)); err != nil {
			return err
		}

	}

	return nil
}

func (t *compositeExtendsCmd) run() error {

	if err := util.CheckGitInstall(); err != nil {
		return err
	}

	if t.directory == "" {
		t.directory, _ = os.Getwd()
	}

	err := util.CreateGitSubmodule(t.directory, t.gitRepoURL)

	if err != nil {
		return err
	}

	fmt.Fprintln(t.out, content.CmdCompositeExtendsSuccess)

	// generate override and service files
	genContent(t.directory)

	// replace root config.yaml from extends folder if empty
	tryReplaceRootConfigYamlFile(t.directory)

	if t.serverDir {
		genServerContent(t.directory)
	}

	return nil
}
