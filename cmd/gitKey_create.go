package main

import (
	"io"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type gitKeyCreateCmd struct {
	out            io.Writer
	client         coreo.Interface
	teamID         string
	resourceSecret string
	resourceName   string
}

func newGitKeyCreateCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	gitKeyCreate := &gitKeyCreateCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdAddUse,
		Short: content.CmdGitKeyAddShort,
		Long:  content.CmdGitKeyAddLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckGitKeyAddFlags(gitKeyCreate.resourceName, gitKeyCreate.resourceSecret); err != nil {
				return err
			}

			if gitKeyCreate.client == nil {
				gitKeyCreate.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			gitKeyCreate.teamID = teamID

			return gitKeyCreate.run()
		},
	}
	f := cmd.Flags()

	f.StringVarP(&gitKeyCreate.resourceSecret, content.CmdFlagSecretLong, content.CmdFlagSecretShort, "", content.CmdFlagSecretDescription)
	f.StringVarP(&gitKeyCreate.resourceName, content.CmdFlagNameLong, content.CmdFlagNameShort, "", content.CmdFlagNameDescription)

	return cmd
}

func (t *gitKeyCreateCmd) run() error {
	gitKey, err := t.client.CreateGitKey(t.teamID, t.resourceSecret, t.resourceName)
	if err != nil {
		return err
	}

	util.PrintResult(
		t.out,
		gitKey,
		[]string{"ID", "Name", "TeamID"},
		map[string]string{
			"ID":     "Git Key ID",
			"Name":   "Git Key Name",
			"TeamID": "Team ID",
		},
		json,
		verbose)

	return nil
}
