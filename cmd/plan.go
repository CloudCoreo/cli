// Copyright © 2016 Paul Allen <paul@cloudcoreo.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"io"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

func newPlanCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:               content.CmdPlanUse,
		Short:             content.CmdPlanShort,
		Long:              content.CmdPlanLong,
		PersistentPreRunE: setupCoreoConfig,
	}

	cmd.AddCommand(newPlanListCmd(nil, out))
	cmd.AddCommand(newPlanShowCmd(nil, out))
	cmd.AddCommand(newPlanDisableCmd(nil, out))
	cmd.AddCommand(newPlanEnableCmd(nil, out))
	cmd.AddCommand(newPlanDeleteCmd(nil, out))

	return cmd
}

//PlanListCmd struct
type PlanListCmd struct {
	out         io.Writer
	client      coreo.Interface
	teamID      string
	compositeID string
}

func newPlanListCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	PlanList := &PlanListCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdListUse,
		Short: content.CmdPlanListShort,
		Long:  content.CmdPlanListLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckCompositeShowOrDeleteFlag(PlanList.compositeID, verbose); err != nil {
				return err
			}

			if PlanList.client == nil {
				PlanList.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			PlanList.teamID = teamID

			return PlanList.run()
		},
	}

	f := cmd.Flags()

	f.StringVarP(&PlanList.compositeID, content.CmdFlagCompositeIDLong, "", "", content.CmdFlagCompositeIDDescription)

	return cmd
}

func (t *PlanListCmd) run() error {
	Plans, err := t.client.ListPlans(t.teamID, t.compositeID)
	if err != nil {
		return err
	}

	b := make([]interface{}, len(Plans))
	for i := range Plans {
		b[i] = Plans[i]
	}

	util.PrintResult(
		t.out,
		b,
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
