package main

import (
	"io"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type gitKeyShowCmd struct {
	out      io.Writer
	client   coreo.Interface
	teamID   string
	gitKeyID string
}

func newGitKeyShowCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	gitKeyShow := &gitKeyShowCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdShowUse,
		Short: content.CmdGitKeyShowShort,
		Long:  content.CmdGitKeyShowLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckGitKeyShowOrDeleteFlag(gitKeyShow.gitKeyID, verbose); err != nil {
				return err
			}

			if gitKeyShow.client == nil {
				gitKeyShow.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			gitKeyShow.teamID = teamID

			return gitKeyShow.run()
		},
	}
	f := cmd.Flags()

	f.StringVarP(&gitKeyShow.gitKeyID, content.CmdFlagGitKeyIDLong, "", "", content.CmdFlagGitKeyIDDescription)

	return cmd
}

func (t *gitKeyShowCmd) run() error {
	gitKey, err := t.client.ShowGitKeyByID(t.teamID, t.gitKeyID)
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
