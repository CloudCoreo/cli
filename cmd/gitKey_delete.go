package main

import (
	"io"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type gitKeyDeleteCmd struct {
	out      io.Writer
	client   coreo.Interface
	teamID   string
	gitKeyID string
}

func newGitKeyDeleteCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	gitKeyDelete := &gitKeyDeleteCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdDeleteUse,
		Short: content.CmdGitKeyDeleteShort,
		Long:  content.CmdGitKeyDeleteLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckGitKeyShowOrDeleteFlag(gitKeyDelete.gitKeyID, verbose); err != nil {
				return err
			}

			if gitKeyDelete.client == nil {
				gitKeyDelete.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			gitKeyDelete.teamID = teamID

			return gitKeyDelete.run()
		},
	}
	f := cmd.Flags()

	f.StringVarP(&gitKeyDelete.gitKeyID, content.CmdFlagGitKeyIDLong, "", "", content.CmdFlagGitKeyIDDescription)

	return cmd
}

func (t *gitKeyDeleteCmd) run() error {
	err := t.client.DeleteGitKeyByID(t.teamID, t.gitKeyID)
	if err != nil {
		return err
	}

	return nil
}
