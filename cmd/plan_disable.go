package main

import (
	"io"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type planDisableCmd struct {
	out         io.Writer
	client      coreo.Interface
	teamID      string
	compositeID string
	planID      string
}

func newPlanDisableCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	planDisable := &planDisableCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdDisableUse,
		Short: content.CmdPlanDisableShort,
		Long:  content.CmdPlanDisableLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckCompositeIDAndPlandIDFlag(planDisable.compositeID, planDisable.planID, verbose); err != nil {
				return err
			}

			if planDisable.client == nil {
				planDisable.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			planDisable.teamID = teamID

			return planDisable.run()
		},
	}

	f := cmd.Flags()

	f.StringVarP(&planDisable.compositeID, content.CmdFlagCompositeIDLong, "", "", content.CmdFlagCompositeIDDescription)
	f.StringVarP(&planDisable.planID, content.CmdFlagPlanIDLong, "", "", content.CmdFlagPlanIDDescription)

	return cmd
}

func (t *planDisableCmd) run() error {
	plan, err := t.client.DisablePlanByID(t.teamID, t.compositeID, t.planID)
	if err != nil {
		return err
	}

	util.PrintResult(
		t.out,
		plan,
		planSchema,
		planHeader,
		jsonFormat,
		verbose)

	return nil
}
