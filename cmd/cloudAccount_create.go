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

	"github.com/CloudCoreo/cli/pkg/aws"

	"github.com/CloudCoreo/cli/client"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/command"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type cloudCreateCmd struct {
	out            io.Writer
	client         command.Interface
	cloud          command.CloudProvider
	teamID         string
	resourceName   string
	roleName       string
	externalID     string
	roleArn        string
	awsProfile     string
	awsProfilePath string
	policy         string
}

func newCloudCreateCmd(client command.Interface, out io.Writer) *cobra.Command {
	cloudCreate := &cloudCreateCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:     content.CmdAddUse,
		Short:   content.CmdCloudAddShort,
		Long:    content.CmdCloudAddLong,
		Example: content.CmdCloudAddExample,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckCloudAddFlags(cloudCreate.externalID, cloudCreate.roleArn, cloudCreate.roleName); err != nil {
				return err
			}

			if cloudCreate.client == nil {
				cloudCreate.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			if cloudCreate.cloud == nil {
				cloudCreate.cloud = aws.NewService(cloudCreate.awsProfile, cloudCreate.awsProfilePath, "")
			}

			cloudCreate.teamID = teamID

			return cloudCreate.run()
		},
	}

	f := cmd.Flags()

	f.StringVarP(&cloudCreate.resourceName, content.CmdFlagNameLong, content.CmdFlagNameShort, "", content.CmdFlagNameDescription)
	f.StringVarP(&cloudCreate.roleName, content.CmdFlagRoleName, "", "", content.CmdFlagRoleNameDescription)
	f.StringVarP(&cloudCreate.roleArn, content.CmdFlagRoleArn, "", "", content.CmdFlagRoleArnDescription)
	f.StringVarP(&cloudCreate.externalID, content.CmdFlagRoleExternalID, "", "", content.CmdFlagRoleExternalIDDescription)
	f.StringVarP(&cloudCreate.awsProfile, content.CmdFlagAwsProfile, "", "", content.CmdFlagAwsProfileDescription)
	f.StringVarP(&cloudCreate.awsProfilePath, content.CmdFlagAwsProfilePath, "", "", content.CmdFlagAwsProfilePathDescription)
	f.StringVarP(&cloudCreate.policy, content.CmdFlagAwsPolicy, "", content.CmdFlagAwsPolicyDefault, content.CmdFlagAwsPolicyDescription)
	return cmd
}

func (t *cloudCreateCmd) run() error {
	input := &client.CreateCloudAccountInput{
		TeamID:     t.teamID,
		CloudName:  t.resourceName,
		RoleName:   t.roleName,
		ExternalID: t.externalID,
		RoleArn:    t.roleArn,
		Policy:     t.policy,
	}
	if t.roleName != "" {
		info, err := t.client.GetRoleCreationInfo(input)
		if err != nil {
			return err
		}
		arn, externalId, err := t.cloud.CreateNewRole(info)

		if err != nil {
			return err
		}

		input.RoleArn = arn
		input.ExternalID = externalId
	}

	cloud, err := t.client.CreateCloudAccount(input)
	if err != nil {
		t.cloud.DeleteRole(t.roleName, t.policy)
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
