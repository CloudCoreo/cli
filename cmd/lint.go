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
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/pkg/lint"
	"github.com/CloudCoreo/cli/pkg/lint/support"

	"github.com/spf13/cobra"
)

type lintCmd struct {
	strict bool
	paths  []string
	out    io.Writer
}

func newLintCmd(out io.Writer) *cobra.Command {
	l := &lintCmd{
		paths: []string{"."},
		out:   out,
	}
	cmd := &cobra.Command{
		Use:   content.CmdLintUse,
		Short: content.CmdLintShort,
		Long:  content.CmdLintLong,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				l.paths = args
			}
			return l.run()
		},
	}

	cmd.Flags().BoolVar(&l.strict, "strict", false, "fail on lint warnings")

	return cmd
}

var errLintNoComposite = errors.New("no composite found for linting (missing Config.yaml)")

func (l *lintCmd) run() error {
	var lowestTolerance int
	if l.strict {
		lowestTolerance = support.WarningSev
	} else {
		lowestTolerance = support.ErrorSev
	}

	var total int
	var failures int
	for _, path := range l.paths {
		if linter, err := lintComposite(path); err != nil {
			fmt.Println("==> Skipping", path)
			fmt.Println(err)
		} else {
			fmt.Println("==> Linting", path)

			if len(linter.Messages) == 0 {
				fmt.Println("Lint OK")
			}

			for _, msg := range linter.Messages {
				fmt.Println(msg)
			}

			total = total + 1
			if linter.HighestSeverity >= lowestTolerance {
				failures = failures + 1
			}
		}
		fmt.Println("")
	}

	msg := fmt.Sprintf("%d composite(s) linted", total)
	if failures > 0 {
		return fmt.Errorf("%s, %d composite(s) failed", msg, failures)
	}

	fmt.Fprintf(l.out, "%s, no failures\n", msg)

	return nil
}

func lintComposite(path string) (support.Linter, error) {
	linter := support.Linter{}

	// Guard: Error out of this is not composite.
	if _, err := os.Stat(filepath.Join(path, "config.yaml")); err != nil {
		return linter, errLintNoComposite
	}

	return lint.All(path), nil
}
