package main

import (
	"io"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

func newTeamCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:               content.CmdTeamUse,
		Short:             content.CmdTeamShort,
		Long:              content.CmdTeamLong,
		PersistentPreRunE: setupCoreoConfig,
	}

	cmd.AddCommand(newTeamListCmd(nil, out))
	cmd.AddCommand(newTeamShowCmd(nil, out))

	return cmd
}

type teamListCmd struct {
	out    io.Writer
	client coreo.Interface
}

func newTeamListCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	teamList := &teamListCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdListUse,
		Short: content.CmdTeamListShort,
		Long:  content.CmdTeamListLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if teamList.client == nil {
				teamList.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			return teamList.run()
		},
	}

	return cmd
}

func (t *teamListCmd) run() error {
	teams, err := t.client.ListTeams()
	if err != nil {
		return err
	}

	b := make([]interface{}, len(teams))
	for i := range teams {
		b[i] = teams[i]
	}

	util.PrintResult(t.out, b,
		[]string{"ID", "TeamName", "TeamDescription"},
		map[string]string{
			"ID":              "Team ID",
			"TeamName":        "Team Name",
			"TeamDescription": "Team Description",
		},
		json,
		verbose)

	return nil
}
