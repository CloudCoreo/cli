package main

import (
	"io"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type cloudCreateCmd struct {
	out            io.Writer
	client         coreo.Interface
	teamID         string
	cloudID        string
	resourceKey    string
	resourceSecret string
	resourceName   string
}

func newCloudCreateCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	cloudCreate := &cloudCreateCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdCreateUse,
		Short: content.CmdCloudAddShort,
		Long:  content.CmdCloudAddLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckCloudAddFlags(cloudCreate.resourceName, cloudCreate.resourceKey, cloudCreate.resourceSecret); err != nil {
				return err
			}

			if cloudCreate.client == nil {
				cloudCreate.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			cloudCreate.teamID = teamID

			return cloudCreate.run()
		},
	}

	f := cmd.Flags()

	f.StringVarP(&cloudCreate.cloudID, content.CmdFlagCloudIDLong, "", "", content.CmdFlagCloudIDDescripton)
	f.StringVarP(&cloudCreate.resourceKey, content.CmdFlagKeyLong, content.CmdFlagKeyShort, "", content.CmdFlagKeyDescription)
	f.StringVarP(&cloudCreate.resourceSecret, content.CmdFlagSecretLong, content.CmdFlagSecretShort, "", content.CmdFlagSecretDescription)
	f.StringVarP(&cloudCreate.resourceName, content.CmdFlagNameLong, content.CmdFlagNameShort, "", content.CmdFlagNameDescription)

	return cmd
}

func (t *cloudCreateCmd) run() error {
	cloud, err := t.client.CreateCloudAccount(t.teamID, t.resourceKey, t.resourceSecret, t.resourceName)
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
