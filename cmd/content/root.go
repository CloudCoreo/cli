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

	//SecretKey secret key
	SecretKey = "SECRET_KEY"

	//TeamID team id
	TeamID = "TEAM_ID"

	//DefaultFolder defautl folder
	DefaultFolder = ".cloudcoreo"

	//DefaultFile default file
	DefaultFile = "profiles.yaml"

	//None none
	None = "None"

	//EndpointAddress endpoint
	EndpointAddress = "https://app.cloudcoreo.com/api"

	//MaskLength mask length
	MaskLength = 4

	//Mask mask
	Mask = "****************"

	//ErrorMissingFile error
	ErrorMissingFile = "[ ERROR ] Missing %s file, creating it.\n"

	//ErrorDirectoryPathNotProvided error
	ErrorDirectoryPathNotProvided = "[ ERROR ] Directory path not provided.\n"

	//ErrorGitURLNotProvided error
	ErrorGitURLNotProvided = "[ ERROR ] Git url not provided. Use flag -g\n"

	//ErrorInvalidDirectory error
	ErrorInvalidDirectory = "[ ERROR ] Error switching to directory to %s. Please make sure it's a valid directory.\n"

	//ErrorGitSubmoduleFailed error
	ErrorGitSubmoduleFailed = "[ ERROR ] git submodule add failed with:\n%s\n"

	//ErrorGitRepoURLMissing error
	ErrorGitRepoURLMissing = "[ ERROR ] A SSH git repo url is required: -g\n"

	//ErrorInvalidGitRepoURL error
	ErrorInvalidGitRepoURL = "[ ERROR ] Use a SSH git repo url for example : [-g git@github.com:CloudCoreo/audit-aws.git]\n"

	//ErrorNameMissing error
	ErrorNameMissing = "[ ERROR ] Name is required for this command. Use flag --name\n"

	//ErrorKeyMissing error
	ErrorKeyMissing = "[ ERROR ] Key is required for this command. Use flag --key\n"

	//ErrorSecretMissing error
	ErrorSecretMissing = "[ ERROR ] Secret is required for this command. Use flag --secret\n"

	//ErrorGitInitNotRan error
	ErrorGitInitNotRan = "[ ERROR ] Not a git repository (or any of the parent directories): .git\n"

	//ErrorNoUserProfileFound error
	ErrorNoUserProfileFound = "[ ERROR ] No user profile found. Set profile using 'coreo configure' command\n"

	//ErrorAcceptsNoArgs error
	ErrorAcceptsNoArgs = "[ ERROR ] This command accepts no argument(s).\n"

	//InfoCreatingGitSubmodule info
	InfoCreatingGitSubmodule = "Creating gitsubmodule %s in %s...\n"

	//CmdFlagDirectoryLong cmd
	CmdFlagDirectoryLong = "directory"

	//CmdFlagDirectoryShort cmd
	CmdFlagDirectoryShort = "d"

	//CmdFlagDirectoryDescription cmd
	CmdFlagDirectoryDescription = "the working directory <fully-qualified-path>"

	//CmdFlagGitRepoLong cmd
	CmdFlagGitRepoLong = "git-repo"

	//CmdFlagGitRepoShort cmd
	CmdFlagGitRepoShort = "g"

	//CmdFlagGitRepoDescription cmd
	CmdFlagGitRepoDescription = "Git repro url"

	//CmdFlagNameLong cmd
	CmdFlagNameLong = "name"

	//CmdFlagNameShort cmd
	CmdFlagNameShort = "n"

	//CmdFlagNameDescription cmd
	CmdFlagNameDescription = "Name flag"

	//CmdFlagServerLong cmd
	CmdFlagServerLong = "add-server-directories"

	//CmdFlagServerShort cmd
	CmdFlagServerShort = "s"

	//CmdFlagServerDescription cmd
	CmdFlagServerDescription = "Create server files"

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
	CmdFlagAPISecretDescription = "Coreo API Secret"

	//CmdFlagAPIKeyLong api key flag long
	CmdFlagAPIKeyLong = "api-key"

	//CmdFlagAPIKeyDescription api key flag description
	CmdFlagAPIKeyDescription = "Coreo API Key"

	//CmdFlagTeamIDLong team id long
	CmdFlagTeamIDLong = "team-id"

	//CmdFlagTeamIDDescription team id flag description
	CmdFlagTeamIDDescription = "Coreo team id"

	//CmdFlagJSONLong json long
	CmdFlagJSONLong = "json"

	//CmdFlagJSONDescription json flag description
	CmdFlagJSONDescription = "Output in json format"

	//CmdFlagVerboseLong verbose long
	CmdFlagVerboseLong = "verbose"

	//CmdFlagVerboseDescription verbose flag description
	CmdFlagVerboseDescription = "Enable verbose output"

	//CmdFlagConfigLong config flag long
	CmdFlagConfigLong = "config"

	//CmdFlagConfigDescription config flag description
	CmdFlagConfigDescription = "Config file (default is $HOME/.cloudcoreo/profiles.yaml)"

	//CmdFlagProfileLong profile flag long
	CmdFlagProfileLong = "profile"

	//CmdFlagProfileDescription secret flag description
	CmdFlagProfileDescription = "Coreo profile to use."

	//CmdCoreoUse Coreo cmd
	CmdCoreoUse = "coreo"

	//CmdCoreoShort Coreo short description
	CmdCoreoShort = "CloudCoreo CLI"

	//CmdCoreoLong Coreo long description
	CmdCoreoLong = `CloudCoreo CLI.`

	//CmdConfigureUse configure cmd
	CmdConfigureUse = "configure"

	//CmdListUse list cmd
	CmdListUse = "list"

	//CmdCreateUse create cmd
	CmdCreateUse = "create"

	//CmdAddUse add cmd
	CmdAddUse = "add"

	//CmdDeleteUse delete cmd
	CmdDeleteUse = "delete"

	//CmdShowUse show cmd
	CmdShowUse = "show"

	//CmdDisableUse disable cmd
	CmdDisableUse = "disable"

	//CmdEnableUse enable cmd
	CmdEnableUse = "enable"

	//CmdRunUse run cmd
	CmdRunUse = "run"

	//CmdInitUse init cmd
	CmdInitUse = "init"

	//CmdGendocUse gendoc cmd
	CmdGendocUse = "gendoc"

	//CmdLayerUse layer cmd
	CmdLayerUse = "layer"

	//CmdExtendsUse extends cmd
	CmdExtendsUse = "extends"

	//InfoUsingProfile info using profile
	InfoUsingProfile = "[ OK ] Using Profile %s\n"
)
