package main

import (
	"io"

	"github.com/CloudCoreo/cli/client"

	"github.com/CloudCoreo/cli/pkg/aws"
	"github.com/CloudCoreo/cli/pkg/coreo"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/pkg/command"
	"github.com/spf13/cobra"
)

type cloudScanCmd struct {
	out             io.Writer
	awsProfile      string
	awsProfilePath  string
	roleArn         string
	client          command.Interface
	cloud           command.CloudProvider
	policy          string
	roleSessionName string
	duration        int64
	teamID          string
}

func newCloudScanCmd(client command.Interface, out io.Writer) *cobra.Command {
	cloudScan := &cloudScanCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdScanUse,
		Short: content.CmdCloudScanShort,
		Long:  content.CmdCloudScanLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if cloudScan.client == nil {
				cloudScan.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			if cloudScan.cloud == nil {
				newServiceInput := &aws.NewServiceInput{
					AwsProfile:      cloudScan.awsProfile,
					AwsProfilePath:  cloudScan.awsProfilePath,
					RoleArn:         cloudScan.roleArn,
					Policy:          cloudScan.policy,
					RoleSessionName: cloudScan.roleSessionName,
					Duration:        cloudScan.duration,
				}
				cloudScan.cloud = aws.NewService(newServiceInput)
			}
			return cloudScan.run()
		},
	}

	return cmd
}

func (t *cloudScanCmd) run() error {
	roots, err := t.cloud.GetOrgTree()
	if err != nil {
		return err
	}
	for _, root := range roots {
		err = t.InOrderTraversalTree(root)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *cloudScanCmd) InOrderTraversalTree(root *command.TreeNode) error {
	switch root.Info.Type {
	case "ORGANIZATIONAL_UNIT":
		// Do nothing for Organizational Unit
	case "ACCOUNT":
		input := &client.CreateCloudAccountInput{
			IsDraft:   true,
			TeamID:    t.teamID,
			CloudName: root.Info.Name,
			Email:     root.Info.Properties["email"],
			UserName:  root.Info.ID,
		}
		_, err := t.client.CreateCloudAccount(input)
		if err != nil {
			return err
		}
	default:
		return client.NewError("Unknown type found while creating drafts")
	}

	for _, child := range root.Children {
		err := t.InOrderTraversalTree(child)
		if err != nil {
			return err
		}
	}
	return nil
}
