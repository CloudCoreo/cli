package main

import (
	"io"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/spf13/cobra"

	"fmt"

	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"
)

type configureCmd struct {
	out         io.Writer
	client      coreo.Interface
	teamID      string
	compositeID string
}

func newConfigureCmd(out io.Writer) *cobra.Command {

	configure := &configureCmd{
		out: out,
	}

	cmd := &cobra.Command{
		Use:   content.CmdConfigureUse,
		Short: content.CmdConfigureShort,
		Long:  content.CmdConfigureLong,
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
		getValueFromUser(&userAPIkey, fmt.Sprintf(content.CmdConfigurePromptAPIKEY, apiKeyValue))
		getValueFromUser(&userSecretKey, fmt.Sprintf(content.CmdConfigurePromptSecretKEY, secretKeyValue))
		getValueFromUser(&userTeamID, fmt.Sprintf(content.CmdConfigurePromptTeamID, teamIDValue))
	}

	// replace values in config
	util.UpdateConfig(apiKey, userAPIkey)
	util.UpdateConfig(secretKey, userSecretKey)
	util.UpdateConfig(teamIDKey, userTeamID)

	// save config
	util.SaveViperConfig()
	return nil
}

func getValueFromUser(userKey *string, prompt string) {
	fmt.Print(prompt)
	fmt.Scanln(userKey)
}
