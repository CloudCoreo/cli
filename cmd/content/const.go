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
	ACCESS_KEY       = "API_KEY"
	SECRET_KEY       = "SECRET_KEY"
	TEAM_ID          = "TEAM_ID"
	DEFAULT_FOLDER   = ".cloudcoreo"
	DEFAULT_FILE     = "profiles.yaml"
	NONE             = "None"
	ENDPOINT_ADDRESS = "https://staging.hack.cloudcoreo.com/api"
	MASK_LENGTH      = 4
	MASK             = "****************"

	// errors
	ERROR_MISSING_API_KEY_SECRET_KEY  = "Missing API key or/and Secret key. Please run 'Coreo configure' to configure them."
	ERROR_MISSING_FILE                = "Missing %s file, creating it."
	ERROR_DIRECTORY_PATH_NOT_PROVIDED = "Directory path not provided"
	ERROR_GIT_URL_NOT_PROVIDED        = "Git url not provided"
	ERROR_INVALID_DIRECTORY           = "Error switching to directory to %s. Please make sure it's a valid directory"
	ERROR_GIT_SUBMODULE_FAILED        = "git submodule add failed with:\n%s\n"
	ERROR_GIT_REPO_URL_MISSING        = "A SSH git repo url is required: -g"
	ERROR_INVALID_GIT_REPO_URL        = "Use a SSH git repo url for example : [-g git@github.com:CloudCoreo/audit-aws.git]"
	ERROR_NAME_MISSING                = "Name is required for this command"
	ERROR_KEY_MISSING                 = "Key is required for this command"
	ERROR_SECRET_MISSING              = "Secret is required for this command"
	ERROR_DESCRIPTION_MISSING         = "Description flag is required for this command"
	ERROR_ID_MISSING                  = "ID is required for this command"
	ERROR_CLOUDACCOUNT_DELETED        = "Cloud account was deleted"

	// info
	INFO_CREATING_GITSUBMODULE = "Creating gitsubmodule %s in %s...\n"

	//Flags
	CMD_FLAG_DIRECTORY_LONG        = "directory"
	CMD_FLAG_DIRECTORY_SHORT       = "d"
	CMD_FLAG_DIRECTORY_DESCRIPTION = "the working directory <fully-qualified-path>"

	CMD_FLAG_GIT_REPO_LONG        = "git-repo"
	CMD_FLAG_GIT_REPO_SHORT       = "g"
	CMD_FLAG_GIT_REPO_DESCRIPTION = "Git repro url"

	CMD_FLAG_NAME_LONG        = "name"
	CMD_FLAG_NAME_SHORT       = "n"
	CMD_FLAG_NAME_DESCRIPTION = "Name flag"

	CMD_FLAG_DESCRIPTION_LONG        = "description"
	CMD_FLAG_DESCRIPTION_SHORT       = "d"
	CMD_FLAG_DESCRIPTION_DESCRIPTION = "Description flag"

	CMD_FLAG_SERVER_LONG        = "add-server-directories"
	CMD_FLAG_SERVER_SHORT       = "s"
	CMD_FLAG_SERVER_DESCRIPTION = "Create server files"

	CMD_FLAG_ID_LONG  = "id"
	CMD_FLAG_ID_SHORT = "i"

	CMD_FLAG_PLAN_ID_LONG  = "plan-id"
	CMD_FLAG_PLAN_ID_SHORT = "p"

	CMD_FLAG_CLOUDID_DESCRIPTION   = "Coreo cloud id"
	CMD_FLAG_TOKENID_DESCRIPTION   = "Coreo token id"
	CMD_FLAG_COMPOSITE_DESCRIPTION = "Coreo composite id"
	CMD_FLAG_GITKEY_DESCRIPTION    = "Coreo gitkey id"
	CMD_FLAG_PLANID_DESCRIPTION    = "Coreo plan id"

	CMD_FLAG_KEY_LONG        = "key"
	CMD_FLAG_KEY_SHORT       = "K"
	CMD_FLAG_KEY_DESCRIPTION = "Resource key"

	CMD_FLAG_SECRET_LONG        = "secret"
	CMD_FLAG_SECRET_SHORT       = "S"
	CMD_FLAG_SECRET_DESCRIPTION = "Resource secret"

	// version command
	CMD_VERSION_USE   = "version"
	CMD_VERSION_SHORT = "Print the version number of Coreo CLI"
	CMD_VERSION_LONG  = `Print the version number of Coreo CLI.`

	// root command
	CMD_COREO_USE   = "coreo"
	CMD_COREO_SHORT = "CloudCoreo CLI"
	CMD_COREO_LONG  = `CloudCoreo CLI.`

	CMD_LIST_USE    = "list"
	CMD_CREATE_USE  = "create"
	CMD_ADD_USE     = "add"
	CMD_DELETE_USE  = "delete"
	CMD_SHOW_USE    = "show"
	CMD_DISABLE_USE = "disable"
	CMD_RUN_USE     = "run"

	// plan command
	CMD_PLAN_USE   = "plan"
	CMD_PLAN_SHORT = "CloudCoreo plan command"
	CMD_PLAN_LONG  = `CloudCoreo plan command.`

	// plan list command
	CMD_PLAN_LIST_SHORT = "CloudCoreo plan list command"
	CMD_PLAN_LIST_LONG  = `CloudCoreo plan list command.`

	// plan create command
	CMD_PLAN_CREATE_SHORT = "CloudCoreo plan create command"
	CMD_PLAN_CREATE_LONG  = `CloudCoreo plan create command.`

	// plan delete command
	CMD_PLAN_DELETE_SHORT = "CloudCoreo plan delete command"
	CMD_PLAN_DELETE_LONG  = `CloudCoreo plan delete command.`

	// plan show command
	CMD_PLAN_SHOW_SHORT = "CloudCoreo plan show command"
	CMD_PLAN_SHOW_LONG  = `CloudCoreo plan show command.`

	// plan disable command
	CMD_PLAN_DISABLE_SHORT = "CloudCoreo plan disable command"
	CMD_PLAN_DISABLE_LONG  = `CloudCoreo plan disable command.`

	// plan run command
	CMD_PLAN_RUN_SHORT = "CloudCoreo plan run command"
	CMD_PLAN_RUN_LONG  = `CloudCoreo plan run command.`

	// token command
	CMD_TOKEN_USE   = "token"
	CMD_TOKEN_SHORT = "CloudCoreo Token command"
	CMD_TOKEN_LONG  = `CloudCoreo Token command.`

	// token list command
	CMD_TOKEN_LIST_USE   = "list"
	CMD_TOKEN_LIST_SHORT = "Show list of tokens"
	CMD_TOKEN_LIST_LONG  = `Show list of tokens.`

	// token add command
	CMD_TOKEN_ADD_USE   = "add"
	CMD_TOKEN_ADD_SHORT = "Add a token"
	CMD_TOKEN_ADD_LONG  = `Add a token.`

	// token show command
	CMD_TOKEN_SHOW_USE   = "show"
	CMD_TOKEN_SHOW_SHORT = "Show a token"
	CMD_TOKEN_SHOW_LONG  = `Show a token.`

	// token delete command
	CMD_TOKEN_DELETE_USE   = "delete"
	CMD_TOKEN_DELETE_SHORT = "delete a token"
	CMD_TOKEN_DELETE_LONG  = `delete a token.`

	// gitKey command
	CMD_GITKEY_USE   = "git-key"
	CMD_GITKEY_SHORT = "CloudCoreo gitKey command"
	CMD_GITKEY_LONG  = `CloudCoreo gitKey command.`

	// gitKey list command
	CMD_GITKEY_LIST_USE   = "list"
	CMD_GITKEY_LIST_SHORT = "Show list of gitKey"
	CMD_GITKEY_LIST_LONG  = `Show list of gitKey.`

	// gitKey add command
	CMD_GITKEY_ADD_USE   = "add"
	CMD_GITKEY_ADD_SHORT = "Add a gitKey"
	CMD_GITKEY_ADD_LONG  = `Add a gitKey.`

	// gitKey show command
	CMD_GITKEY_SHOW_USE   = "show"
	CMD_GITKEY_SHOW_SHORT = "Show a gitKey"
	CMD_GITKEY_SHOW_LONG  = `Show a gitKey.`

	// gitKey delete command
	CMD_GITKEY_DELETE_USE   = "delete"
	CMD_GITKEY_DELETE_SHORT = "Delete a gitKey"
	CMD_GITKEY_DELETE_LONG  = `Delete a gitKey.`

	// team command
	CMD_TEAM_USE   = "team"
	CMD_TEAM_SHORT = "CloudCoreo team command"
	CMD_TEAM_LONG  = `CloudCoreo team command.`

	// team list command
	CMD_TEAM_LIST_USE   = "list"
	CMD_TEAM_LIST_SHORT = "Show list of team"
	CMD_TEAM_LIST_LONG  = `Show list of team.`

	// token add command
	CMD_TEAM_ADD_USE   = "add"
	CMD_TEAM_ADD_SHORT = "Add a team"
	CMD_TEAM_ADD_LONG  = `Add a team.`

	// team show command
	CMD_TEAM_SHOW_USE   = "show"
	CMD_TEAM_SHOW_SHORT = "Show a team"
	CMD_TEAM_SHOW_LONG  = `Show a team.`

	// TEAM delete command
	CMD_TEAM_DELETE_USE   = "delete"
	CMD_TEAM_DELETE_SHORT = "delete a team"
	CMD_TEAM_DELETE_LONG  = `delete a team.`

	// cloud command
	CMD_CLOUD_USE   = "cloud"
	CMD_CLOUD_SHORT = "CloudCoreo cloud command"
	CMD_CLOUD_LONG  = `CloudCoreo cloud command.`

	// cloud list command
	CMD_CLOUD_LIST_USE   = "list"
	CMD_CLOUD_LIST_SHORT = "Show list of tokens"
	CMD_CLOUD_LIST_LONG  = `Show list of tokens.`

	// cloud add command
	CMD_CLOUD_ADD_USE   = "add"
	CMD_CLOUD_ADD_SHORT = "Add a token"
	CMD_CLOUD_ADD_LONG  = `Add a token.`

	// cloud show command
	CMD_CLOUD_SHOW_USE   = "show"
	CMD_CLOUD_SHOW_SHORT = "Show a token"
	CMD_CLOUD_SHOW_LONG  = `Show a token.`

	// cloud delete command
	CMD_CLOUD_DELETE_USE   = "delete"
	CMD_CLOUD_DELETE_SHORT = "DELETE a token"
	CMD_CLOUD_DELETE_LONG  = `DELETE a token.`

	// configure command
	CMD_CONFIG_USE               = "configure"
	CMD_CONFIG_SHORT             = "create a new configuration"
	CMD_CONFIG_LONG              = `Configure  Coreo  CLI  options.`
	CMD_CONFIG_PROMPT_API_KEY    = "Enter CloudCoreo api key[%s]: "
	CMD_CONFIG_PROMPT_SECRET_KEY = "Enter CloudCoreo secret key[%s]: "
	CMD_CONFIG_PROMPT_TEAM_ID    = "Enter CloudCoreo team ID[%s]: "

	// composite command
	CMD_COMPOSITE_USE   = "composite"
	CMD_COMPOSITE_SHORT = "create a new configuration"
	CMD_COMPOSITE_LONG  = `Configure  Coreo  CLI  options.`

	// composite init command
	CMD_COMPOSITE_INIT_USE     = "init"
	CMD_COMPOSITE_INIT_SHORT   = "create a new configuration"
	CMD_COMPOSITE_INIT_LONG    = `Configure  Coreo  CLI  options.`
	CMD_COMPOSITE_INIT_SUCCESS = "Initialization completed, default files were generated"

	// composite create command
	CMD_COMPOSITE_CREATE_USE   = "create"
	CMD_COMPOSITE_CREATE_SHORT = "create readme documentation"
	CMD_COMPOSITE_CREATE_LONG  = `create readme documentation.`

	// composite gendoc command
	CMD_COMPOSITE_GENDOC_USE     = "gendoc"
	CMD_COMPOSITE_GENDOC_SHORT   = "generate readme documentation"
	CMD_COMPOSITE_GENDOC_LONG    = `generate readme documentation.`
	CMD_COMPOSITE_GENDOC_SUCCESS = "Documentation completed, README.md was generated"

	// composite list command
	CMD_COMPOSITE_LIST_USE   = "list"
	CMD_COMPOSITE_LIST_SHORT = "show list of composites"
	CMD_COMPOSITE_LIST_LONG  = `show list of composites.`

	// composite show command
	CMD_COMPOSITE_SHOW_USE   = "show"
	CMD_COMPOSITE_SHOW_SHORT = "show a particular composite"
	CMD_COMPOSITE_SHOW_LONG  = `show a particular compsoite`

	// composite layer command
	CMD_COMPOSITE_LAYER_USE     = "layer"
	CMD_COMPOSITE_LAYER_SHORT   = "generate composite layer"
	CMD_COMPOSITE_LAYER_LONG    = `generate composite layer.`
	CMD_COMPOSITE_LAYER_SUCCESS = "Composite successfully added."

	// composite extends command
	CMD_COMPOSITE_EXTENDS_USE     = "extends"
	CMD_COMPOSITE_EXTENDS_SHORT   = "generate composite layer"
	CMD_COMPOSITE_EXTENDS_LONG    = `generate composite layer.`
	CMD_COMPOSITE_EXTENDS_SUCCESS = "Composite successfully extended."

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

	DEFAULT_FILES_ORDER_YAML_FILE = "order.yaml"

	DEFAULT_FILES_OVERRIDES_FOLDER = "/overrides"

	DEFAULT_FILES_OPERATIONAL_SCRIPTS_FOLDER = "/operational-scripts"

	DEFAULT_FILES_BOOT_SCRIPTS_FOLDER = "/boot-scripts"

	DEFAULT_FILES_SHUTDOWN_SCRIPTS_FOLDER = "/shutdown-scripts"

	DEFAULT_FILES_OPERATIONAL_README_CONTENT = `## This file was auto-generated by CloudCoreo CLI.
Scripts contained in this directory will be exposed to the CloudCoreo UI and can be run on an adhoc basis.`

	DEFAULT_FILES_SHUTDOWN_README_CONTENT = `## This file was auto-generated by CloudCoreo CLI
Scripts contained in this directory will be proccessed (if possible) when an instance "is asked to shut down.
The order in which these scripts will be run is defined by the order.yaml`

	DEFAULT_FILES_SHUTDOWN_ORDER_YAML_CONTENT = `## This file was auto-generated by CloudCoreo CLI
## this yaml file dictates the order in which the scripts will run. i.e:\n" +
## script-order:
##   - deregister_dns.sh
##   - delete_backups.sh`

	DEFAULT_FILES_BOOT_README_CONTENT = `## This file was auto-generated by CloudCoreo CLI
Scripts contained in this directory will be proccessed when an instance is booting.
The order in which these scripts will be run is defined by the order.yaml`

	DEFAULT_FILES_BOOT_ORDER_YAML_CONTENT = `## This file was auto-generated by CloudCoreo CLI
## this yaml file dictates the order in which the scripts will run. i.e
## script-order:
##   - install_chef.sh
##   - run_chef.sh"`

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
