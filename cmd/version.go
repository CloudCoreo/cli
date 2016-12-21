package main

import (
	"fmt"
	"io"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/spf13/cobra"
)

const versionDesc = `
Show the client and server versions for Helm and tiller.

This will print a representation of the client and server versions of Helm and
Tiller. The output will look something like this:

Client: &version.Version{SemVer:"v2.0.0", GitCommit:"ff52399e51bb880526e9cd0ed8386f6433b74da1", GitTreeState:"clean"}
Server: &version.Version{SemVer:"v2.0.0", GitCommit:"b0c113dfb9f612a9add796549da66c0d294508a3", GitTreeState:"clean"}

- SemVer is the semantic version of the release.
- GitCommit is the SHA for the commit that this version was built from.
- GitTreeState is "clean" if there are no local code changes when this binary was
  built, and "dirty" if the binary was built from locally modified code.

To print just the client version, use '--client'. To print just the server version,
use '--server'.
`

type versionCmd struct {
	out           io.Writer
	clientVersion string
	clientGithash string
	clientBuildID string
}

var (
	version string
	gitHash string
	buildID string
)

func newVersionCmd(out io.Writer) *cobra.Command {
	v := &versionCmd{
		out:           out,
		clientVersion: version,
		clientGithash: gitHash,
		clientBuildID: buildID,
	}

	cmd := &cobra.Command{
		Use:   content.CmdVersionUse,
		Short: content.CmdVersionShort,
		Long:  content.CmdVersionLong,
		RunE: func(cmd *cobra.Command, args []string) error {
			return v.run()
		},
	}

	return cmd
}

func (v *versionCmd) run() error {
	fmt.Fprintf(v.out, "Version: %#v\n", v.clientVersion)
	fmt.Fprintf(v.out, "Git hash: %#v\n", v.clientGithash)
	fmt.Fprintf(v.out, "BuildID: %#v\n", v.clientBuildID)

	return nil
}
