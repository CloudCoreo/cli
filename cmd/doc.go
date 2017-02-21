// Copyright © 2016 Paul Allen <paul@cloudcoreo.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

type docsCmd struct {
	out           io.Writer
	dest          string
	docTypeString string
	topCmd        *cobra.Command
}

func newDocsCmd(out io.Writer) *cobra.Command {
	dc := &docsCmd{out: out}

	cmd := &cobra.Command{
		Use:    content.CmdDocsUse,
		Short:  content.CmdDocsShort,
		Long:   content.CmdDocsLong,
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			dc.topCmd = cmd.Root()
			return dc.run()
		},
	}

	f := cmd.Flags()
	f.StringVarP(&dc.dest, content.CmdFlagDirectoryLong, content.CmdFlagDirectoryShort, "./", content.CmdFlagDirectoryDescription)
	f.StringVarP(&dc.docTypeString, content.CmdFlagTypeLong, content.CmdFlagTypeShort, "markdown", content.CmdFlagTypeDescription)

	return cmd
}

func (d *docsCmd) run() error {
	switch d.docTypeString {
	case "markdown", "mdown", "md":
		return doc.GenMarkdownTree(d.topCmd, d.dest)
	case "man":
		manHdr := &doc.GenManHeader{Title: "COREO", Section: "1"}
		return doc.GenManTree(d.topCmd, manHdr, d.dest)
	case "bash":
		return d.topCmd.GenBashCompletionFile(filepath.Join(d.dest, "completions.bash"))
	default:
		return fmt.Errorf(content.ErrorDocGeneration, d.docTypeString)
	}
}
