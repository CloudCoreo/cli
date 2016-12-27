package main

import (
	"io"

	"fmt"
	"os"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/spf13/cobra"
)

type compositeExtendsCmd struct {
	out        io.Writer
	directory  string
	gitRepoURL string
	serverDir  bool
}

func newCompositeExtendsCmd(out io.Writer) *cobra.Command {
	compositeExtends := &compositeExtendsCmd{
		out: out,
	}

	cmd := &cobra.Command{
		Use:   content.CmdExtendsUse,
		Short: content.CmdCompositeExtendsShort,
		Long:  content.CmdCompositeExtendsLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckExtendFlags(compositeExtends.gitRepoURL); err != nil {
				return err
			}

			return compositeExtends.run()
		},
	}

	f := cmd.Flags()

	f.StringVarP(&compositeExtends.directory, content.CmdFlagDirectoryLong, content.CmdFlagDirectoryShort, "", content.CmdFlagDirectoryDescription)
	f.BoolVarP(&compositeExtends.serverDir, content.CmdFlagServerLong, content.CmdFlagServerShort, false, content.CmdFlagServerDescription)
	f.StringVarP(&compositeExtends.gitRepoURL, content.CmdFlagGitRepoLong, content.CmdFlagGitRepoShort, "", content.CmdFlagGitRepoDescription)

	return cmd
}

func (t *compositeExtendsCmd) run() error {

	if err := util.CheckGitInstall(); err != nil {
		return err
	}

	if t.directory == "" {
		t.directory, _ = os.Getwd()
	}

	err := util.CreateGitSubmodule(t.directory, t.gitRepoURL)

	if err != nil {
		return err
	}

	fmt.Fprintln(t.out, content.CmdCompositeExtendsSuccess)

	// generate override and service files
	genContent(t.directory)

	if t.serverDir {
		genServerContent(t.directory)
	}

	return nil
}
