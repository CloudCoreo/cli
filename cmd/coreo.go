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
	"runtime"
	"strings"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	hostEnvVar         = "CC_API_ENDPOINT"
	defaultAPIEndpoint = "https://app.cloudcoreo.com/api"
)

var (
	cfgFile     string
	userProfile string
	key         string
	secret      string
	teamID      string
	apiEndpoint string
	json        bool
	verbose     bool
)

func newRootCmd(out io.Writer) *cobra.Command {
	cobra.OnInitialize(initConfig)
	cmd := &cobra.Command{
		Use:          content.CmdCoreoUse,
		Short:        content.CmdCoreoShort,
		Long:         content.CmdCoreoLong,
		SilenceUsage: true,
	}

	envAPIEndpoint := os.Getenv(hostEnvVar)

	if envAPIEndpoint == "" {
		envAPIEndpoint = defaultAPIEndpoint
	}

	p := cmd.PersistentFlags()
	p.StringVar(&cfgFile, content.CmdFlagConfigLong, "", content.CmdFlagConfigDescription)
	p.StringVar(&userProfile, content.CmdFlagProfileLong, "default", content.CmdFlagProfileDescription)
	p.StringVar(&key, content.CmdFlagAPIKeyLong, content.None, content.CmdFlagAPIKeyDescription)
	p.StringVar(&secret, content.CmdFlagAPISecretLong, content.None, content.CmdFlagAPISecretDescription)
	p.StringVar(&teamID, content.CmdFlagTeamIDLong, content.None, content.CmdFlagTeamIDDescription)
	p.StringVar(&apiEndpoint, content.CmdFlagAPIEndpointLong, envAPIEndpoint, content.CmdFlagAPIEndpointDescription)
	p.BoolVar(&json, content.CmdFlagJSONLong, false, content.CmdFlagJSONDescription)
	p.BoolVar(&verbose, content.CmdFlagVerboseLong, true, content.CmdFlagVerboseDescription)
	cmd.AddCommand(
		newVersionCmd(out),
		newTeamCmd(out),
		newTokenCmd(out),
		newCloudAccountCmd(out),
		newGitKeyCmd(out),
		newCompositeCmd(out),
		newPlanCmd(out),
		newConfigureCmd(out),
		newLintCmd(out),
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
	viper.SetConfigName("profiles")          // name of config file (without extension)
	viper.AddConfigPath("$HOME/.cloudcoreo") // adding home directory as first search path
	viper.AutomaticEnv()                     // read in environment variables that match

	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	} else {
		path := absPathify("$HOME")

		if err := util.CreateFolder(content.DefaultFolder, path); err != nil {
			fmt.Println("Error creating folder")
		}

		if err := util.CreateFile(content.DefaultFile, filepath.Join(path, content.DefaultFolder), "", false); err != nil {
			fmt.Println("Error creating file")
		}
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", viper.ConfigFileUsed())
	}
}

func setupCoreoConfig(*cobra.Command, []string) error {
	setupCoreoCredentials()
	setupCoreoDefaultTeam()

	// TODO return valid error
	return nil
}

func setupCoreoCredentials() {
	apiKey, err := util.CheckAPIKeyFlag(key, userProfile)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}
	key = apiKey

	secretKey, err := util.CheckSecretKeyFlag(secret, userProfile)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}
	secret = secretKey

	if verbose {
		fmt.Printf(content.InfoUsingProfile, userProfile)
	}
}

func setupCoreoDefaultTeam() {
	tID, err := util.CheckTeamIDFlag(teamID, userProfile, verbose)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}
	teamID = tID
}

func absPathify(inPath string) string {
	if strings.HasPrefix(inPath, "$HOME") {
		inPath = userHomeDir() + inPath[5:]
	}

	if strings.HasPrefix(inPath, "$") {
		end := strings.Index(inPath, string(os.PathSeparator))
		inPath = os.Getenv(inPath[1:end]) + inPath[end:]
	}

	if filepath.IsAbs(inPath) {
		return filepath.Clean(inPath)
	}

	p, err := filepath.Abs(inPath)
	if err == nil {
		return filepath.Clean(p)
	}
	return ""
}

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
