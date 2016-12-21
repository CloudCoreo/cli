package main

import (
	"io"

	"fmt"
	"os"
	"path"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/spf13/cobra"
)

type compositeExtendsCmd struct {
	out        io.Writer
	directory  string
	name       string
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

	return cmd
}

func (t *compositeExtendsCmd) run() error {

	if err := util.CheckGitInstall(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}

	if t.directory == "" {
		t.directory, _ = os.Getwd()
	}

	err := util.CreateFolder("stack-"+t.name, t.directory)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}

	t.directory = path.Join(t.directory, "stack-"+t.name)

	err = util.CreateGitSubmodule(t.directory, t.gitRepoURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}
	fmt.Println(content.CmdCompositeExtendsSuccess)

	// generate override and service files
	genContent(t.directory)

	if t.serverDir {
		genServerContent(t.directory)
	}

	return nil
}
