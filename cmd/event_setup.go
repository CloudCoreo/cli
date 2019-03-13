package main

import (
	"errors"
	"fmt"
	"io"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/aws"
	"github.com/CloudCoreo/cli/pkg/azure"
	"github.com/CloudCoreo/cli/pkg/command"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type eventSetupCmd struct {
	client              command.Interface
	cloud               command.CloudProvider
	out                 io.Writer
	awsProfile          string
	awsProfilePath      string
	cloudID             string
	teamID              string
	ignoreMissingTrails bool
	awsRoleArn          string
	awsExternalID       string
	authFile            string
	region              string
}

func newEventSetupCmd(client command.Interface, provider command.CloudProvider, out io.Writer) *cobra.Command {
	eventSetup := &eventSetupCmd{
		client: client,
		out:    out,
		cloud:  provider,
	}

	cmd := &cobra.Command{
		Use:     content.CmdEventSetupUse,
		Short:   content.CmdEventSetupShort,
		Long:    content.CmdEventSetupLong,
		Example: content.CmdEventSetupExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Check for --cloud-id
			if err := util.CheckCloudShowOrDeleteFlag(eventSetup.cloudID, verbose); err != nil {
				return err
			}
			if eventSetup.client == nil {
				eventSetup.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			eventSetup.teamID = teamID

			return eventSetup.run()
		},
	}
	f := cmd.Flags()
	f.StringVarP(&eventSetup.awsProfile, content.CmdFlagAwsProfile, "", "", content.CmdFlagAwsProfileDescription)
	f.StringVarP(&eventSetup.awsProfilePath, content.CmdFlagAwsProfilePath, "", "", content.CmdFlagAwsProfilePathDescription)
	f.StringVarP(&eventSetup.cloudID, content.CmdFlagCloudIDLong, "", "", content.CmdFlagCloudIDDescription)
	f.BoolVarP(&eventSetup.ignoreMissingTrails, content.CmdFlagIgnoreMissingTrails, "", false, content.CmdFlagIgnoreMissingTrailsDescription)
	f.StringVarP(&eventSetup.awsRoleArn, content.CmdFlagAwsRoleArn, "", "", content.CmdFlagAwsRoleArnDescription)
	f.StringVarP(&eventSetup.awsExternalID, content.CmdFlagAwsExternalID, "", "", content.CmdFlagAwsExternalIDDescription)
	f.StringVarP(&eventSetup.authFile, content.CmdEventAuthFile, "", "", content.CmdEventAuthFileDescription)
	f.StringVarP(&eventSetup.region, content.CmdEventRegion, "", "eastus", content.CmdEventRegionDescription)
	return cmd
}

func (t *eventSetupCmd) run() error {

	config, err := t.client.GetEventStreamConfig(t.teamID, t.cloudID)
	if err != nil {
		return err
	}

	if t.cloud == nil {
		if config.Provider == "AWS" {
			newServiceInput := &aws.NewServiceInput{
				AwsProfile:          t.awsProfile,
				AwsProfilePath:      t.awsProfilePath,
				IgnoreMissingTrails: t.ignoreMissingTrails,
				RoleArn:             t.awsRoleArn,
				ExternalID:          t.awsExternalID,
			}
			t.cloud = aws.NewService(newServiceInput)
		} else if config.Provider == "Azure" {
			newServiceInput := &azure.NewServiceInput{
				AuthFile: t.authFile,
				Region:   t.region,
			}
			t.cloud = azure.NewService(newServiceInput)
		} else {
			return errors.New("unsupported provider type " + config.Provider + " ")
		}

	}

	if config.Provider == "AWS" && len(config.Regions) == 0 {
		return errors.New("No regions returned")
	}
	err = t.cloud.SetupEventStream(config)
	if err != nil {
		return err
	}
	fmt.Fprintln(t.out, "Setup event stream successfully!")
	return nil
}
