package main

import (
	"io"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type tokenShowCmd struct {
	out     io.Writer
	client  coreo.Interface
	tokenID string
}

func newTokenShowCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	tokenShow := &tokenShowCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdShowUse,
		Short: content.CmdTokenShowShort,
		Long:  content.CmdTokenShowLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckTokenShowOrDeleteFlag(tokenShow.tokenID, verbose); err != nil {
				return err
			}

			if tokenShow.client == nil {
				tokenShow.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			return tokenShow.run()
		},
	}

	f := cmd.Flags()

	f.StringVarP(&tokenShow.tokenID, content.CmdFlagTokenIDLong, "", "", content.CmdFlagTokenIDDescription)

	return cmd
}

func (t *tokenShowCmd) run() error {
	token, err := t.client.ShowTokenByID(t.tokenID)
	if err != nil {
		return err
	}

	util.PrintResult(
		t.out,
		token,
		[]string{"ID", "Name", "Description"},
		map[string]string{
			"ID":          "Token ID",
			"Name":        "Token Name",
			"Description": "Token Description",
		},
		json,
		verbose)

	return nil
}
