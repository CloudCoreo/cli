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
	"strings"

	"github.com/CloudCoreo/cli/pkg/aws"

	"github.com/CloudCoreo/cli/pkg/command"

	"fmt"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type cloudDeleteCmd struct {
	out            io.Writer
	client         command.Interface
	cloud          command.CloudProvider
	cloudID        string
	deleteRole     bool
	awsProfile     string
	awsProfilePath string
}

func newCloudDeleteCmd(client command.Interface, out io.Writer) *cobra.Command {
	cloudDelete := &cloudDeleteCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdDeleteUse,
		Short: content.CmdCloudDeleteShort,
		Long:  content.CmdCloudDeleteLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckCloudShowOrDeleteFlag(cloudDelete.cloudID, verbose); err != nil {
				return err
			}

			if cloudDelete.client == nil {
				cloudDelete.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.RefreshToken(key))
			}

			if cloudDelete.deleteRole && (cloudDelete.cloud == nil) {
				newServiceInput := &aws.NewServiceInput{
					AwsProfile:     cloudDelete.awsProfile,
					AwsProfilePath: cloudDelete.awsProfilePath,
				}
				cloudDelete.cloud = aws.NewService(newServiceInput)
			}

			return cloudDelete.run()
		},
	}

	f := cmd.Flags()

	f.StringVarP(&cloudDelete.cloudID, content.CmdFlagCloudIDLong, "", "", content.CmdFlagCloudIDDescription)
	f.BoolVarP(&cloudDelete.deleteRole, content.CmdFlagDeleteRole, "", false, content.CmdFLagDeleteRoleDescription)
	f.StringVarP(&cloudDelete.awsProfile, content.CmdFlagAwsProfile, "", "", content.CmdFlagAwsProfileDescription)
	f.StringVarP(&cloudDelete.awsProfilePath, content.CmdFlagAwsProfilePath, "", "", content.CmdFlagAwsProfilePathDescription)

	return cmd
}

func (t *cloudDeleteCmd) run() error {
	var roleName string
	if t.deleteRole {
		cloud, err := t.client.ShowCloudAccountByID(t.cloudID)
		if err != nil {
			return err
		}

		roleNames := strings.Split(cloud.Arn, "/")
		roleName = roleNames[len(roleNames)-1]

		t.cloud.DeleteRole(roleName)
	}

	err := t.client.DeleteCloudAccountByID(t.cloudID)
	if err != nil {
		return err
	}

	fmt.Fprintln(t.out, content.InfoCloudAccountDeleted)

	return nil
}
