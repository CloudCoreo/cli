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
	"io"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/CloudCoreo/cli/pkg/command"
	"github.com/CloudCoreo/cli/pkg/coreo"
	"github.com/spf13/cobra"
)

func newTokenCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:               content.CmdTokenUse,
		Short:             content.CmdTokenShort,
		Long:              content.CmdTokenLong,
		PersistentPreRunE: setupCoreoConfig,
	}

	cmd.AddCommand(newTokenListCmd(nil, out))
	cmd.AddCommand(newTokenShowCmd(nil, out))
	cmd.AddCommand(newTokenDeleteCmd(nil, out))

	return cmd
}

type tokenListCmd struct {
	out    io.Writer
	client command.Interface
}

func newTokenListCmd(client command.Interface, out io.Writer) *cobra.Command {
	tokenList := &tokenListCmd{
		out:    out,
		client: client,
	}

	cmd := &cobra.Command{
		Use:   content.CmdListUse,
		Short: content.CmdTokenListShort,
		Long:  content.CmdTokenListLong,
		RunE: func(cmd *cobra.Command, args []string) error {

			if tokenList.client == nil {
				tokenList.client = coreo.NewClient(
					coreo.Host(apiEndpoint),
					coreo.APIKey(key),
					coreo.SecretKey(secret))
			}

			return tokenList.run()
		},
	}

	return cmd
}

func (t *tokenListCmd) run() error {
	Tokens, err := t.client.ListTokens()
	if err != nil {
		return err
	}

	b := make([]interface{}, len(Tokens))
	for i := range Tokens {
		b[i] = Tokens[i]
	}

	util.PrintResult(
		t.out,
		b,
		[]string{"ID", "Name", "Description"},
		map[string]string{
			"ID":          "Token ID",
			"Name":        "Token Name",
			"Description": "Token Description",
		},
		jsonFormat,
		verbose)

	return nil
}
