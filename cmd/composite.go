package main

import (
	"io"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

func newCompositeCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:               content.CmdCompositeUse,
		Short:             content.CmdCompositeShort,
		Long:              content.CmdCompositeLong,
		PersistentPreRunE: setupCoreoConfig,
	}

	cmd.AddCommand(newCompositeListCmd(nil, out))
	cmd.AddCommand(newCompositeShowCmd(nil, out))
	cmd.AddCommand(newCompositeCreateCmd(nil, out))
	cmd.AddCommand(newCompositeInitCmd(out))
	cmd.AddCommand(newCompositeLayerCmd(out))
	cmd.AddCommand(newCompositeExtendsCmd(out))
	cmd.AddCommand(newCompositeGendocCmd(out))

	return cmd
}

type compositeListCmd struct {
	out    io.Writer
	client coreo.Interface
	teamID string
}

func newCompositeListCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	compositeList := &compositeListCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdListUse,
		Short: content.CmdCompositeListShort,
		Long:  content.CmdCompositeListLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if compositeList.client == nil {
				compositeList.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			compositeList.teamID = teamID

			return compositeList.run()
		},
	}

	return cmd
}

func (t *compositeListCmd) run() error {
	composites, err := t.client.ListComposites(t.teamID)
	if err != nil {
		return err
	}

	b := make([]interface{}, len(composites))
	for i := range composites {
		b[i] = composites[i]
	}

	util.PrintResult(
		t.out,
		b,
		[]string{"ID", "Name", "TeamID"},
		map[string]string{
			"ID":     "Composite ID",
			"Name":   "Composite Name",
			"TeamID": "Team ID",
		},
		json,
		verbose)

	return nil
}
