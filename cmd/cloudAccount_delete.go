package main

import (
	"io"

	"fmt"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type cloudDeleteCmd struct {
	out     io.Writer
	client  coreo.Interface
	teamID  string
	cloudID string
}

func newCloudDeleteCmd(client coreo.Interface, out io.Writer) *cobra.Command {
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
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			cloudDelete.teamID = teamID

			return cloudDelete.run()
		},
	}

	f := cmd.Flags()

	f.StringVarP(&cloudDelete.cloudID, content.CmdFlagCloudIDLong, "", "", content.CmdFlagCloudIDDescripton)

	return cmd
}

func (t *cloudDeleteCmd) run() error {
	err := t.client.DeleteCloudAccountByID(t.teamID, t.cloudID)
	if err != nil {
		return err
	}

	fmt.Fprintln(t.out, content.InfoCloudAccountDeleted)

	return nil
}
