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

package content

const (
	//CmdConfigureShort short description
	CmdConfigureShort = "Configure CLI options"

	//CmdConfigureLong long description
	CmdConfigureLong = `Configure  CLI  options. If this command is run with no arguments,
you will be prompted for configuration values such as your  VMware Secure State  Access
Key  Id  and your  VMware Secure State  Secret  Access  Key.  You can configure a named pro-
file using the --profile argument.  If your config file does not  exist
(the default location is $HOME/.vss/profiles.yaml), the VSS CLI
will create it for you.  To keep an existing value, hit enter when prompted
for the value. When  you  are prompted for information, the current value
will be displayed in [brackets].  If the config item has no value,
it be displayed as  [None].

Note:  the  values  you  provide  for the VMware Secure State Access Key ID and the VMware SecureState
Secret Access Key will  be  written  to  the  shared  credentials  file
($HOME/.vss/profiles.yaml).

`

	//CmdConfigureExample is examples for vss configure command
	CmdConfigureExample = `  vss configure
  vss configure --api-key VSS_API_KEY --api-secret VSS_API_SECRET --team-id VSS_TEAM_ID
  vss configure list`

	//CmdConfigurePromptAPIKEY prompt for api key
	CmdConfigurePromptAPIKEY = "Enter your VMware Secure State api token key (available on https://app.cloudcoreo.com under Settings -> API Tokens) [%s]: "

	//CmdConfigurePromptSecretKEY prompt for secret key
	CmdConfigurePromptSecretKEY = "Enter your VMware Secure State api token secret key (available on https://app.cloudcoreo.com under Settings -> API Tokens) [%s]: "

	//CmdConfigurePromptTeamID prompt for team ID
	CmdConfigurePromptTeamID = "Enter your default VMware Secure State team ID (available on https://app.cloudcoreo.com under Settings -> Team) [%s]: "

	//CmdConfigureListShort short description
	CmdConfigureListShort = "List all user profiles"

	//CmdConfigureListLong long description
	CmdConfigureListLong = `List all user profiles.`
)
