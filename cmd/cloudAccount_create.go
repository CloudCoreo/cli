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

type cloudCreateCmd struct {
	out            io.Writer
	client         coreo.Interface
	teamID         string
	cloudID        string
	resourceKey    string
	resourceSecret string
	resourceName   string
}

func newCloudCreateCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	cloudCreate := &cloudCreateCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdAddUse,
		Short: content.CmdCloudAddShort,
		Long:  content.CmdCloudAddLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckCloudAddFlags(cloudCreate.resourceName, cloudCreate.resourceKey, cloudCreate.resourceSecret); err != nil {
				return err
			}

			if cloudCreate.client == nil {
				cloudCreate.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			cloudCreate.teamID = teamID

			return cloudCreate.run()
		},
	}

	f := cmd.Flags()

	f.StringVarP(&cloudCreate.cloudID, content.CmdFlagCloudIDLong, "", "", content.CmdFlagCloudIDDescription)
	f.StringVarP(&cloudCreate.resourceKey, content.CmdFlagKeyLong, content.CmdFlagKeyShort, "", content.CmdFlagKeyDescription)
	f.StringVarP(&cloudCreate.resourceSecret, content.CmdFlagSecretLong, content.CmdFlagSecretShort, "", content.CmdFlagSecretDescription)
	f.StringVarP(&cloudCreate.resourceName, content.CmdFlagNameLong, content.CmdFlagNameShort, "", content.CmdFlagNameDescription)

	return cmd
}

func (t *cloudCreateCmd) run() error {
	cloud, err := t.client.CreateCloudAccount(t.teamID, t.resourceKey, t.resourceSecret, t.resourceName)
	if err != nil {
		return err
	}

	util.PrintResult(
		t.out,
		cloud,
		[]string{"ID", "Name", "TeamID"},
		map[string]string{
			"ID":     "Cloud Account ID",
			"Name":   "Cloud Account Name",
			"TeamID": "Team ID",
		},
		jsonFormat,
		verbose)

	return nil
}
