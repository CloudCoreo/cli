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

type devTimeJobsCmd struct {
	out       io.Writer
	client    coreo.Interface
	devTimeID string
	teamID    string
}

func newDevTimeJobsCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	devTimeJobs := &devTimeJobsCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdJobsUse,
		Short: content.CmdDevTimeJobsShort,
		Long:  content.CmdDevTimeJobsLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckDevTimeIDAddFlags(devTimeJobs.devTimeID); err != nil {
				return err
			}

			if devTimeJobs.client == nil {
				devTimeJobs.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			devTimeJobs.teamID = teamID

			return devTimeJobs.run()
		},
	}
	f := cmd.Flags()

	f.StringVarP(&devTimeJobs.devTimeID, content.CmdFlagDevTimeIDLong, "", "", content.CmdFlagDevTimeIDDescription)

	return cmd
}

func (t *devTimeJobsCmd) run() error {

	// gitKey, err := t.client.CreateDevTime(t.teamID, t.context, t.task)
	// if err != nil {
	// 	return err
	// }

	// util.PrintResult(
	// 	t.out,
	// 	gitKey,
	// 	[]string{"DevTimeID", "DevTimeURL", "Context", "Task"},
	// 	map[string]string{
	// 		"DevTimeID":  "DevTime ID",
	// 		"DevTimeURL": "DevTime Url",
	// 		"Context":    "DevTime Context",
	// 		"Task":       "DevTime Task",
	// 	},
	// 	jsonFormat,
	// 	verbose)

	return nil
}
