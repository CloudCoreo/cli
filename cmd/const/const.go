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

package util


const (
	ACCESS_KEY = "API_KEY"
	SECRET_KEY = "SECRET_KEY"
	TEAM_ID = "TEAM_ID"
	DEFAULT_FOLDER = "/.cloudcoreo"
	DEFAULT_FILE = "profiles.yaml"
	NONE = "None"

	// root command
	CMD_COREO_USE = "cloudcoreo"
	CMD_COREO_SHORT = "A brief description of your application"
	CMD_COREO_LONG = `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`

	// Configure command
	CMD_CONFIG_USE = "configure"
	CMD_CONFIG_SHORT = "create a new configuration"
	CMD_CONFIG_LONG = `Configure  Coreo  CLI  options.`
	CMD_CONFIG_PROMPT_API_KEY = "Enter CloudCoreo api key[%s]: "
	CMD_CONFIG_PROMPT_SECRET_KEY = "Enter CloudCoreo secret key[%s]: "
	CMD_CONFIG_PROMPT_TEAM_ID = "Enter CloudCoreo team ID[%s]: "
)