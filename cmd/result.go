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
	"io"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type resultCmd struct {
	client  coreo.Interface
	teamID  string
	cloudID string
}

//Return result command. If teamID and cloudID are not specified,
//will return all violations under the user account
func newResultCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:               content.CmdResultUse,
		Short:             content.CmdResultShort,
		Long:              content.CmdResultLong,
		PersistentPreRunE: setupCoreoCredentials,
	}

	cmd.AddCommand(newResultRuleCmd(nil, out))
	cmd.AddCommand(newResultObjectCmd(nil, out))

	return cmd
}
