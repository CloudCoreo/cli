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

	"github.com/CloudCoreo/cli/pkg/command"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

func newCloudAccountCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:               content.CmdCloudUse,
		Short:             content.CmdCloudShort,
		Long:              content.CmdCloudLong,
		PersistentPreRunE: setupCoreoConfig,
	}

	cmd.AddCommand(newCloudListCmd(nil, out))
	cmd.AddCommand(newCloudShowCmd(nil, out))
	cmd.AddCommand(newCloudDeleteCmd(nil, out))
	cmd.AddCommand(newCloudCreateCmd(nil, out))
	cmd.AddCommand(newCloudScanCmd(nil, out))
	cmd.AddCommand(newCloudUpdateCmd(nil, out))
	cmd.AddCommand(newCloudTestCmd(nil, out))

	return cmd
}

type cloudListCmd struct {
	out    io.Writer
	client command.Interface
	teamID string
}

func newCloudListCmd(client command.Interface, out io.Writer) *cobra.Command {
	cloudList := &cloudListCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdListUse,
		Short: content.CmdCloudListShort,
		Long:  content.CmdCloudListLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if cloudList.client == nil {
				cloudList.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			cloudList.teamID = teamID

			return cloudList.run()
		},
	}

	return cmd
}

func (t *cloudListCmd) run() error {
	clouds, err := t.client.ListCloudAccounts(t.teamID)
	if err != nil {
		return err
	}

	b := make([]interface{}, len(clouds))
	for i := range clouds {
		b[i] = clouds[i]
	}

	util.PrintResult(
		t.out,
		b,
		[]string{"ID", "Name", "TeamID", "AccountID", "IsDraft"},
		map[string]string{
			"ID":        "Cloud Account ID",
			"Name":      "Cloud Account Name",
			"TeamID":    "Team ID",
			"AccountID": "AWS account ID",
			"IsDraft":   "IsDraft",
		},
		jsonFormat,
		verbose)

	return nil
}

type cloudTestCmd struct {
	out     io.Writer
	client  command.Interface
	teamID  string
	cloudID string
}

func newCloudTestCmd(client command.Interface, out io.Writer) *cobra.Command {
	cloudTest := &cloudTestCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdTestUse,
		Short: content.CmdCloudTestShort,
		Long:  content.CmdCloudTestLong,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := util.CheckCloudShowOrDeleteFlag(cloudTest.cloudID, verbose); err != nil {
				return err
			}
			if cloudTest.client == nil {
				cloudTest.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			cloudTest.teamID = teamID
			return cloudTest.run()
		},
	}

	f := cmd.Flags()

	f.StringVarP(&cloudTest.cloudID, content.CmdFlagCloudIDLong, "", "", content.CmdFlagCloudIDDescription)

	return cmd
}

func (t *cloudTestCmd) run() error {
	res, err := t.client.ReValidateRole(t.teamID, t.cloudID)
	if err != nil {
		return err
	}
	if jsonFormat {
		util.PrettyPrintJSON(res)
	} else {
		fmt.Fprintln(t.out, res.Message)
	}
	return nil
}
