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

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type devTimeResultsCmd struct {
	out       io.Writer
	client    coreo.Interface
	devTimeID string
	teamID    string
}

func newDevTimeResultsCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	devTimeResults := &devTimeResultsCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdResultsUse,
		Short: content.CmdDevTimeResultsShort,
		Long:  content.CmdDevTimeResultsLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckDevTimeIDAddFlags(devTimeResults.devTimeID); err != nil {
				return err
			}

			if devTimeResults.client == nil {
				devTimeResults.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			devTimeResults.teamID = teamID

			return devTimeResults.run()
		},
	}
	f := cmd.Flags()

	f.StringVarP(&devTimeResults.devTimeID, content.CmdFlagDevTimeIDLong, "", "", content.CmdFlagDevTimeIDDescription)

	return cmd
}

func (t *devTimeResultsCmd) run() error {
	results, err := t.client.GetDevTimeResults(t.teamID, t.devTimeID)
	if err != nil {
		return err
	}

	util.PrettyPrintJSON(results)

	return nil
}
