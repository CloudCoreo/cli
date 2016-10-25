package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmdConfigure = &cobra.Command{
	Use:   CMD_CONFIG_USE,
	Short: CMD_CONFIG_SHORT,
	Long: CMD_CONFIG_LONG,
	Run: func(cmd *cobra.Command, args []string) {

		//generate config keys based on user profile
		apiKey := fmt.Sprintf("%s.%s", userProfile, ACCESS_KEY)
		secretKey := fmt.Sprintf("%s.%s", userProfile, SECRET_KEY)
		teamIDKey := fmt.Sprintf("%s.%s", userProfile, TEAM_ID)

		// load from config
		apiKeyValue := getValueFromConfig(apiKey)
		secretKeyValue := getValueFromConfig(secretKey)
		teamIDValue := getValueFromConfig(teamIDKey)

		// prompt user for input
		var userApiKey, userSecretKey, userTeamID string
		getValueFromUser(&userApiKey, fmt.Sprintf(CMD_CONFIG_PROMPT_API_KEY, apiKeyValue))
		getValueFromUser(&userSecretKey, fmt.Sprintf(CMD_CONFIG_PROMPT_SECRET_KEY, secretKeyValue))
		getValueFromUser(&userTeamID, fmt.Sprintf(CMD_CONFIG_PROMPT_TEAM_ID, teamIDValue))

		// replace values in config
		updateConfig(apiKey, userApiKey)
		updateConfig(secretKey, userSecretKey)
		updateConfig(teamIDKey, userTeamID)

		// save config
		SaveViperConfig()
	},
}

func updateConfig(key, value string ) {
	if value != "" {
		viper.Set(key, value)
	}
}

func getValueFromUser(userKey *string, prompt string) {
	fmt.Print(prompt)
	fmt.Scanln(userKey)
}

func getValueFromConfig(key string) (value string) {
	if value = viper.GetString(key); value == "" {
		value = NONE
	}

	return value
}

func init() {
	RootCmd.AddCommand(cmdConfigure)

}
