package main

import (
	"fmt"
	"io"

	"github.com/CloudCoreo/cli/cmd/util"

	"github.com/CloudCoreo/cli/pkg/aws"

	"github.com/CloudCoreo/cli/pkg/coreo"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/pkg/command"

	"github.com/spf13/cobra"
)

type eventSetupCmd struct {
	client         command.Interface
	cloud          command.CloudProvider
	out            io.Writer
	awsProfile     string
	awsProfilePath string
	cloudID        string
	teamID         string
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
			if eventSetup.cloud == nil {
				newServiceInput := &aws.NewServiceInput{
					AwsProfile:     eventSetup.awsProfile,
					AwsProfilePath: eventSetup.awsProfilePath,
				}
				eventSetup.cloud = aws.NewService(newServiceInput)
			}

			eventSetup.teamID = teamID

			return eventSetup.run()
		},
	}
	f := cmd.Flags()
	f.StringVarP(&eventSetup.awsProfile, content.CmdFlagAwsProfile, "", "", content.CmdFlagAwsProfileDescription)
	f.StringVarP(&eventSetup.awsProfilePath, content.CmdFlagAwsProfilePath, "", "", content.CmdFlagAwsProfilePathDescription)
	f.StringVarP(&eventSetup.cloudID, content.CmdFlagCloudIDLong, "", "", content.CmdFlagCloudIDDescription)
	return cmd
}

func (t *eventSetupCmd) run() error {

	config, err := t.client.GetEventStreamConfig(t.teamID, t.cloudID)
	if err != nil {
		return err
	}

	err = t.cloud.SetupEventStream(config)
	if err != nil {
		return err
	}
	fmt.Fprintln(t.out, "Setup event stream successfully!")
	return nil
}
