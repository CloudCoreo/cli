package main

import (
	"io"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

func newCloudAccountCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:               content.CmdCloudUse,
		Short:             content.CmdCloudShort,
		Long:              content.CmdCloudLong,
		PersistentPreRunE: setupCoreoConfig,
	}

	cmd.AddCommand(newCloudListCmd(nil, out))
	cmd.AddCommand(newCloudShowCmd(nil, out))
	cmd.AddCommand(newCloudDeleteCmd(nil, out))
	cmd.AddCommand(newCloudCreateCmd(nil, out))

	return cmd
}

type cloudListCmd struct {
	out    io.Writer
	client coreo.Interface
	teamID string
}

func newCloudListCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	cloudList := &cloudListCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdListUse,
		Short: content.CmdCloudListShort,
		Long:  content.CmdCloudListLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if cloudList.client == nil {
				cloudList.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			cloudList.teamID = teamID

			return cloudList.run()
		},
	}

	return cmd
}

func (t *cloudListCmd) run() error {
	clouds, err := t.client.ListCloudAccounts(t.teamID)
	if err != nil {
		return err
	}

	b := make([]interface{}, len(clouds))
	for i := range clouds {
		b[i] = clouds[i]
	}

	util.PrintResult(
		t.out,
		b,
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
