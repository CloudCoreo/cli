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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
	"runtime"
	"github.com/cloudcoreo/cli/cmd/content"
	"github.com/cloudcoreo/cli/cmd/util"
)

var key, secret, teamID,  userProfile, cfgFile, resourceKey, resourceSecret, resourceName, format string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   content.CMD_COREO_USE,
	Short: content.CMD_COREO_SHORT,
	Long: content.CMD_COREO_LONG,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cloudcoreo/profiles.yaml)")
	RootCmd.PersistentFlags().StringVar(&userProfile, "profile", "default", "user profile (default)")
	RootCmd.PersistentFlags().StringVar(&key, "api-key", content.NONE, "Coreo api key")
	RootCmd.PersistentFlags().StringVar(&secret, "api-secret", content.NONE, "Coreo secret key")
	RootCmd.PersistentFlags().StringVar(&teamID, "team-id", content.NONE, "Coreo team id")
	RootCmd.PersistentFlags().StringVar(&format, "format", "table", "Output format")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	viper.SetConfigType("yaml")
	viper.SetConfigName("profiles") // name of config file (without extension)
	viper.AddConfigPath("$HOME/.cloudcoreo")       // adding home directory as first search path
	viper.AutomaticEnv()               // read in environment variables that match


	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	} else {
		path := absPathify("$HOME")

		if err := util.CreateFolder(content.DEFAULT_FOLDER, path); err != nil {
			fmt.Println("Error creating folder")
		}

		if err := util.CreateFile(content.DEFAULT_FILE, filepath.Join(path, content.DEFAULT_FOLDER), "", false); err != nil {
			fmt.Println("Error creating file")
		}
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", viper.ConfigFileUsed())
	}
}

func SetupCoreoCredentials() {
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

	tID, err := util.CheckTeamIDFlag(teamID, userProfile)

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