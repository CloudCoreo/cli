// Copyright © 2018 Zechen Jiang <zechen@cloudcoreo.com>
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

type resultObjectCmd struct {
	client   command.Interface
	out      io.Writer
}

func newResultObjectCmd(client command.Interface, out io.Writer) *cobra.Command {
	resultObject := &resultObjectCmd{
		client: client,
		out:    out,
	}
	cmd := &cobra.Command{
		Use:     content.CmdResultObjectUse,
		Short:   content.CmdResultObjectShort,
		Long:    content.CmdResultObjectLong,
		Run: func(cmd *cobra.Command, args []string) {
			if resultObject.client == nil {
				resultObject.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.RefreshToken(key))
			}
			fmt.Fprint(resultObject.out, "Findings results are deprecated, please follow the link to swagger API doc `https://api.securestate.vmware.com` \n")
		},
	}
	return cmd
}