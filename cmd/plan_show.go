package main

import (
	"io"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type planShowCmd struct {
	out         io.Writer
	client      coreo.Interface
	teamID      string
	compositeID string
	planID      string
}

func newPlanShowCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	planShow := &planShowCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdShowUse,
		Short: content.CmdPlanShowShort,
		Long:  content.CmdPlanShowLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckCompositeIDAndPlandIDFlag(planShow.compositeID, planShow.planID, verbose); err != nil {

			}

			if planShow.client == nil {
				planShow.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			planShow.teamID = teamID

			return planShow.run()
		},
	}

	f := cmd.Flags()

	f.StringVarP(&planShow.compositeID, content.CmdFlagCompositeIDLong, "", "", content.CmdFlagCompositeIDDescription)
	f.StringVarP(&planShow.planID, content.CmdFlagPlanIDLong, "", "", content.CmdFlagPlanIDDescription)

	return cmd
}

func (t *planShowCmd) run() error {
	plan, err := t.client.ShowPlanByID(t.teamID, t.compositeID, t.planID)
	if err != nil {
		return err
	}

	util.PrintResult(
		t.out,
		plan,
		[]string{"ID", "Name", "Enabled", "Branch", "RefreshInterval"},
		map[string]string{
			"ID":              "Plan ID",
			"Name":            "Plan Name",
			"Enabled":         "Active",
			"Branch":          "Git Branch",
			"RefreshInterval": "Interval",
		},
		json,
		verbose)

	return nil
}
