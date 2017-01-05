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
	"fmt"
	"io"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/spf13/cobra"
)

type versionCmd struct {
	out           io.Writer
	clientVersion string
	clientGithash string
	clientBuildID string
}

var (
	version string
	githash string
	buildID string
)

func newVersionCmd(out io.Writer) *cobra.Command {
	v := &versionCmd{
		out:           out,
		clientVersion: version,
		clientGithash: githash,
		clientBuildID: buildID,
	}

	cmd := &cobra.Command{
		Use:   content.CmdVersionUse,
		Short: content.CmdVersionShort,
		Long:  content.CmdVersionLong,
		RunE: func(cmd *cobra.Command, args []string) error {
			return v.run()
		},
	}

	return cmd
}

func (v *versionCmd) run() error {
	fmt.Fprintf(v.out, "Version: %#v\n", v.clientVersion)
	fmt.Fprintf(v.out, "Git hash: %#v\n", v.clientGithash)
	fmt.Fprintf(v.out, "BuildID: %#v\n", v.clientBuildID)

	return nil
}
