package main

import (
	"io"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type cloudShowCmd struct {
	out     io.Writer
	client  coreo.Interface
	teamID  string
	cloudID string
}

func newCloudShowCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	cloudShow := &cloudShowCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdShowUse,
		Short: content.CmdCloudShowShort,
		Long:  content.CmdCloudShowLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckCloudShowOrDeleteFlag(cloudShow.cloudID, verbose); err != nil {
				return err
			}

			if cloudShow.client == nil {
				cloudShow.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			cloudShow.teamID = teamID

			return cloudShow.run()
		},
	}

	f := cmd.Flags()

	f.StringVarP(&cloudShow.cloudID, content.CmdFlagCloudIDLong, "", "", content.CmdFlagCloudIDDescripton)

	return cmd
}

func (t *cloudShowCmd) run() error {
	cloud, err := t.client.ShowCloudAccountByID(t.teamID, t.cloudID)
	if err != nil {
		return err
	}

	util.PrintResult(
		t.out,
		cloud,
		[]string{"ID", "Name", "TeamID"},
		map[string]string{
			"ID":     "Cloud Account ID",
			"Name":   "Cloud Account Name",
			"TeamID": "Team ID",
		},
		json,
		verbose)

	return nil
}
