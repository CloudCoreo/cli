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
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type devTimeStartCmd struct {
	out       io.Writer
	client    coreo.Interface
	devTimeID string
	teamID    string
}

func newDevTimeStartCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	devTimeStart := &devTimeStartCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdStartUse,
		Short: content.CmdDevTimeStartShort,
		Long:  content.CmdDevTimeStartLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckDevTimeIDAddFlags(devTimeStart.devTimeID); err != nil {
				return err
			}

			if devTimeStart.client == nil {
				devTimeStart.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			devTimeStart.teamID = teamID

			return devTimeStart.run()
		},
	}
	f := cmd.Flags()

	f.StringVarP(&devTimeStart.devTimeID, content.CmdFlagDevTimeIDLong, "", "", content.CmdFlagDevTimeIDDescription)

	return cmd
}

func (t *devTimeStartCmd) run() error {

	err := t.client.StartDevTime(t.teamID, t.devTimeID)
	if err != nil {
		return err
	}

	fmt.Printf(content.InfoDevTimeStarted, t.devTimeID)

	return nil
}
