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
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type compositeDeleteCmd struct {
	out         io.Writer
	client      coreo.Interface
	teamID      string
	compositeID string
}

func newCompositeDeleteCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	compositeDelete := &compositeDeleteCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdDeleteUse,
		Short: content.CmdCompositeDeleteShort,
		Long:  content.CmdCompositeDeleteLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckCompositeShowOrDeleteFlag(compositeDelete.compositeID, verbose); err != nil {
				return err
			}

			if compositeDelete.client == nil {
				compositeDelete.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			compositeDelete.teamID = teamID

			return compositeDelete.run()
		},
	}

	f := cmd.Flags()

	f.StringVarP(&compositeDelete.compositeID, content.CmdFlagCompositeIDLong, "", "", content.CmdFlagCompositeIDDescription)

	return cmd
}

func (t *compositeDeleteCmd) run() error {
	err := t.client.DeleteCompositeByID(t.teamID, t.compositeID)
	if err != nil {
		return err
	}

	fmt.Fprintln(t.out, content.InfoCompositeDeletedSuccess)

	return nil
}
