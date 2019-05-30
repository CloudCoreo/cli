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
	//AccessKey api key
	AccessKey = "API_KEY"

	//TeamID team id
	TeamID = "TEAM_ID"

	//DefaultFolder default folder
	DefaultFolder = ".vss"

	//DefaultFile default file
	DefaultFile = "profiles.yaml"

	//None none
	None = "None"

	//MaskLength mask length
	MaskLength = 4

	//Mask mask
	Mask = "****************"

	//ErrorAPIKeyMissing error
	ErrorAPIKeyMissing = "VMware Secure State API Key is required for this command. Use flag --api-key\n"

	//ErrorSecretMissing error
	ErrorSecretMissing = "Secret is required for this command. Use flag --secret\n"

	//ErrorAPISecretMissing error
	ErrorAPISecretMissing = "API Secret is required for this command. Use flag --api-secret\n"

	//ErrorGitInitNotRan error
	ErrorGitInitNotRan = "Not a git repository (or any of the parent directories): .git\n"

	//ErrorNoUserProfileFound error
	ErrorNoUserProfileFound = "No user profile found. Set profile using 'vss configure' command"

	//ErrorAcceptsNoArgs error
	ErrorAcceptsNoArgs = "This command accepts no argument(s).\n"

	//InfoCreatingGitSubmodule info
	InfoCreatingGitSubmodule = "Creating gitsubmodule %s in %s...\n"

	//CmdFlagDirectoryLong cmd
	CmdFlagDirectoryLong = "directory"

	//CmdFlagDirectoryShort cmd
	CmdFlagDirectoryShort = "d"

	//CmdFlagDirectoryDescription cmd
	CmdFlagDirectoryDescription = "The working directory <fully-qualified-path>"

	//CmdFlagNameLong cmd
	CmdFlagNameLong = "name"

	//CmdFlagNameShort cmd
	CmdFlagNameShort = "n"

	//CmdFlagDescriptionLong cmd
	CmdFlagDescriptionLong = "description"

	//CmdFlagDescriptionShort cmd
	CmdFlagDescriptionShort = "d"

	//CmdFlagNameDescription cmd
	CmdFlagNameDescription = "Name flag"

	//CmdFlagKeyLong cmd
	CmdFlagKeyLong = "key"

	//CmdFlagKeyShort cmd
	CmdFlagKeyShort = "K"

	//CmdFlagKeyDescription cmd
	CmdFlagKeyDescription = "Resource key"

	//CmdFlagSecretLong secret flag long
	CmdFlagSecretLong = "secret"

	//CmdFlagSecretShort secret flag short
	CmdFlagSecretShort = "S"

	//CmdFlagSecretDescription secret flag description
	CmdFlagSecretDescription = "Resource secret"

	//CmdFlagAPISecretLong secret flag long
	CmdFlagAPISecretLong = "api-secret"

	//CmdFlagAPISecretDescription api secret flag description
	CmdFlagAPISecretDescription = "VMware Secure State API Secret"

	//CmdFlagAPIKeyLong api key flag long
	CmdFlagAPIKeyLong = "api-key"

	//CmdFlagAPIKeyDescription api key flag description
	CmdFlagAPIKeyDescription = "VMware Secure State API Key"

	//CmdFlagTeamIDLong team id long
	CmdFlagTeamIDLong = "team-id"

	//CmdFlagTeamIDDescription team id flag description
	CmdFlagTeamIDDescription = "VMware Secure State team id"

	//CmdFlagJSONLong json long
	CmdFlagJSONLong = "json"

	//CmdFlagJSONDescription json flag description
	CmdFlagJSONDescription = "Output in json format"

	//CmdFlagVerboseLong verbose long
	CmdFlagVerboseLong = "verbose"

	//CmdFlagVerboseDescription verbose flag description
	CmdFlagVerboseDescription = "Enable verbose output"

	//CmdFlagFileLong JSON file flag
	CmdFlagFileLong = "file"

	//CmdFlagFileShort JSON file flag
	CmdFlagFileShort = "f"

	//CmdFlagConfigLong config flag long
	CmdFlagConfigLong = "home"

	//CmdFlagConfigDescription config flag description
	CmdFlagConfigDescription = "Location of your VMware Secure State config. Overrides $VSS_HOME."

	//CmdFlagProfileLong profile flag long
	CmdFlagProfileLong = "profile"

	//CmdFlagProfileDescription secret flag description
	CmdFlagProfileDescription = "VMware Secure State CLI profile to use. Overrides $VSS_PROFILE."

	//CmdFlagAPIEndpointLong api endpoint flag long
	CmdFlagAPIEndpointLong = "endpoint"

	//CmdFlagAPIEndpointDescription api endpoint description
	CmdFlagAPIEndpointDescription = "VMware Secure State API endpoint. Overrides $VSS_API_ENDPOINT."

	//CmdCoreoUse Coreo cmd
	CmdCoreoUse = "vss"

	//CmdCoreoShort Coreo short description
	CmdCoreoShort = "VMware Secure State CLI"

	//CmdCoreoLong Coreo long description
	CmdCoreoLong = `VMware Secure State CLI.`

	//CmdConfigureUse configure cmd
	CmdConfigureUse = "configure"

	//CmdListUse list cmd
	CmdListUse = "list"

	//CmdAddUse add cmd
	CmdAddUse = "add"

	//CmdUpdateUse update cmd
	CmdUpdateUse = "update"

	//CmdTestUse test cmd
	CmdTestUse = "test"

	//CmdScanUse scan cmd
	CmdScanUse = "scan"

	//CmdDeleteUse delete cmd
	CmdDeleteUse = "delete [flags]"

	//CmdShowUse show cmd
	CmdShowUse = "show [flags]"

	//InfoUsingProfile info using profile
	InfoUsingProfile = "[ OK ] Using Profile %s\n"

	//InfoCommandSuccess info command was executed successfully
	InfoCommandSuccess = "[ OK ] Command was executed successfully"
)
