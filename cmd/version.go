package cmd

import (
	"fmt"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/spf13/cobra"
)

var version = "No Version Provided"
var buildstamp = "Unknown buildstamp"
var githash = "Unknown githash"

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   content.CMD_VERSION_USE,
	Short: content.CMD_VERSION_SHORT,
	Long:  content.CMD_VERSION_LONG,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\n", version)
		fmt.Printf("Git hash: %s\n", githash)
		fmt.Printf("Buildstamp: %s\n", buildstamp)
	},
}
