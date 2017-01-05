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
	"os"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/CloudCoreo/cli/cmd/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Profile struct for user
type Profile struct {
	ProfileName string
	APIKey      string
	SecretKey   string
	TeamID      string
}

type configureListCmd struct {
	out io.Writer
}

func newConfigureListCmd(out io.Writer) *cobra.Command {
	compositeInit := &configureListCmd{
		out: out,
	}

	cmd := &cobra.Command{
		Use:   content.CmdListUse,
		Short: content.CmdConfigureListShort,
		Long:  content.CmdConfigureListLong,
		RunE: func(cmd *cobra.Command, args []string) error {
			return compositeInit.run()
		},
	}

	return cmd
}

func (t *configureListCmd) run() error {

	var config interface{}
	err := viper.Unmarshal(&config)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}

	profiles := []*Profile{}
	for k := range config.(map[string]interface{}) {

		//generate config keys based on user profile
		apiKey := fmt.Sprintf("%s.%s", k, content.AccessKey)
		secretKey := fmt.Sprintf("%s.%s", k, content.SecretKey)
		teamIDKey := fmt.Sprintf("%s.%s", k, content.TeamID)

		profile := &Profile{
			ProfileName: k,
			APIKey:      util.GetValueFromConfig(apiKey, true),
			SecretKey:   util.GetValueFromConfig(secretKey, true),
			TeamID:      util.GetValueFromConfig(teamIDKey, false),
		}

		profiles = append(profiles, profile)

	}

	if len(profiles) == 0 {
		fmt.Println(content.ErrorNoUserProfileFound)
		os.Exit(-1)
	}

	b := make([]interface{}, len(profiles))
	for i := range profiles {
		b[i] = profiles[i]
	}

	table := util.NewTable()
	table.UseObj(b)
	fmt.Println(table.Render())
	return nil
}
