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
	CmdCloudAddLong = `
Add a cloud account. The result would be shown as the following if successful.
         -----------------------------  -----------------------  -----------------------------
               Cloud Account ID           Cloud Account Name               Team ID
         -----------------------------  -----------------------  -----------------------------
             *********************           ************           ************************
         -----------------------------  -----------------------  -----------------------------
`
	//CmdCloudAddExample ...
	CmdCloudAddExample = `  coreo cloud add --name YOUR_NEW_ACCOUNT_NAME --role NAME_FOR_NEW_ROLE
  coreo cloud add --name YOUR_NEW_ACCOUNT_NAME --arn YOUR_ROLE_ARN --external-id EXTERNAL_ID_OF_YOUR_ROLE`

	//CmdCloudShowShort short description
	CmdCloudShowShort = "Show a cloud account"

	//CmdCloudShowLong long description
	CmdCloudShowLong = `Show a cloud account.`

	//CmdCloudDeleteShort short desription
	CmdCloudDeleteShort = "Delete a cloud account"

	//CmdCloudDeleteLong long desription
	CmdCloudDeleteLong = `Delete a cloud account.`

	//CmdFlagCloudIDLong flag
	CmdFlagCloudIDLong = "cloud-id"

	//CmdFlagRoleName is the name of the role to add a cloud account
	CmdFlagRoleName = "role"

	//CmdFlagRoleNameDescription is description for --role
	CmdFlagRoleNameDescription = "The name of the role you want to create"

	//CmdFlagRoleArn is flag for the Arn of the role you'd like to add for a new cloud account
	CmdFlagRoleArn = "arn"

	//CmdFlagRoleArnDescription is description for --arn
	CmdFlagRoleArnDescription = "The arn of the role to connect"

	//CmdFlagRoleExternalID is flag for external-id used to assume the provided role
	CmdFlagRoleExternalID = "external-id"

	//CmdFlagAwsProfile = "aws-profile"
	CmdFlagAwsProfile = "aws-profile"

	//CmdFlagAwsProfileDescription ...
	CmdFlagAwsProfileDescription = "Aws shared credential file. If empty default provider chain will be used to look for credentials with the following order.\n" +
		"  1. Environment variables.\n" + "  2. Shared credentials file.\n" + "  3. If your application is running on an Amazon EC2 instance, IAM role for Amazon EC2."

	//CmdFlagAwsProfilePath  ...
	CmdFlagAwsProfilePath = "aws-profile-path"

	//CmdFlagAwsProfilePathDescription ...
	CmdFlagAwsProfilePathDescription = "The file path of aws profile. If empty will look for AWS_SHARED_CREDENTIALS_FILE env variable. " +
		"If the env value is empty will default to current user's home directory.\n" + "  Linux/OSX: \"$HOME/.aws/credentials\"\n" + "  Windows:   \"%USERPROFILE%\\.aws\\credentials\""

	//CmdFlagAwsPolicy is the flag for policy
	CmdFlagAwsPolicy = "policy-arn"

	//CmdFlagAwsPolicyDefault is the default policy arn cli will use
	CmdFlagAwsPolicyDefault = "arn:aws:iam::aws:policy/SecurityAudit"

	//CmdFlagAwsPolicyDescription describes flag policy-arn
	CmdFlagAwsPolicyDescription = "The arn of the policy you'd like to attach for role creation, SecurityAudit policy arn by default"

	//CmdFlagRoleExternalIDDescription is description for flag --external-id
	CmdFlagRoleExternalIDDescription = "The external-id used to assume the provided role"

	//CmdFlagCloudIDDescription flag description
	CmdFlagCloudIDDescription = "Coreo cloud id"

	//InfoCloudAccountDeleted info
	InfoCloudAccountDeleted = "Cloud account was deleted"

	//InfoUsingCloudAccount info
	InfoUsingCloudAccount = "[ OK ] Using Cloud Account ID %s\n"

	//ErrorCloudIDRequired error message
	ErrorCloudIDRequired = "Cloud Account ID is required for this command. Use flag '--cloud-id'\n"

	//CmdFlagLevelLong is flag --severity
	CmdFlagLevelLong = "severity"

	//CmdFlagLevelDescription is description for flag --severity
	CmdFlagLevelDescription = "The severity level you'd like to show in violation results"
)
