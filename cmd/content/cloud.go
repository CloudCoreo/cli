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

	//CmdCloudTestShort short description
	CmdCloudTestShort = "test role"

	//CmdCloudTestLong long description
	CmdCloudTestLong = "test whether role is valid"

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
	//CmdCloudUpdateShort short description
	CmdCloudUpdateShort = "Update cloud account info"

	//CmdCloudUpdateLong long description
	CmdCloudUpdateLong = "Update cloud account info"

	//CmdCloudScanShort short description
	CmdCloudScanShort = "Scan your root account and create skeletons"

	//CmdCloudScanLong long description
	CmdCloudScanLong = "Scan your root account, get organization and create skeletons for each account"

	//CmdCloudAddExample ...
	CmdCloudAddExample = `  vss cloud add --name YOUR_NEW_ACCOUNT_NAME --role NAME_FOR_NEW_ROLE
  vss cloud add --name YOUR_NEW_ACCOUNT_NAME --arn YOUR_ROLE_ARN --external-id EXTERNAL_ID_OF_YOUR_ROLE`

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

	//CmdFlagRoleSessionName is the name of session for cloud scan command
	CmdFlagRoleSessionName = "role-session"

	//CmdFlagRoleSessionNameDescription is the description of flag roleSessionName
	CmdFlagRoleSessionNameDescription = "The session name to assume the role"

	//CmdFlagIgnoreMissingTrails will make CLI skip on current region of which cloudTrail is not enabled and go on.
	CmdFlagIgnoreMissingTrails = "ignore-missing-trails"

	//CmdFlagIgnoreMissingTrailsDescription describes the usage of CmdFlagIgnoreMissingTrails flag
	CmdFlagIgnoreMissingTrailsDescription = "CLI will continue on event steam setup even if CloudTrail is not enabled in all regions"

	//CmdFlagDuration is the duration of session keys for cloud scan command
	CmdFlagDuration = "duartion"

	//CmdFlagDurationDescription describes the flag duration
	CmdFlagDurationDescription = "The duration for session in seconds"

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

	//CmdFlagIsDraft will add a draft account
	CmdFlagIsDraft = "draft"

	//CmdFlagIsDraftDescription describes the usage of draft flag
	CmdFlagIsDraftDescription = "Create a draft"

	//CmdFlagUserName is the owner username when adding a cloud account, optional
	CmdFlagUserName = "username"

	//CmdFlagUserNameDescription describes the username flag
	CmdFlagUserNameDescription = "User name of account owner"

	//CmdFlagEnvironmentShort is the short flag for environment flag
	CmdFlagEnvironmentShort = "e"

	//CmdFlagEnvironmentLong is the long flag for environment flag
	CmdFlagEnvironmentLong = "environment"

	//CmdFlagEnvironmentDescription describes the usage of environment flag
	CmdFlagEnvironmentDescription = "Environment label for cloud account, four options available: Production, Staging, Development, Test"

	//CmdFlagProvider is the cloud account provider
	CmdFlagProvider = "provider"

	//CmdFlagProviderDescription describes the usage of provider flag
	CmdFlagProviderDescription = "Your cloud account provider, AWS or Azure"

	//CmdFlagEmail is the owner email when adding cloud accounts, optional
	CmdFlagEmail = "email"

	//CmdFlagEmailDescription describes the usage of email flag
	CmdFlagEmailDescription = "Email address of account owner"

	//CmdFlagAwsAssumeRolePolicy is the flag for assume role policy
	CmdFlagAwsAssumeRolePolicy = "policy"

	//CmdFlagAwsAssumeRolePolicyDescription describes the usage of flag policy
	CmdFlagAwsAssumeRolePolicyDescription = "The policy you'd like to use to assume that role"

	//CmdFlagAwsAssumeRolePolicyDefault is the default assume role policy value
	CmdFlagAwsAssumeRolePolicyDefault = "{\"Version\":\"2012-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Action\":\"organizations:*\",\"Resource\":\"*\"}]}"

	//CmdFlagRoleExternalIDDescription is description for flag --external-id
	CmdFlagRoleExternalIDDescription = "The external-id used to assume the provided role"

	//CmdFlagCloudIDDescription flag description
	CmdFlagCloudIDDescription = "Vss cloud id"

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

	//CmdFlagDeleteRole is a flag to delete the role while deleting a cloud account
	CmdFlagDeleteRole = "role"

	//CmdFlagDeleteRoleDescription describes flag --role
	CmdFLagDeleteRoleDescription = "Use this flag to delete the role while deleting a cloud account"

	//CmdFlagKeyValue is the flag for key value
	CmdFlagKeyValue = "key-value"

	//CmdFlagKeyValueDescription is the description for flag --key-value
	CmdFlagKeyValueDescription = "Key required for adding Azure cloud accounts"

	//CmdFlagApplicationID is the flag for application ID
	CmdFlagApplicationID = "application-id"

	//CmdFlagApplicationIDDescription is the description for flag --application-id
	CmdFlagApplicationIDDescription = "Application ID required for adding Azure cloud accounts"

	//CmdFlagDirectoryID is the flag for directory ID
	CmdFlagDirectoryID = "directory-id"

	//CmdFlagDirectoryIDDescription is the description for flag --directory-id
	CmdFlagDirectoryIDDescription = "Directory ID required for adding Azure cloud accounts"

	//CmdFlagSubscriptionID is the flag for subscription ID
	CmdFlagSubscriptionID = "subscription-id"

	//CmdFlagSubscriptionID is the description for flag --subscription-id
	CmdFlagSubscriptionIDDescription = "Subscription ID required for adding Azure cloud accounts"
)
