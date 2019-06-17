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
	"fmt"
	"io"

	"github.com/CloudCoreo/cli/client"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/command"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type resultObjectCmd struct {
	client   command.Interface
	teamID   string
	cloudID  string
	out      io.Writer
	level    string
	provider string
	retry    uint
}

//Object is violation by objects
type Object struct {
	ID string `json:"objectName"`
	client.Info
	RiskScore int    `json:"riskScore"`
	TeamName  string `json:"teamName"`
	TeamID    string `json:"teamID"`
	Region    string `json:"region"`
}

//ObjectWrapper contains info other than object details
type ObjectWrapper struct {
	AccountName   string   `json:"accountName,omitempty"`
	AccountNumber string   `json:"accountNumber,omitempty"`
	TotalItems    int      `json:"totalItems"`
	Objects       []Object `json:"violations"`
}

func newResultObjectCmd(client command.Interface, out io.Writer) *cobra.Command {
	resultObject := &resultObjectCmd{
		client: client,
		out:    out,
	}
	cmd := &cobra.Command{
		Use:     content.CmdResultObjectUse,
		Short:   content.CmdResultObjectShort,
		Long:    content.CmdResultObjectLong,
		Example: content.CmdResultObjectExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			if resultObject.client == nil {
				resultObject.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.RefreshToken(key))
			}

			resultObject.teamID = teamID
			return resultObject.run()
		},
	}
	f := cmd.Flags()
	f.StringVar(&resultObject.cloudID, content.CmdFlagCloudIDLong, content.None, content.CmdFlagCloudIDDescription)
	f.StringVar(&resultObject.level, content.CmdFlagLevelLong, "", content.CmdFlagLevelDescription)
	f.StringVar(&resultObject.provider, content.CmdFlagProvider, "", content.CmdFlagProviderDescription)
	f.UintVar(&resultObject.retry, content.CmdFlagRetry, 1, content.CmdFlagRetryDescription)
	return cmd
}

func (t *resultObjectCmd) run() error {
	res, err := t.client.ShowResultObject(t.teamID, t.cloudID, t.level, t.provider, t.retry)
	if err != nil {
		return err
	}
	return t.prettyPrintObjects(res)
}

func (t *resultObjectCmd) prettyPrintObjects(wrappers []*client.ResultObjectWrapper) error {
	result := make([]ObjectWrapper, len(wrappers))
	for i, wrapper := range wrappers {
		result[i].AccountName = wrapper.AccountName
		result[i].AccountNumber = wrapper.AccountNumber
		result[i].TotalItems = wrapper.TotalItems
		result[i].Objects = make([]Object, len(wrapper.Objects))
		for j, object := range wrapper.Objects {
			result[i].Objects[j].TeamName = object.TInfo.Name
			result[i].Objects[j].TeamID = object.TInfo.ID
			result[i].Objects[j].RiskScore = object.RiskScore
			result[i].Objects[j].ID = object.ID
			result[i].Objects[j].Info = object.Info
			result[i].Objects[j].Region = object.Region
		}
	}
	_, err := fmt.Fprintln(t.out, util.PrettyJSON(result))
	return err
}
