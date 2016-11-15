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

package cmd

import (
	"fmt"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/spf13/cobra"
)

var cmdConfigure = &cobra.Command{
	Use:   content.CMD_CONFIG_USE,
	Short: content.CMD_CONFIG_SHORT,
	Long:  content.CMD_CONFIG_LONG,
	Run: func(cmd *cobra.Command, args []string) {

		//generate config keys based on user profile
		apiKey := fmt.Sprintf("%s.%s", userProfile, content.ACCESS_KEY)
		secretKey := fmt.Sprintf("%s.%s", userProfile, content.SECRET_KEY)
		teamIDKey := fmt.Sprintf("%s.%s", userProfile, content.TEAM_ID)

		userAPIkey := ""
		userSecretKey := ""
		userTeamID := ""

		if key != "None" {
			userAPIkey = key
		}

		if secret != "None" {
			userSecretKey = secret
		}

		if teamID != "None" {
			userTeamID = teamID
		}

		if userAPIkey == "" && userSecretKey == "" && userTeamID == "" {
			// load from config
			apiKeyValue := util.GetValueFromConfig(apiKey, true)
			secretKeyValue := util.GetValueFromConfig(secretKey, true)
			teamIDValue := util.GetValueFromConfig(teamIDKey, false)

			// prompt user for input
			getValueFromUser(&userAPIkey, fmt.Sprintf(content.CMD_CONFIG_PROMPT_API_KEY, apiKeyValue))
			getValueFromUser(&userSecretKey, fmt.Sprintf(content.CMD_CONFIG_PROMPT_SECRET_KEY, secretKeyValue))
			getValueFromUser(&userTeamID, fmt.Sprintf(content.CMD_CONFIG_PROMPT_TEAM_ID, teamIDValue))
		}

		// replace values in config
		util.UpdateConfig(apiKey, userAPIkey)
		util.UpdateConfig(secretKey, userSecretKey)
		util.UpdateConfig(teamIDKey, userTeamID)

		// save config
		util.SaveViperConfig()
	},
}

func getValueFromUser(userKey *string, prompt string) {
	fmt.Print(prompt)
	fmt.Scanln(userKey)
}

func init() {
	RootCmd.AddCommand(cmdConfigure)

}
