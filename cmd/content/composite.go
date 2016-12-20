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
	//CmdCompositeUse command
	CmdCompositeUse = "composite"

	//CmdCompositeShort short descripton
	CmdCompositeShort = "Manage Composites"

	//CmdCompositeLong long descripton
	CmdCompositeLong = `Manage Composites. Composites are inheritable cloud reference designs that
can be easily extended and updated. A Composite can specify a single micro-service,
a group of services, a complex application, or an entire data center definition.`

	//CmdCompositeInitShort short descripton
	CmdCompositeInitShort = "Create a new configuration"

	//CmdCompositeInitLong long descripton
	CmdCompositeInitLong = `Configure  Coreo  CLI  options.`

	//CmdCompositeInitSuccess message
	CmdCompositeInitSuccess = "[Done] Composite initialization completed, default files and folders were generated.\n"

	//CmdCompositeCreateShort short descripton
	CmdCompositeCreateShort = "Create new composite"

	//CmdCompositeCreateLong long descripton
	CmdCompositeCreateLong = `Create new composite.`

	//CmdCompositeGendocShort short descripton
	CmdCompositeGendocShort = "Generate readme documentation"

	//CmdCompositeGendocLong long descripton
	CmdCompositeGendocLong = `Generate readme documentation.`

	//CmdCompositeGendocSuccess message
	CmdCompositeGendocSuccess = "[Done] Composite documentation completed, README.md was generated.\n"

	//CmdCompositeListShort short descripton
	CmdCompositeListShort = "Show list of composites"

	//CmdCompositeListLong long descripton
	CmdCompositeListLong = `Show list of composites.`

	//CmdCompositeShowShort short descripton
	CmdCompositeShowShort = "Show a particular composite"

	//CmdCompositeShowLong long descripton
	CmdCompositeShowLong = `Show a particular compsoite`

	//CmdCompositeLayerShort short descripton
	CmdCompositeLayerShort = "Generate composite layer"

	//CmdCompositeLayerLong long descripton
	CmdCompositeLayerLong = `Generate composite layer.`

	//CmdCompositeLayerSuccess layer success
	CmdCompositeLayerSuccess = "Composite successfully added."

	//CmdCompositeExtendsShort short descripton
	CmdCompositeExtendsShort = "Extends a particular composite"

	//CmdCompositeExtendsLong long descripton
	CmdCompositeExtendsLong = `Extends a particular composite.`

	//CmdCompositeExtendsSuccess message
	CmdCompositeExtendsSuccess = "[Done] Composite successfully extended.\n"

	//InfoUsingCompsiteID informational using Composite id
	InfoUsingCompsiteID = "[ OK ] Using Composite ID %s\n"

	//ErrorCompositeIDRequired error message
	ErrorCompositeIDRequired = "[ ERROR ] Composite ID is required for this command. Use flag '--composite-id'\n"

	//CmdFlagCompositeIDLong flag
	CmdFlagCompositeIDLong = "composite-id"

	//CmdFlagCompositeIDDescription flag description
	CmdFlagCompositeIDDescription = "Coreo composite id"

	//DefaultFilesConfigYAMLContent default content
	DefaultFilesConfigYAMLContent = `## This file was auto-generated by CloudCoreo CLI
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

	//DefaultFilesConfigRBContent default content
	DefaultFilesConfigRBContent = `## This file was auto-generated by CloudCoreo CLI
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

	//DefaultFilesConfigYAMLName default content
	DefaultFilesConfigYAMLName = "config.yaml"

	//DefaultFilesConfigRBName default content
	DefaultFilesConfigRBName = "config.rb"

	//DefaultFilesReadMEName default content
	DefaultFilesReadMEName = "README.md"

	//DefaultFilesOrderYAMLName default content
	DefaultFilesOrderYAMLName = "order.yaml"

	//DefaultFilesOverrideFolderName default content
	DefaultFilesOverrideFolderName = "/overrides"

	//DefaultFilesOperationalScriptsFolder default content
	DefaultFilesOperationalScriptsFolder = "/operational-scripts"

	//DefaultFilesBootScriptsFolder default content
	DefaultFilesBootScriptsFolder = "/boot-scripts"

	//DefaultFilesShutdownScriptsFolder default content
	DefaultFilesShutdownScriptsFolder = "/shutdown-scripts"

	//DefaultFilesOperationalReadMeContent default content
	DefaultFilesOperationalReadMeContent = `## This file was auto-generated by CloudCoreo CLI.
	Scripts contained in this directory will be exposed to the CloudCoreo UI and can be run on an adhoc basis.`

	//DefaultFilesShutDownReadMeContent default content
	DefaultFilesShutDownReadMeContent = `## This file was auto-generated by CloudCoreo CLI
	Scripts contained in this directory will be proccessed (if possible) when an instance "is asked to shut down.
	The order in which these scripts will be run is defined by the order.yaml`

	//DefaultFilesShutDownOrderYAMLContent default content
	DefaultFilesShutDownOrderYAMLContent = `## This file was auto-generated by CloudCoreo CLI
	## this yaml file dictates the order in which the scripts will run. i.e:\n" +
	## script-order:
	##   - deregister_dns.sh
	##   - delete_backups.sh`

	//DefaultFilesBootReadMeContent default content
	DefaultFilesBootReadMeContent = `## This file was auto-generated by CloudCoreo CLI
	Scripts contained in this directory will be proccessed when an instance is booting.
	The order in which these scripts will be run is defined by the order.yaml`

	//DefaultFilesBootOrderYAMLContent default content
	DefaultFilesBootOrderYAMLContent = `## This file was auto-generated by CloudCoreo CLI
	## this yaml file dictates the order in which the scripts will run. i.e
	## script-order:
	##   - install_chef.sh
	##   - run_chef.sh"`

	//DefautlFilesGenDocReadMeRquiredNoDefaultHeader default content
	DefautlFilesGenDocReadMeRquiredNoDefaultHeader = "Required variables with no default"

	//DefautlFilesGenDocReadMeRquiredDefaultHeader default content
	DefautlFilesGenDocReadMeRquiredDefaultHeader = "Required variables with default"

	//DefautlFilesGenDocReadMeNoRquiredNoDefaultHeader default content
	DefautlFilesGenDocReadMeNoRquiredNoDefaultHeader = "Optional variables with no default"

	//DefautlFilesGenDocReadMeNoRquiredDefaultHeader default content
	DefautlFilesGenDocReadMeNoRquiredDefaultHeader = "Optional variables with default"

	//DefaultFilesOverridesReadMeHeader default content
	DefaultFilesOverridesReadMeHeader = `## This file was auto-generated by CloudCoreo CLI
	Anything contained in this directory will act as an override for the stack in which it is contained.

	The paths should be considered relative to the parent of this directory.

	For example, if you have a directory structure like this,
	`

	//DefaultFilesOverridesReadMeFooter default content
	DefaultFilesOverridesReadMeFooter = `Because the directory structure within the override directory matches the structure of the parent,
	the 'order.yaml' file will be ignored in the stack-a directory and instead the overrides/stack-a order.yaml
	file will be used.`

	//DefaultFilesReadMeCodeTicks default content
	DefaultFilesReadMeCodeTicks = "```"

	//DefaultFilesOverridesReadMeTree default content
	DefaultFilesOverridesReadMeTree = `
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

	//DefaultFilesServicesFolder default content
	DefaultFilesServicesFolder = "/services"

	//DefaultFilesServicesReadMeHeader code snippet default content
	DefaultFilesServicesReadMeHeader = `## This file was auto-generated by CloudCoreo CLI
	This is your services directory. Place a config.rb file in here containing CloudCoreo service
	syntax. For example, your config.rb might contain the following in order to create a VPC
	`

	//DefaultFilesServicesReadMeCode code snippet default content
	DefaultFilesServicesReadMeCode = `%s
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
