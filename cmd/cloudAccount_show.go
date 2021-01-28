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

	"github.com/CloudCoreo/cli/pkg/command"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type cloudShowCmd struct {
	out           io.Writer
	client        command.Interface
	accountNumber string
	provider      string
}

func newCloudShowCmd(client command.Interface, out io.Writer) *cobra.Command {
	cloudShow := &cloudShowCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdShowUse,
		Short: content.CmdCloudShowShort,
		Long:  content.CmdCloudShowLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckCloudShowOrDeleteFlag(cloudShow.accountNumber, verbose); err != nil {
				return err
			}

			if cloudShow.client == nil {
				cloudShow.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.RefreshToken(key))
			}

			return cloudShow.run()
		},
	}

	f := cmd.Flags()

	f.StringVarP(&cloudShow.accountNumber, content.CmdFlagAccountIDLong, "", "", content.CmdFlagAccountIDDescription)
	f.StringVarP(&cloudShow.provider, content.CmdFlagProvider, "", "AWS", content.CmdFlagProviderDescription)
	return cmd
}

func (t *cloudShowCmd) run() error {
	cloud, err := t.client.ShowCloudAccountByID(t.accountNumber, t.provider)
	if err != nil {
		return err
	}

	util.PrintResult(
		t.out,
		cloud,
		[]string{"Name", "AccountID", "Provider"},
		map[string]string{
			"Name":      "Cloud Account Name",
			"AccountID": "Cloud Account ID",
			"Provider":  "Cloud Provider",
		},
		jsonFormat,
		verbose)

	return nil
}
