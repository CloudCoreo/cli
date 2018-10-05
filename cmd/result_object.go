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
)

type resultObjectCmd struct{
	client coreo.Interface
	teamID string
	cloudID string
}

func newResultObjectCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	resultObject := &resultObjectCmd{
		client: client,
	}
	cmd := &cobra.Command{
		Use: content.CmdResultObjectUse,
		Short: content.CmdResultObjectShort,
		Long: content.CmdResultObjectLong,
		RunE: func(cmd *cobra.Command, args []string) error {
			if resultObject.client == nil {
				resultObject.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			resultObject.teamID = teamID

			return resultObject.run()
		},
	}
	f := cmd.Flags()
	f.StringVar(&resultObject.cloudID, content.CmdFlagCloudIDLong, content.None, content.CmdFlagCloudIDDescription)
	return cmd
}

func (t *resultObjectCmd) run() error {
	res, err := t.client.ShowResultObject(t.teamID, t.cloudID)
	if err != nil {
		return err
	}
	b := make([]interface{}, len(res))
	for i := range res {
		b[i] = res[i]
	}
	util.PrettyPrintJSON(b)

	return nil
}