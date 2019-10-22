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
	"github.com/CloudCoreo/cli/pkg/command"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type teamShowCmd struct {
	out    io.Writer
	client command.Interface
}

func newTeamShowCmd(client command.Interface, out io.Writer) *cobra.Command {
	teamShow := &teamShowCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdShowUse,
		Short: content.CmdTeamShowShort,
		Long:  content.CmdTeamShowLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if teamShow.client == nil {
				teamShow.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.RefreshToken(key))
			}

			_, err := fmt.Fprint(out, "Teams are deprecated` \n")
			return err
		},
	}

	return cmd
}
