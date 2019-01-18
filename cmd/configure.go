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
	"github.com/CloudCoreo/cli/pkg/command"
	"github.com/spf13/cobra"

	"fmt"

	"github.com/CloudCoreo/cli/cmd/util"
)

type configureCmd struct {
	out         io.Writer
	client      command.Interface
	teamID      string
	compositeID string
}

func newConfigureCmd(out io.Writer) *cobra.Command {

	configure := &configureCmd{
		out: out,
	}

	cmd := &cobra.Command{
		Use:     content.CmdConfigureUse,
		Short:   content.CmdConfigureShort,
		Long:    content.CmdConfigureLong,
		Example: content.CmdConfigureExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			return configure.run()
		},
	}

	cmd.AddCommand(newConfigureListCmd(out))

	return cmd
}

func (t *configureCmd) run() error {

	//generate config keys based on user profile
	apiKey := fmt.Sprintf("%s.%s", userProfile, content.AccessKey)
	secretKey := fmt.Sprintf("%s.%s", userProfile, content.SecretKey)
	teamIDKey := fmt.Sprintf("%s.%s", userProfile, content.TeamID)

	userAPIkey := ""
	userSecretKey := ""
	userTeamID := ""

	// load from config
	apiKeyValue := util.GetValueFromConfig(apiKey, true)
	secretKeyValue := util.GetValueFromConfig(secretKey, true)
	teamIDValue := util.GetValueFromConfig(teamIDKey, false)

	// prompt user for input
	setValue(&userAPIkey, key, fmt.Sprintf(content.CmdConfigurePromptAPIKEY, apiKeyValue))
	setValue(&userSecretKey, secret, fmt.Sprintf(content.CmdConfigurePromptSecretKEY, secretKeyValue))
	setValue(&teamIDValue, teamID, fmt.Sprintf(content.CmdConfigurePromptTeamID, teamIDValue))

	// replace values in config
	util.UpdateConfig(apiKey, userAPIkey)
	util.UpdateConfig(secretKey, userSecretKey)
	util.UpdateConfig(teamIDKey, userTeamID)

	// save config
	if err := util.SaveViperConfig(); err != nil {
		println("Unable to save config")
	}
	return nil
}

func getValueFromUser(userKey *string, prompt string) {
	fmt.Print(prompt)
	fmt.Scanln(userKey)
}

func setValue(userKey *string, current string, prompt string) {
	if current != content.None {
		*userKey = current
	} else {
		getValueFromUser(userKey, prompt)
	}
}
