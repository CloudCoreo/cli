package main

import (
	"io"

	"fmt"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type tokenDeleteCmd struct {
	out     io.Writer
	client  coreo.Interface
	tokenID string
}

func newTokenDeleteCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	tokenDelete := &tokenDeleteCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdDeleteUse,
		Short: content.CmdTokenDeleteShort,
		Long:  content.CmdTokenDeleteLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckTokenShowOrDeleteFlag(tokenDelete.tokenID, verbose); err != nil {
				return err
			}

			if tokenDelete.client == nil {
				tokenDelete.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			return tokenDelete.run()
		},
	}

	f := cmd.Flags()

	f.StringVarP(&tokenDelete.tokenID, content.CmdFlagTokenIDLong, "", "", content.CmdFlagTokenIDDescription)

	return cmd
}

func (t *tokenDeleteCmd) run() error {
	err := t.client.DeleteTokenByID(t.tokenID)
	if err != nil {
		return err
	}

	fmt.Fprintln(t.out, content.InfoTokenDeleted)

	return nil
}
