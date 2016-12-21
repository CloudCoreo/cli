package main

import (
	"io"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type planEnableCmd struct {
	out         io.Writer
	client      coreo.Interface
	teamID      string
	compositeID string
	planID      string
}

func newPlanEnableCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	planEnable := &planEnableCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdEnableUse,
		Short: content.CmdPlanEnableShort,
		Long:  content.CmdPlanEnableLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckCompositeIDAndPlandIDFlag(planEnable.compositeID, planEnable.planID, verbose); err != nil {

			}

			if planEnable.client == nil {
				planEnable.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			planEnable.teamID = teamID

			return planEnable.run()
		},
	}

	f := cmd.Flags()

	f.StringVarP(&planEnable.compositeID, content.CmdFlagCompositeIDLong, "", "", content.CmdFlagCompositeIDDescription)
	f.StringVarP(&planEnable.planID, content.CmdFlagPlanIDLong, "", "", content.CmdFlagPlanIDDescription)

	return cmd
}

func (t *planEnableCmd) run() error {
	plan, err := t.client.EnablePlanByID(t.teamID, t.compositeID, t.planID)
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
