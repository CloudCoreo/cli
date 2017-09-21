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

func newDevTimeCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:               content.CmdDevTimeKeyUse,
		Short:             content.CmdDevTimeShot,
		Long:              content.CmdDevTimeLong,
		PersistentPreRunE: setupCoreoConfig,
	}

	cmd.AddCommand(newDevTimeCreateCmd(nil, out))
	cmd.AddCommand(newDevTimeStopCmd(nil, out))
	cmd.AddCommand(newDevTimeStartCmd(nil, out))
	cmd.AddCommand(newDevTimeJobsCmd(nil, out))
	cmd.AddCommand(newDevTimeResultsCmd(nil, out))

	return cmd
}

type devTimeCreateCmd struct {
	out     io.Writer
	client  coreo.Interface
	teamID  string
	task    string
	context string
}

func newDevTimeCreateCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	devTimeCreate := &devTimeCreateCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdAddUse,
		Short: content.CmdDevTimeCreateShort,
		Long:  content.CmdDevTimeCreateLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckDevTimeAddFlags(devTimeCreate.context, devTimeCreate.task); err != nil {
				return err
			}

			if devTimeCreate.client == nil {
				devTimeCreate.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			devTimeCreate.teamID = teamID

			return devTimeCreate.run()
		},
	}
	f := cmd.Flags()

	f.StringVarP(&devTimeCreate.context, content.CmdFlagDevTimeContextLong, "", "", content.CmdFlagDevTimeContextDescription)
	f.StringVarP(&devTimeCreate.task, content.CmdFlagDevTimeTaskLong, "", "", content.CmdFlagDevTimeTaskDescription)

	return cmd
}

func (t *devTimeCreateCmd) run() error {

	gitKey, err := t.client.CreateDevTime(t.teamID, t.context, t.task)
	if err != nil {
		return err
	}

	util.PrintResult(
		t.out,
		gitKey,
		[]string{"DevTimeID", "DevTimeURL", "Context", "Task"},
		map[string]string{
			"DevTimeID":  "DevTime ID",
			"DevTimeURL": "DevTime Url",
			"Context":    "DevTime Context",
			"Task":       "DevTime Task",
		},
		jsonFormat,
		verbose)

	return nil
}
