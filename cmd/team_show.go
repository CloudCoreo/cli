package main

import (
	"io"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type teamShowCmd struct {
	out    io.Writer
	client coreo.Interface
	teamID string
}

func newTeamShowCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	teamShow := &teamShowCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdShowUse,
		Short: content.CmdTeamShowShort,
		Long:  content.CmdTeamShowLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if teamShow.client == nil {
				teamShow.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			teamShow.teamID = teamID

			return teamShow.run()
		},
	}

	return cmd
}

func (t *teamShowCmd) run() error {
	team, err := t.client.ShowTeamByID(t.teamID)
	if err != nil {
		return err
	}

	util.PrintResult(t.out, team,
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
