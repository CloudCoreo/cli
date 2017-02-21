package main

import (
	"io"

	"fmt"
	"os"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type planInitCmd struct {
	out         io.Writer
	client      coreo.Interface
	teamID      string
	compositeID string
	name        string
	interval    int
	region      string
	cloudID     string
	revision    string
	branch      string
	directory   string
}

func newPlanInitCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	planInit := &planInitCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   "add [flags]",
		Short: content.CmdPlanInitShort,
		Long:  content.CmdPlanInitLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckPlanInitRequiredFlags(planInit.compositeID, planInit.cloudID, planInit.name); err != nil {
				return err
			}

			if planInit.client == nil {
				planInit.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			planInit.teamID = teamID

			return planInit.run()
		},
	}

	f := cmd.Flags()

	f.StringVarP(&planInit.compositeID, content.CmdFlagCompositeIDLong, "", "", content.CmdFlagCompositeIDDescription)
	f.StringVarP(&planInit.name, content.CmdFlagNameLong, "", "", content.CmdFlagNameDescription)
	f.StringVarP(&planInit.region, content.CmdFlagCloudRegionLong, "", "us-east-1", content.CmdFlagCloudRegionDescription)
	f.StringVarP(&planInit.cloudID, content.CmdFlagCloudIDLong, "", "", content.CmdFlagCloudIDDescripton)
	f.StringVarP(&planInit.revision, content.CmdFlagGitRevisionLong, "", "HEAD", content.CmdFlagGitRevisionDescription)
	f.StringVarP(&planInit.branch, content.CmdFlagBranchLong, "", "master", content.CmdFlagBranchDescription)
	f.StringVarP(&planInit.directory, content.CmdFlagDirectoryLong, content.CmdFlagDirectoryShort, "", content.CmdFlagDirectoryDescription)
	f.IntVarP(&planInit.interval, content.CmdFlagIntervalLong, "", 1440, content.CmdFlagIntervalDescription)

	return cmd
}

func (t *planInitCmd) run() error {
	planConfig, err := t.client.InitPlan(t.branch, t.name, t.region, t.teamID, t.cloudID, t.compositeID, t.revision, t.interval)
	if err != nil {
		return err
	}

	if t.directory == "" {
		t.directory, _ = os.Getwd()
	}
	util.CreateFile(fmt.Sprintf("%s.json", t.name), t.directory, util.PrettyJSON(planConfig), true)

	if verbose {
		fmt.Printf(content.InfoPlanJSONFileCreated, t.name, t.directory)
	}
	return nil
}
