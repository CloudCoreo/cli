package main

import (
	"fmt"
	"io"

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

func newEventSetupCmd(client command.Interface, out io.Writer) *cobra.Command {
	eventSetup := &eventSetupCmd{
		client: client,
		out:    out,
	}

	cmd := &cobra.Command{
		Use:     content.CmdEventSetupUse,
		Short:   content.CmdEventSetupShort,
		Long:    content.CmdEventSetupLong,
		Example: content.CmdEventSetupExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			if eventSetup.client == nil {
				eventSetup.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}
			if eventSetup.cloud == nil {
				eventSetup.cloud = aws.NewService(eventSetup.awsProfile, eventSetup.awsProfilePath, "")
			}
			return eventSetup.run()
		},
	}
	eventSetup.teamID = teamID
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

	/*
		input := &command.SetupEventStreamInput{
			AwsProfile:     t.awsProfile,
			AwsProfilePath: t.awsProfilePath,
			Config:         config,
		}
	*/
	err = t.cloud.SetupEventStream(config)
	if err != nil {
		return err
	}
	fmt.Fprintln(t.out, "Setup event stream successfully!")
	return nil
}
