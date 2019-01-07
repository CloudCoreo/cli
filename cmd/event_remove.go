package main

import (
	"fmt"
	"io"

	"github.com/CloudCoreo/cli/pkg/aws"
	"github.com/CloudCoreo/cli/pkg/coreo"

	"github.com/CloudCoreo/cli/cmd/util"

	"github.com/CloudCoreo/cli/cmd/content"

	"github.com/CloudCoreo/cli/pkg/command"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type eventRemoveCmd struct {
	client         command.Interface
	cloud          command.CloudProvider
	out            io.Writer
	awsProfile     string
	awsProfilePath string
	cloudID        string
	teamID         string
}

func newEventRemoveCmd(client command.Interface, provider command.CloudProvider, out io.Writer) *cobra.Command {
	eventRemove := &eventRemoveCmd{
		client: client,
		out:    out,
		cloud:  provider,
	}

	cmd := &cobra.Command{
		Use:     content.CmdEventRemoveUse,
		Short:   content.CmdEventRemoveShort,
		Long:    content.CmdEventRemoveLong,
		Example: content.CmdEventRemoveExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Check for --cloud-id
			if err := util.CheckCloudShowOrDeleteFlag(eventRemove.cloudID, verbose); err != nil {
				return err
			}
			if eventRemove.client == nil {
				eventRemove.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}
			if eventRemove.cloud == nil {
				newServiceInput := &aws.NewServiceInput{
					AwsProfile:     eventRemove.awsProfile,
					AwsProfilePath: eventRemove.awsProfilePath,
				}
				eventRemove.cloud = aws.NewService(newServiceInput)
			}

			eventRemove.teamID = teamID
			return eventRemove.run()
		},
	}
	f := cmd.Flags()
	f.StringVarP(&eventRemove.awsProfile, content.CmdFlagAwsProfile, "", "", content.CmdFlagAwsProfileDescription)
	f.StringVarP(&eventRemove.awsProfilePath, content.CmdFlagAwsProfilePath, "", "", content.CmdFlagAwsProfilePathDescription)
	f.StringVarP(&eventRemove.cloudID, content.CmdFlagCloudIDLong, "", "", content.CmdFlagCloudIDDescription)
	return cmd
}

func (t *eventRemoveCmd) run() error {
	config, err := t.client.GetEventRemoveConfig(t.teamID, t.cloudID)
	if err != nil {
		return err
	}

	if len(config.Regions) == 0 {
		return errors.New("No regions returned")
	}

	err = t.cloud.RemoveEventStream(config)
	if err != nil {
		return err
	}

	fmt.Fprintln(t.out, "Removed event stream successfully!")
	return nil
}
