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

package cmd

import (
	"github.com/cloudcoreo/cli/cmd/content"
	"github.com/spf13/cobra"
)

var planID string

// PlanCmd represents the based command for plan subcommands
var PlanCmd = &cobra.Command{
	Use: content.CMD_PLAN_USE,
	Short: content.CMD_PLAN_SHORT,
	Long: content.CMD_PLAN_LONG,
}

func init() {
	RootCmd.AddCommand(PlanCmd)
}
