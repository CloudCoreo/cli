package main

import (
	"io"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type planDeleteCmd struct {
	out         io.Writer
	client      coreo.Interface
	teamID      string
	compositeID string
	planID      string
}

func newPlanDeleteCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	planDelete := &planDeleteCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdDeleteUse,
		Short: content.CmdPlanDeleteShort,
		Long:  content.CmdPlanDeleteLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckCompositeIDAndPlandIDFlag(planDelete.compositeID, planDelete.planID, verbose); err != nil {

			}

			if planDelete.client == nil {
				planDelete.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			planDelete.teamID = teamID

			return planDelete.run()
		},
	}

	f := cmd.Flags()

	f.StringVarP(&planDelete.compositeID, content.CmdFlagCompositeIDLong, "", "", content.CmdFlagCompositeIDDescription)
	f.StringVarP(&planDelete.planID, content.CmdFlagPlanIDLong, "", "", content.CmdFlagPlanIDDescription)

	return cmd
}

func (t *planDeleteCmd) run() error {
	err := t.client.DeletePlanByID(t.teamID, t.compositeID, t.planID)
	if err != nil {
		return err
	}

	return nil
}
