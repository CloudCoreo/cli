package main

import (
	"github.com/CloudCoreo/cli/pkg/coreo"
	"io"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/pkg/command"
	"github.com/spf13/cobra"
)

type cloudScanCmd struct {
	out            io.Writer
	awsProfile     string
	awsProfilePath string
	roleArn        string
	client         command.Interface
	cloud          command.CloudProvider
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


			return cloudScan.run()
		},
	}

	return cmd
}

func (t *cloudScanCmd) run() error {
	return nil
}
