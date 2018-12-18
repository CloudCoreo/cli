package main

import (
	"io"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/spf13/cobra"
)

func newEventCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:               content.CmdEventUse,
		Short:             content.CmdEventShort,
		Long:              content.CmdEventLong,
		PersistentPreRunE: setupCoreoConfig,
	}
	cmd.AddCommand(newEventSetupCmd(nil, nil, out))
	return cmd
}
