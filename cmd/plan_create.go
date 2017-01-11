package main

import (
	"io"

	"io/ioutil"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type planCreateCmd struct {
	out                io.Writer
	client             coreo.Interface
	planConfigJSONFile string
}

func newPlanCreateCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	planCreate := &planCreateCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdPlanFinalizeUse,
		Short: content.CmdPlanCreateShort,
		Long:  content.CmdPlanCreateLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckPlanCreateJSONFileFlag(planCreate.planConfigJSONFile); err != nil {
				return err
			}

			if planCreate.client == nil {
				planCreate.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			return planCreate.run()
		},
	}

	f := cmd.Flags()

	f.StringVarP(&planCreate.planConfigJSONFile, content.CmdFlagJSONFileLong, content.CmdFlagJSONFileShort, "", content.CmdFlagJSONFileDescription)

	return cmd
}

func (t *planCreateCmd) run() error {

	planConfigJSON, err := ioutil.ReadFile(t.planConfigJSONFile)

	plan, err := t.client.CreatePlan(planConfigJSON)
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
		jsonFormat,
		verbose)

	return nil
}
