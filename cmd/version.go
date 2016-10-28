package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Coreo CLI",
	Long:  `Print the version number of Coreo CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Coreo CLI v0.0.1")
	},
}