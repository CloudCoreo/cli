package main

import (
	"fmt"
	"io"

	"github.com/CloudCoreo/cli/pkg/coreo"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/pkg/command"

	"github.com/spf13/cobra"
)

type eventSetupCmd struct {
	client         command.Interface
	out            io.Writer
	awsProfile     string
	awsProfilePath string
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
			return eventSetup.run()
		},
	}
	f := cmd.Flags()
	f.StringVarP(&eventSetup.awsProfile, content.CmdFlagAwsProfile, "", "", content.CmdFlagAwsProfileDescription)
	f.StringVarP(&eventSetup.awsProfilePath, content.CmdFlagAwsProfilePath, "", "", content.CmdFlagAwsProfilePathDescription)

	return cmd
}

func (t *eventSetupCmd) run() error {
	input := &command.SetupEventStreamInput{
		AwsProfile:     t.awsProfile,
		AwsProfilePath: t.awsProfilePath,
	}
	err := t.client.SetupEventStream(input)
	if err != nil {
		return err
	}

	fmt.Fprintln(t.out, "Setup event stream successfully!")
	return nil
}
