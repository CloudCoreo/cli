package main

import (
	"io"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

type compositeCreateCmd struct {
	out        io.Writer
	client     coreo.Interface
	teamID     string
	gitRepoURL string
	name       string
}

func newCompositeCreateCmd(client coreo.Interface, out io.Writer) *cobra.Command {
	compositeCreate := &compositeCreateCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdCreateUse,
		Short: content.CmdCompositeCreateShort,
		Long:  content.CmdCompositeCreateLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckCompositeCreateFlags(compositeCreate.name, compositeCreate.gitRepoURL); err != nil {
				return err
			}

			if compositeCreate.client == nil {
				compositeCreate.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			compositeCreate.teamID = teamID

			return compositeCreate.run()
		},
	}

	f := cmd.Flags()

	f.StringVarP(&compositeCreate.name, content.CmdFlagNameLong, content.CmdFlagNameShort, "", content.CmdFlagNameDescription)
	f.StringVarP(&compositeCreate.gitRepoURL, content.CmdFlagGitRepoLong, content.CmdFlagGitRepoShort, "", content.CmdFlagGitRepoDescription)

	return cmd
}

func (t *compositeCreateCmd) run() error {
	composite, err := t.client.CreateComposite(t.teamID, t.gitRepoURL, t.name)
	if err != nil {
		return err
	}

	util.PrintResult(
		t.out,
		composite,
		[]string{"ID", "Name", "TeamID"},
		map[string]string{
			"ID":     "Composite ID",
			"Name":   "Composite Name",
			"TeamID": "Team ID",
		},
		json,
		verbose)

	return nil
}
