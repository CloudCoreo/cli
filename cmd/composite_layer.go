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

type compositeLayerCmd struct {
	out        io.Writer
	directory  string
	name       string
	gitRepoURL string
	serverDir  bool
}

func newCompositeLayerCmd(out io.Writer) *cobra.Command {
	compositeLayer := &compositeLayerCmd{
		out: out,
	}

	cmd := &cobra.Command{
		Use:   content.CmdLayerUse,
		Short: content.CmdCompositeLayerShort,
		Long:  content.CmdCompositeLayerLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := util.CheckLayersFlags(compositeLayer.name, compositeLayer.gitRepoURL); err != nil {
				return err
			}

			return compositeLayer.run()
		},
	}

	f := cmd.Flags()

	f.StringVarP(&compositeLayer.directory, content.CmdFlagDirectoryLong, content.CmdFlagDirectoryShort, "", content.CmdFlagDirectoryDescription)
	f.StringVarP(&compositeLayer.name, content.CmdFlagNameLong, content.CmdFlagNameShort, "", content.CmdFlagNameDescription)
	f.StringVarP(&compositeLayer.gitRepoURL, content.CmdFlagGitRepoLong, content.CmdFlagGitRepoShort, "", content.CmdFlagGitRepoDescription)
	f.BoolVarP(&compositeLayer.serverDir, content.CmdFlagServerLong, content.CmdFlagServerShort, false, content.CmdFlagServerDescription)

	return cmd
}

func (t *compositeLayerCmd) run() error {

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
	fmt.Println(content.CmdCompositeLayerSuccess)

	// generate override and service files
	genContent(t.directory)

	if t.serverDir {
		genServerContent(t.directory)
	}

	return nil
}
