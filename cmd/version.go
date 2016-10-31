package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/cloudcoreo/cli/cmd/content"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   content.CMD_VERSION_USE,
	Short: content.CMD_VERSION_SHORT,
	Long:  content.CMD_VERSION_LONG,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(content.CMD_VERSION)
	},
}