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
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	hostEnvVar         = "VSS_API_ENDPOINT"
	homeEnvVar         = "VSS_HOME"
	profileEnvVar      = "VSS_PROFILE"
	defaultAPIEndpoint = "https://app.securestate.vmware.com/api"
	defaultProfile     = "default"
)

var (
	coreoHome   string
	userProfile string
	key         string
	teamID      string
	apiEndpoint string
	jsonFormat  bool
	verbose     bool
)

func newRootCmd(out io.Writer) *cobra.Command {
	cobra.OnInitialize(initConfig)
	cmd := &cobra.Command{
		// The first word of Use is the name of this command
		Use:          content.CmdCoreoUse,
		Short:        content.CmdCoreoShort,
		Long:         content.CmdCoreoLong,
		SilenceUsage: true,
	}

	userProfileToUse := os.Getenv(profileEnvVar)
	if userProfileToUse == "" {
		userProfileToUse = defaultProfile
	}

	envAPIEndpoint := os.Getenv(hostEnvVar)
	if envAPIEndpoint == "" {
		envAPIEndpoint = defaultAPIEndpoint
	}

	p := cmd.PersistentFlags()
	p.StringVar(&coreoHome, content.CmdFlagConfigLong, defaultCoreoHome(), content.CmdFlagConfigDescription)
	p.StringVar(&userProfile, content.CmdFlagProfileLong, userProfileToUse, content.CmdFlagProfileDescription)
	p.StringVar(&key, content.CmdFlagAPIKeyLong, content.None, content.CmdFlagAPIKeyDescription)
	p.StringVar(&teamID, content.CmdFlagTeamIDLong, content.None, content.CmdFlagTeamIDDescription)
	p.StringVar(&apiEndpoint, content.CmdFlagAPIEndpointLong, envAPIEndpoint, content.CmdFlagAPIEndpointDescription)
	p.BoolVar(&jsonFormat, content.CmdFlagJSONLong, false, content.CmdFlagJSONDescription)
	p.BoolVar(&verbose, content.CmdFlagVerboseLong, false, content.CmdFlagVerboseDescription)
	cmd.AddCommand(
		newVersionCmd(out),
		newTeamCmd(out),
		newTokenCmd(out),
		newCloudAccountCmd(out),
		newConfigureCmd(out),
		newCompletionCmd(out),
		newResultCmd(out),
		// Hidden documentation generator command: 'coreo docs'
		newDocsCmd(out),
		newEventCmd(out),
	)

	return cmd
}

func main() {
	cmd := newRootCmd(os.Stdout)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	viper.SetConfigType("yaml")
	viper.SetConfigName("profiles") // name of config file (without extension)
	viper.AddConfigPath(coreoHome)  // adding home directory as first search path
	viper.AutomaticEnv()            // read in environment variables that match

	path := homePath()

	if err := util.CreateFolder("", path); err != nil {
		fmt.Println("Error creating folder")
	}

	if err := util.CreateFile(content.DefaultFile, path, "", false); err != nil {
		fmt.Println("Error creating file")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", viper.ConfigFileUsed())
	}
}

func setupCoreoConfig(cmd *cobra.Command, args []string) error {
	return setupCoreoCredentials(cmd, args)
}

func setupCoreoCredentials(cmd *cobra.Command, args []string) error {
	apiKey, err := util.CheckAPIKeyFlag(key, userProfile)

	if err != nil {
		return err

	}
	key = apiKey

	if verbose {
		fmt.Printf(content.InfoUsingProfile, userProfile)
	}

	return nil
}

func defaultCoreoHome() string {
	if home := os.Getenv(homeEnvVar); home != "" {
		return home
	}

	return filepath.Join(os.Getenv("HOME"), content.DefaultFolder)
}

func homePath() string {
	return os.ExpandEnv(coreoHome)
}
