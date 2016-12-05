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
	//CmdCloudUse command
	CmdCloudUse = "cloud"

	//CmdCloudShort short description
	CmdCloudShort = "Manage Cloud Accounts"

	//CmdCloudLong long description
	CmdCloudLong = `Connect to your cloud accounts.`

	//CmdCloudListShort short description
	CmdCloudListShort = "Show list of tokens"

	//CmdCloudListLong long description
	CmdCloudListLong = `Show list of tokens.`

	//CmdCloudAddShort short description
	CmdCloudAddShort = "Add a cloud account"

	//CmdCloudAddLong long description
	CmdCloudAddLong = `Add a cloud account.`

	//CmdCloudShowShort short description
	CmdCloudShowShort = "Show a cloud account"

	//CmdCloudShowLong long description
	CmdCloudShowLong = `Show a cloud account.`

	//CmdCloudDeleteShort short desription
	CmdCloudDeleteShort = "Delete a cloud account"

	//CmdCloudDeleteLong long desription
	CmdCloudDeleteLong = `Delete a cloud account.`

	// CmdFlagCloudIDLong flag
	CmdFlagCloudIDLong = "cloud-id"

	// CmdFlagCloudIDDescripton flag description
	CmdFlagCloudIDDescripton = "Coreo cloud id"

	//InfoCloudAccountDeleted info
	InfoCloudAccountDeleted = "Cloud account was deleted"

	//InfoUsingCloudAccount info
	InfoUsingCloudAccount = "[ OK ] Using Cloud Account ID %s\n"

	//ErrorCloudIDRequired error message
	ErrorCloudIDRequired = "[ ERROR ] Cloud Account ID is required for this command. Use flag '--cloud-id'\n"
)
