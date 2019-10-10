package main

import (
	"fmt"
	"io"
	"time"

	"github.com/CloudCoreo/cli/pkg/aws"

	"github.com/CloudCoreo/cli/client"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"

	"github.com/CloudCoreo/cli/cmd/content"

	"github.com/spf13/cobra"

	"github.com/CloudCoreo/cli/pkg/command"
)

type cloudUpdateCmd struct {
	out            io.Writer
	client         command.Interface
	cloud          command.CloudProvider
	cloudID        string
	resourceName   string
	roleName       string
	externalID     string
	roleArn        string
	isDraft        bool
	userName       string
	email          string
	environment    string
	awsProfile     string
	awsProfilePath string
	policy         string
	tags           string
}

func newCloudUpdateCmd(client command.Interface, out io.Writer) *cobra.Command {
	cloudUpdate := &cloudUpdateCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdUpdateUse,
		Short: content.CmdCloudUpdateShort,
		Long:  content.CmdCloudUpdateLong,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := util.CheckCloudShowOrDeleteFlag(cloudUpdate.cloudID, verbose); err != nil {
				return err
			}

			if cloudUpdate.client == nil {
				cloudUpdate.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.RefreshToken(key))
			}

			if cloudUpdate.cloud == nil {
				newServiceInput := &aws.NewServiceInput{
					AwsProfile:     cloudUpdate.awsProfile,
					AwsProfilePath: cloudUpdate.awsProfilePath,
				}
				cloudUpdate.cloud = aws.NewService(newServiceInput)
			}

			return cloudUpdate.run()
		},
	}
	f := cmd.Flags()

	f.StringVarP(&cloudUpdate.resourceName, content.CmdFlagNameLong, content.CmdFlagNameShort, "", content.CmdFlagNameDescription)
	f.StringVarP(&cloudUpdate.roleArn, content.CmdFlagRoleArn, "", "", content.CmdFlagRoleArnDescription)
	f.StringVarP(&cloudUpdate.externalID, content.CmdFlagRoleExternalID, "", "", content.CmdFlagRoleExternalIDDescription)
	f.BoolVarP(&cloudUpdate.isDraft, content.CmdFlagIsDraft, "", false, content.CmdFlagIsDraftDescription)
	f.StringVarP(&cloudUpdate.email, content.CmdFlagEmail, "", "", content.CmdFlagEmailDescription)
	f.StringVarP(&cloudUpdate.userName, content.CmdFlagUserName, "", "", content.CmdFlagUserNameDescription)
	f.StringVarP(&cloudUpdate.environment, content.CmdFlagEnvironmentLong, content.CmdFlagEnvironmentShort, "", content.CmdFlagEnvironmentDescription)
	f.StringVarP(&cloudUpdate.cloudID, content.CmdFlagCloudIDLong, "", "", content.CmdFlagCloudIDDescription)
	f.StringVarP(&cloudUpdate.awsProfile, content.CmdFlagAwsProfile, "", "", content.CmdFlagAwsProfileDescription)
	f.StringVarP(&cloudUpdate.awsProfilePath, content.CmdFlagAwsProfilePath, "", "", content.CmdFlagAwsProfilePathDescription)
	f.StringVarP(&cloudUpdate.policy, content.CmdFlagAwsPolicy, "", content.CmdFlagAwsPolicyDefault, content.CmdFlagAwsPolicyDescription)
	f.StringVarP(&cloudUpdate.roleName, content.CmdFlagRoleName, "", "", content.CmdFlagRoleNameDescription)
	f.StringVarP(&cloudUpdate.tags, content.CmdFlagTags, "", "", content.CmdFlagTagsDescription)
	return cmd

}

func (t *cloudUpdateCmd) run() error {
	input := &client.UpdateCloudAccountInput{
		CreateCloudAccountInput: client.CreateCloudAccountInput{
			CloudName:   t.resourceName,
			RoleName:    t.roleName,
			ExternalID:  t.externalID,
			RoleArn:     t.roleArn,
			IsDraft:     t.isDraft,
			Email:       t.email,
			UserName:    t.userName,
			Environment: t.environment,
			Policy:      t.policy,
			Tags:        t.tags,
		},
		CloudID: t.cloudID,
	}

	if t.roleName != "" {
		info, err := t.client.GetRoleCreationInfo(&input.CreateCloudAccountInput)
		if err != nil {
			return err
		}
		arn, externalID, err := t.cloud.CreateNewRole(info)
		time.Sleep(10 * time.Second)
		if err != nil {
			return err
		}

		input.RoleArn = arn
		input.ExternalID = externalID
	}

	cloud, err := t.client.UpdateCloudAccount(input)
	if err != nil {
		if t.roleName != "" {
			fmt.Println("Cloud account update failed! Will delete created role.")
			t.cloud.DeleteRole(t.roleName)
		}
		return err
	}
	util.PrintResult(
		t.out,
		cloud,
		[]string{"ID", "Name"},
		map[string]string{
			"ID":   "Cloud Account ID",
			"Name": "Cloud Account Name",
		},
		jsonFormat,
		verbose)
	return nil
}
