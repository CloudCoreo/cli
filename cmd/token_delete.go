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

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/command"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type tokenDeleteCmd struct {
	out     io.Writer
	client  command.Interface
	tokenID string
}

func newTokenDeleteCmd(client command.Interface, out io.Writer) *cobra.Command {
	tokenDelete := &tokenDeleteCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdDeleteUse,
		Short: content.CmdTokenDeleteShort,
		Long:  content.CmdTokenDeleteLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckTokenShowOrDeleteFlag(tokenDelete.tokenID, verbose); err != nil {
				return err
			}

			if tokenDelete.client == nil {
				tokenDelete.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.RefreshToken(key))
			}
			_, err := fmt.Fprint(out, "Tokens are deprecated, only csp token is required` \n")
			return err
		},
	}

	f := cmd.Flags()
	f.StringVarP(&tokenDelete.tokenID, content.CmdFlagTokenIDLong, "", "", content.CmdFlagTokenIDDescription)

	return cmd
}
