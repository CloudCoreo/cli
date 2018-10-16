// Copyright © 2018 Zechen Jiang <zechen@cloudcoreo.com>
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
	"github.com/CloudCoreo/cli/pkg/coreo"
	"io"
	"github.com/spf13/cobra"
	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"fmt"
)

type resultRuleCmd struct{
	client coreo.Interface
	teamID string
	cloudID string
	out io.Writer
	level string
}

func newResultRuleCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	resultRule := &resultRuleCmd{
		client: client,
		out: out,
	}
	cmd := &cobra.Command{
		Use: content.CmdResultRuleUse,
		Short: content.CmdResultRuleShort,
		Long: content.CmdResultRuleLong,
		RunE: func(cmd *cobra.Command, args []string) error {
			if resultRule.client == nil {
				resultRule.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			resultRule.teamID = teamID

			return resultRule.run()
		},
	}
	f := cmd.Flags()
	f.StringVar(&resultRule.cloudID, content.CmdFlagCloudIDLong, content.None, content.CmdFlagCloudIDDescription)
	f.StringVar(&resultRule.level, content.CmdFlagLevelLong, content.None, content.CmdFlagLevelDescription)
	return cmd
}

func (t *resultRuleCmd) run() error {
	res, err := t.client.ShowResultRule(t.teamID, t.cloudID, t.level)
	if err != nil {
		return err
	}
	b := make([]interface{}, len(res))
	for i := range res {
		b[i] = res[i]
	}
	fmt.Fprintln(t.out, util.PrettyJSON(b))

	return nil
}