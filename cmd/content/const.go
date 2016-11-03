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
	ACCESS_KEY = "API_KEY"
	SECRET_KEY = "SECRET_KEY"
	TEAM_ID = "TEAM_ID"
	DEFAULT_FOLDER = ".cloudcoreo"
	DEFAULT_FILE = "profiles.yaml"
	NONE = "None"
	ENDPOINT_ADDRESS = "http://localhost:3000/api"
	MASK_LENGTH = 4
	MASK = "****************"

	// errors
	ERROR_MISSING_API_KEY_SECRET_KEY = "Missing API key or/and Secret key. Please run 'Coreo configure' to configure them."
	ERROR_MISSING_FILE = "Missing %s file, creating it."

	//Flags
	CMD_FLAG_DIRECTORY_LONG = "directory"
	CMD_FLAG_DIRECTORY_SHORT = "d"
	CMD_FLAG_DIRECTORY_DESCRIPTION = "the working directory <fully-qualified-path>"


	// version command
	CMD_VERSION_USE = "version"
	CMD_VERSION_SHORT = "Print the version number of Coreo CLI"
	CMD_VERSION_LONG = `Print the version number of Coreo CLI.`
	CMD_VERSION = "Coreo CLI v0.0.1"


	// root command
	CMD_COREO_USE = "coreo"
	CMD_COREO_SHORT = "CloudCoreo CLI"
	CMD_COREO_LONG = `CloudCoreo CLI.`

	// token command
	CMD_TOKEN_USE = "token"
	CMD_TOKEN_SHORT = "CloudCoreo Token command"
	CMD_TOKEN_LONG = `CloudCoreo Token command.`

	// token list command
	CMD_TOKEN_LIST_USE = "list"
	CMD_TOKEN_LIST_SHORT = "Show list of tokens"
	CMD_TOKEN_LIST_LONG = `Show list of tokens.`

	// token add command
	CMD_TOKEN_ADD_USE = "add"
	CMD_TOKEN_ADD_SHORT = "Add a token"
	CMD_TOKEN_ADD_LONG = `Add a token.`


	// token show command
	CMD_TOKEN_SHOW_USE = "show"
	CMD_TOKEN_SHOW_SHORT = "Show a token"
	CMD_TOKEN_SHOW_LONG = `Show a token.`


	// token delete command
	CMD_TOKEN_DELETE_USE = "delete"
	CMD_TOKEN_DELETE_SHORT = "DELETE a token"
	CMD_TOKEN_DELETE_LONG = `DELETE a token.`

	// gitKey command
	CMD_GITKEY_USE = "git-key"
	CMD_GITKEY_SHORT = "CloudCoreo GITKEY command"
	CMD_GITKEY_LONG = `CloudCoreo GITKEY command.`

	// gitKey list command
	CMD_GITKEY_LIST_USE = "list"
	CMD_GITKEY_LIST_SHORT = "Show list of GITKEY"
	CMD_GITKEY_LIST_LONG = `Show list of GITKEY.`

	// gitKey add command
	CMD_GITKEY_ADD_USE = "add"
	CMD_GITKEY_ADD_SHORT = "Add a GITKEY"
	CMD_GITKEY_ADD_LONG = `Add a GITKEY.`


	// gitKey show command
	CMD_GITKEY_SHOW_USE = "show"
	CMD_GITKEY_SHOW_SHORT = "Show a GITKEY"
	CMD_GITKEY_SHOW_LONG = `Show a GITKEY.`


	// gitKey delete command
	CMD_GITKEY_DELETE_USE = "delete"
	CMD_GITKEY_DELETE_SHORT = "DELETE a GITKEY"
	CMD_GITKEY_DELETE_LONG = `DELETE a GITKEY.`

	// team command
	CMD_TEAM_USE = "team"
	CMD_TEAM_SHORT = "CloudCoreo team command"
	CMD_TEAM_LONG = `CloudCoreo team command.`


	// team list command
	CMD_TEAM_LIST_USE = "list"
	CMD_TEAM_LIST_SHORT = "Show list of team"
	CMD_TEAM_LIST_LONG = `Show list of team.`

	// token add command
	CMD_TEAM_ADD_USE = "add"
	CMD_TEAM_ADD_SHORT = "Add a team"
	CMD_TEAM_ADD_LONG = `Add a team.`


	// team show command
	CMD_TEAM_SHOW_USE = "show"
	CMD_TEAM_SHOW_SHORT = "Show a team"
	CMD_TEAM_SHOW_LONG = `Show a team.`


	// TEAM delete command
	CMD_TEAM_DELETE_USE = "delete"
	CMD_TEAM_DELETE_SHORT = "delete a team"
	CMD_TEAM_DELETE_LONG = `delete a team.`

	// cloud command
	CMD_CLOUD_USE = "cloud"
	CMD_CLOUD_SHORT = "CloudCoreo cloud command"
	CMD_CLOUD_LONG = `CloudCoreo cloud command.`

	// cloud list command
	CMD_CLOUD_LIST_USE = "list"
	CMD_CLOUD_LIST_SHORT = "Show list of tokens"
	CMD_CLOUD_LIST_LONG = `Show list of tokens.`

	// cloud add command
	CMD_CLOUD_ADD_USE = "add"
	CMD_CLOUD_ADD_SHORT = "Add a token"
	CMD_CLOUD_ADD_LONG = `Add a token.`


	// cloud show command
	CMD_CLOUD_SHOW_USE = "show"
	CMD_CLOUD_SHOW_SHORT = "Show a token"
	CMD_CLOUD_SHOW_LONG = `Show a token.`


	// cloud delete command
	CMD_CLOUD_DELETE_USE = "delete"
	CMD_CLOUD_DELETE_SHORT = "DELETE a token"
	CMD_CLOUD_DELETE_LONG = `DELETE a token.`

	// configure command
	CMD_CONFIG_USE = "configure"
	CMD_CONFIG_SHORT = "create a new configuration"
	CMD_CONFIG_LONG = `Configure  Coreo  CLI  options.`
	CMD_CONFIG_PROMPT_API_KEY = "Enter CloudCoreo api key[%s]: "
	CMD_CONFIG_PROMPT_SECRET_KEY = "Enter CloudCoreo secret key[%s]: "
	CMD_CONFIG_PROMPT_TEAM_ID = "Enter CloudCoreo team ID[%s]: "


	// composite command
	CMD_COMPOSITE_USE = "composite"
	CMD_COMPOSITE_SHORT = "create a new configuration"
	CMD_COMPOSITE_LONG = `Configure  Coreo  CLI  options.`

	// composite init command
	CMD_COMPOSITE_INIT_USE = "init"
	CMD_COMPOSITE_INIT_SHORT = "create a new configuration"
	CMD_COMPOSITE_INIT_LONG = `Configure  Coreo  CLI  options.`
	CMD_COMPOSITE_INIT_SUCCESS = "Initialization completed, default files were generated"


	// composite gendoc command
	CMD_COMPOSITE_GENDOC_USE = "gendoc"
	CMD_COMPOSITE_GENDOC_SHORT = "generate readme documentation"
	CMD_COMPOSITE_GENDOC_LONG = `generate readme documentation.`
	CMD_COMPOSITE_GENDOC_SUCCESS = "Documentation completed, README.md was generated"

	// composite list command
	CMD_COMPOSITE_LIST_USE = "list"
	CMD_COMPOSITE_LIST_SHORT = "show list of composites"
	CMD_COMPOSITE_LIST_LONG = `show list of composites.`


	DEFAULT_FILES_CONFIG_YAML = `## This file was auto-generated by CloudCoreo CLI
## This file was automatically generated using the CloudCoreo CLI
##
## This config.rb file exists to create and maintain services not related to compute.
## for example, a VPC might be maintained using:
##
## coreo_aws_vpc_vpc "my-vpc" do
##   action :sustain
##   cidr "12.0.0.0/16"
##   internet_gateway true
## end
##
`
	DEFAULT_FILES_CONFIG_RB = `## This file was auto-generated by CloudCoreo CLI
## This file was automatically generated using the CloudCoreo CLI
##
## This config.rb file exists to create and maintain services not related to compute.
## for example, a VPC might be maintained using:
##
## coreo_aws_vpc_vpc "my-vpc" do
##   action :sustain
##   cidr "12.0.0.0/16"
##   internet_gateway true
## end
##
`
	DEFAULT_FILES_CONFIG_YAML_FILE = "config.yaml"

	DEFAULT_FILES_CONFIG_RB_FILE = "config.rb"

	DEFAULT_FILES_README_FILE = "README.md"

	DEFAULT_FILES_OVERRIDES_FOLDER = "/overrides"

	DEFAULT_FILES_GENDOC_README_REQUIRED_NO_DEFAULT_HEADER = "Required variables with no default"

	DEFAULT_FILES_GENDOC_README_REQUIRED_DEFAULT_HEADER = "Required variables with default"

	DEFAULT_FILES_GENDOC_README_NO_REQUIRED_DEFAULT_HEADER = "Optional variables with no default"

	DEFAULT_FILES_GENDOC_README_NO_REQUIRED_NO_DEFAULT_HEADER = "Optional variables with default"

	DEFAULT_FILES_OVERRIDES_README_HEADER = `## This file was auto-generated by CloudCoreo CLI
Anything contained in this directory will act as an override for the stack in which it is contained.

The paths should be considered relative to the parent of this directory.

For example, if you have a directory structure like this,
`

	DEFAULT_FILES_OVERRIDES_README_FOOTER = `Because the directory structure within the override directory matches the structure of the parent,
the 'order.yaml' file will be ignored in the stack-a directory and instead the overrides/stack-a order.yaml
file will be used.`

	DEFAULT_FILES_README_CODE_TICKS = "```"

	DEFAULT_FILES_OVERRIDES_README_TREE = `
%s
+-- parent
|   +-- overrides
|   |   +-- stack-a
|   |   |   +-- boot-scripts
|   |   |   |   +-- order.yaml
|   +-- stack-a
|   |   +-- boot-scripts
|   |   |   +-- order.yaml
%s
`
	DEFAULT_FILES_SERVICES_FOLDER = "/services"

	DEFAULT_FILES_SERVICES_README_HEADER = `## This file was auto-generated by CloudCoreo CLI
This is your services directory. Place a config.rb file in here containing CloudCoreo service
syntax. For example, your config.rb might contain the following in order to create a VPC
`

	DEFAULT_FILES_SERVICES_README_CODE = `%s
	coreo_aws_vpc_vpc "my-vpc" do
  action :sustain
  cidr "12.0.0.0/16"
  internet_gateway true
end

coreo_aws_vpc_routetable "my-public-route" do
  action :create
  vpc "my-vpc"
  number_of_tables 3
  routes [
         { :from => "0.0.0.0/0", :to => "my-vpc", :type => :igw },
        ]
end

coreo_aws_vpc_subnet "my-public-subnet" do
  action :create
  number_of_zones 3
  percent_of_vpc_allocated 50
  route_table "my-public-route"
  vpc "my-vpc"
  map_public_ip_on_launch true
end
%s`
)
