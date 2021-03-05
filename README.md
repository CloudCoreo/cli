# CloudHealth Secure State CLI

[![Build Status](https://travis-ci.org/CloudCoreo/cli.svg?branch=master)](https://travis-ci.org/CloudCoreo/cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/CloudCoreo/cli)](https://goreportcard.com/report/github.com/CloudCoreo/cli)

CLI is a tool for managing CloudHealth Secure State resources. 

Use CLI to...

- Add/remove cloud accounts and API tokens
- Event stream setup and removal
- Get violation results

**NOTE: Secure State recently changed our name to CloudHealth Secure State. The CLI will still include references to vss.

## Install

**DISCLAIMER: These are PRE-RELEASE binaries -- use at your own peril for now**

### OSX

Download `vss` from [https://github.com/CloudCoreo/cli/releases/download/v0.0.51/vss_darwin_amd64](https://github.com/CloudCoreo/cli/releases/download/v0.0.51/vss_darwin_amd64)

```sh
 mkdir vss && cd vss
 wget -q -O vss https://github.com/CloudCoreo/cli/releases/download/v0.0.51/vss_darwin_amd64
 chmod +x vss
 export PATH=$PATH:${PWD}   # Add current dir where vss has been downloaded to
 vss
```

### Linux

Download `vss` from [https://github.com/CloudCoreo/cli/releases/download/v0.0.51/vss_linux_amd64](https://github.com/CloudCoreo/cli/releases/download/v0.0.51/vss_linux_amd64)

```sh
 mkdir vss && cd vss
 wget -q -O vss https://github.com/CloudCoreo/cli/releases/download/v0.0.51/vss_linux_amd64
 chmod +x vss
 export PATH=$PATH:${PWD}   # Add current dir where vss has been downloaded to
 vss
```

### Windows

Download `vss.exe` from [https://github.com/CloudCoreo/cli/releases/download/v0.0.51/vss_windows_amd64.exe](https://github.com/CloudCoreo/cli/releases/download/v0.0.51/vss_windows_amd64.exe)

```
C:\Users\Username\Downloads> rename vss_windows_amd64.exe vss.exe
C:\Users\Username\Downloads> vss.exe
```

### Building from source

Build instructions are as follows (see [install golang](https://docs.minio.io/docs/how-to-install-golang) for setting up a working golang environment):

```sh
 mkdir -p $GOPATH/src/github.com/CloudCoreo
 cd $GOPATH/src/github.com/CloudCoreo
 git clone https://github.com/CloudCoreo/cli.git
 go get -d github.com/CloudCoreo/cli
 cd $GOPATH/src/github.com/CloudCoreo/cli/cmd
 go build -o $GOPATH/bin/vss
 vss
```
## Getting started
Get your access keys on [VMware CSP User Portal](https://console.cloud.vmware.com/csp/gateway/portal/#/user/tokens).

You use API tokens to authenticate yourself when you make authorized API connections. An API token authorizes access per organization.

You can generate more than one API token. A token is valid for six months, after which time you must regenerate it if you want to continue using APIs that rely on a token. If you feel the token has been compromised, you can revoke the token to prevent unauthorized access. You generate a new token to renew authorization.

Procedure
1. On the VMware Cloud Services toolbar, click your user name and select My Account > API Tokens.
2. Click New Token.
3. Click Copy to Clipboard.
4. Paste the token in a safe place so you can retrieve it to use later on.


You may need to configure your access key the first time using CLI but you can also skip this step and pass these to CLI using flags `--api-key`. You may set up configuration using:
	`vss  configure`
	
And then type your access key information. You may check you current configuration settings using
	`vss configure list`

Team id concept is deprecated in the latest CLI release and is not required anymore.
## Usage
```sh
vss <command> [--help] [--verbose] [--json] [<args>]
```

The most commonly used VSS commands are:


|Command         |Usage      | Sub-commands|
| --------   | :-------------:| :-------------:|
|cloud     | Manage your cloud accounts                    | add, delete, list, scan, show, update, test|
|configure | Configure CLI options. You may also view your current configuration using 'list' subcommand| list|
|team      | Manage your team(Deprecated, this info is not required anymore)                              | add, list, show|
|result    | Get violation results (Deprecated, please follow the link to swagger API doc 'https://api.securestate.vmware.com')  | rule, object|
|token     | Manage your api tokens(Deprecated, please manage your token through CSP portal)                        | delete, list, show|
|completion| Generate bash autocompletions script|
|event     | Manage event stream                           | setup|
|help      | Help about any command|
|version   | Print the version number of the Secure State CLI|
-------------      
 
## Configurable variables
|Variable | Option | Environment Variable | Description |
| ------ | ------ | :--------:| :-------- |
|api-key | --api-key| |VSS API Token, will read api-key in configure file by default| 
|endpoint| --endpoint |$VSS_API_ENDPOINT| VSS API endpoint, default https://app.securestate.vmware.com/api |
|help    | --help, -h| | Get user manual for command
|home    | --home | $VSS_HOME | Location of your VSS config. Overrides $VSS_HOME.
|json    |--json | | Output in json format
|profile | --profile | $VSS_PROFILE | VSS profile to use. Overrides $VSS_PROFILE, default "default" |
|team-id | --team-id | | Secure State team id. This flag is deprecated in the latest CLI release and not required anymore|
|verbose | --verbose | | Enable verbose output

The values passing by flags will override environment variables.  
Flags for specific commands are listed in Docs section.

## Example
You may use CLI to do scriptable onboarding with two commands:
```sh
 vss cloud add --name YOUR_NEW_ACCOUNT_NAME --role NAME_FOR_NEW_ROLE [--aws-profile PROFILE_NAME] [–aws-profile-path PROFILE_PATH] [--policy-arn YOUR_POLICY_ARN]  
 vss event setup --account-id YOUR_ACCOUNT_ID [--aws-profile PROFILE_NAME] [--aws-profile-path PROFILE_PATH] 
```
team-id flag is not required from CLI release v0.0.51
## Docs

Get started with [VSS commands](docs/vss/vss.md), setup for [VSS bash completion](docs/bash-completion.md)

#### cloud
Manage Cloud Accounts
* add
    * Usage
        * `vss cloud add --name YOUR_NEW_ACCOUNT_NAME --role NAME_FOR_NEW_ROLE [flags]`  
        * `vss cloud add --name YOUR_NEW_ACCOUNT_NAME --arn YOUR_ROLE_ARN --external-id EXTERNAL_ID_OF_YOUR_ROLE [flags]`
    * Flags
    
        |Variable | Option | Description |
        | ------ | ------ | :-------- |
        | arn | --arn |  The arn of the role to connect |
        | name | --name| The name of the new cloud account you want to add, this flag is required |
        | role | --role | The name of the role you want to create
        | policy arn| --policy-arn | The arn of the policy you'd like to attach for role creation, SecurityAudit policy arn by default|
        | external id| --external-id | The external id used to assume provided role|
        | aws profile | --aws-profile |  Aws shared credential file. If empty default provider chain will be used to look for credentials with the following order. <br> <br> 1. Environment variables.<br>2. Shared credentials file. <br>3. If your application is running on an Amazon EC2 instance, IAM role for Amazon EC2.
        | aws profile path| --aws-profile-path| The file path of aws profile. If empty will look for AWS_SHARED_CREDENTIALS_FILE env variable. If the env value is empty will default to current user's home directory. <br> <br> Linux/OSX: &nbsp; "$HOME/.aws/credentials"<br> Windows: &nbsp;&nbsp;&nbsp; "%USERPROFILE%\.aws\credentials"
        | draft| --draft| Will add a draft account with this flag|
        | Environment| --env| Environment label for the cloud account to add, must be one of these: Production, Staging, Development, Test"|
        | email|--email|The email address of account owner|
        | username|--username| The username of account owner|
        | provider|--provider| Cloud provider type, either AWS or Azure, AWS by default|
        | application id |--application-id| Application ID is required for adding Azure cloud accounts|
        | key |--key-value| Key is required for adding Azure cloud accounts
        | subscription id |--subscription-id| Subscription ID is required for adding Azure cloud accounts |
        | directory id |--directory-id| Directory ID is required for adding Azure Cloud Accounts |
        | cloud account tags| --tags| Cloud account tags|
        
    * You need to either use your own role or let CLI create one for you. 
        * To use your own role, you need to pass the role arn and external id to CLI. 
        * To make CLI create one for you, you need to pass the role name to CLI
    * Examples:
        * `vss cloud add --name YOUR_NEW_ACCOUNT_NAME --provider AWS --role NAME_FOR_NEW_ROLE --aws-profile AWS_PROFILE --tags "key1:value1|key2:value2"`
        * `vss cloud add --name YOUR_NEW_ACCOUNT_NAME --provider Azure --application-id AZURE_APPLICATION_ID --key-value KEY_VALUE --subscription-id SUBSCRIPTION_ID --directory-id DIRECTORY_ID`
        
* delete
    * Usage
        * `vss cloud delete --account-id YOUR_ACCOUNT_ID --provider PROVIDER [flags]`
    * Flags  
    
        |Variable | Option | Description |
        | ------ | ------ | :-------- |
        | account id| --account-id| Cloud account id of which account you'd like to delete, this flag is required|
        | provider|--provider| Cloud provider type, you may use AWS, Azure or GCP, AWS by default|
        | aws profile | --aws-profile |  Aws shared credential file. If empty default provider chain will be used to look for credentials with the following order. <br> <br> 1. Environment variables.<br>2. Shared credentials file. <br>3. If your application is running on an Amazon EC2 instance, IAM role for Amazon EC2.
        | aws profile path| --aws-profile-path| The file path of aws profile. If empty will look for AWS_SHARED_CREDENTIALS_FILE env variable. If the env value is empty will default to current user's home directory. <br> <br> Linux/OSX: &nbsp; "$HOME/.aws/credentials"<br> Windows: &nbsp;&nbsp;&nbsp; "%USERPROFILE%\.aws\credentials"
* list
    * Usage
        *  `vss cloud list [flags]`

* show
    * Usage
        * `vss cloud show --account-id YOUR_ACCOUNT_ID --provider PROVIDER [flags]`
    * Flags
    
        |Variable | Option | Description |
        | ------ | ------ | :-------- |
        | account id| --account-id| Cloud account id of which account you'd like to delete, this flag is required|
        | provider|--provider| Cloud provider type, you may use AWS, Azure or GCP, AWS by default|
* update
    * Usage
        * `vss cloud update --account-id YOUR_ACCOUNT_ID --provider PROVIDER [flags]`
    * Flags 
    
        |Variable | Option | Description |
        | ------ | ------ | :-------- |
        | arn | --arn |  The arn of the role to connect |
        | name | --name| The name of the new cloud account you want to add, this flag is required |
        | role | --role | The name of the role you want to create
        | policy arn| --policy-arn | The arn of the policy you'd like to attach for role creation, SecurityAudit policy arn by default|
        | external id| --external-id | The external id used to assume provided role|
        | aws profile | --aws-profile |  Aws shared credential file. If empty default provider chain will be used to look for credentials with the following order. <br> <br> 1. Environment variables.<br>2. Shared credentials file. <br>3. If your application is running on an Amazon EC2 instance, IAM role for Amazon EC2.
        |aws profile path| --aws-profile-path| The file path of aws profile. If empty will look for AWS_SHARED_CREDENTIALS_FILE env variable. If the env value is empty will default to current user's home directory. <br> <br> Linux/OSX: &nbsp; "$HOME/.aws/credentials"<br> Windows: &nbsp;&nbsp;&nbsp; "%USERPROFILE%\.aws\credentials"
        |draft| --draft| Will update the account with draft status|
        |Environment| --env| Environment label for the cloud account to add, must be one of these: Production, Staging, Development, Test"|
        |email|--email|The email address of account owner|
        |username|--username| The username of account owner|
        | account id| --account-id| Cloud account id of which account you'd like to delete, this flag is required|
        | provider|--provider| Cloud provider type, you may use AWS, Azure or GCP, AWS by default|
        | cloud account tags| --tags| Cloud account tags|
    * For role update, you may either provide your own role or let CLI create one
    * You may need to use --draft flag if you still want to keep it as draft status, otherwise VSS CLI will switch it to non-draft status
        
* test
    * Usage
        * `vss cloud test --account-id YOUR_ACCOUNT_ID --provider PROVIDER`
        * Flags
        
            |Variable | Option | Description |
            | ------ | ------ | :-------- |
            | account id| --account-id| Cloud account id of which account you'd like to delete, this flag is required|
            | provider|--provider| Cloud provider type, you may use AWS, Azure or GCP, AWS by default|
            
#### configure
Configure CLI options
* Usage
    * `vss configure [flags]` &nbsp; :configure CLI options
    * `vss configure list` &nbsp; : list current configuration
* Examples
    * `vss configure`
    * `vss configure --api-key VSS_API_TOKEN`
    * `vss configure list`
    
#### team
Manage Teams(These commands are deprecated from CLI version v0.0.51)
* add
    * Usage
        * `vss team add -n YOUR_NEW_TEAM_NAME -d YOUR_TEAM_DESCRIPTION [flags]`
    * Flags
        
        |Variable | Option | Description |
        | ------ | ------ | :-------- |
        |name| -n, --name | Provide team name |
        |description| -d, --description | Provide team description|
    * Flag name and description are required for team add command
* list  
    Get all the teams under user's account
    * Usage
        * `vss team list [flags]`
* show  
    Show info of one team
    * Usage 
        * `vss team show [flags]`        

#### result
Show violation results (Deprecated, please follow the link to swagger API doc 'https://api.securestate.vmware.com'
* object
    * Usage
        * `vss result object [flags]`
* rule
    * Usage
        * `vss result rule [flags]`
        
#### token
Manage API Tokens(These commands are deprecated from CLI version v0.0.51, please use [CSP portal](https://console.cloud.vmware.com/csp/gateway/portal/#/user/tokens)) to manage your token)
* delete
    * Usage
        * `vss token delete --token-id YOUR_TOKEN_ID [flags]`             
* list
    * Usage 
        * `vss token list [flags]`
* show
    * Usage
        * `vss token show --token-id YOUR_TOKEN_ID [flags]`
        
#### completion
Generate bash auto-completions script
* Usage
    * `vss completion [flags]`
    
#### event
Manage event stream
* setup
    * Usage 
        * `vss event setup --account-id YOUR_ACCOUNT_ID --provider PROVIDER [flags]`
    * Flags
    
        |Variable | Option | Description |
        | ------ | ------ | :-------- |
        | aws profile | --aws-profile |  Aws shared credential file. If empty default provider chain will be used to look for credentials with the following order. <br> <br> 1. Environment variables.<br>2. Shared credentials file. <br>3. If your application is running on an Amazon EC2 instance, IAM role for Amazon EC2.
        |aws profile path| --aws-profile-path| The file path of aws profile. If empty will look for AWS_SHARED_CREDENTIALS_FILE env variable. If the env value is empty will default to current user's home directory. <br> <br> Linux/OSX: &nbsp; "$HOME/.aws/credentials"<br> Windows: &nbsp;&nbsp;&nbsp; "%USERPROFILE%\.aws\credentials"
        | account id| --account-id| Cloud account id of which account you'd like to delete, this flag is required|
        | provider|--provider| Cloud provider type, you may use AWS, Azure or GCP, AWS by default|
        |ignore-missing-trails|--ignore-missing-trails| With this flag, CLI will skip regions of which CloudTrail in not enables and continue on other regions.|

* remove
    * Usage 
        * `vss event remove --account-id YOUR_ACCOUNT_ID --provider PROVIDER [flags]`
    * Flags
        
        |Variable | Option | Description |
        | ------ | ------ | :-------- |
        | aws profile | --aws-profile |  Aws shared credential file. If empty default provider chain will be used to look for credentials with the following order. <br> <br> 1. Environment variables.<br>2. Shared credentials file. <br>3. If your application is running on an Amazon EC2 instance, IAM role for Amazon EC2.
        |aws profile path| --aws-profile-path| The file path of aws profile. If empty will look for AWS_SHARED_CREDENTIALS_FILE env variable. If the env value is empty will default to current user's home directory. <br> <br> Linux/OSX: &nbsp; "$HOME/.aws/credentials"<br> Windows: &nbsp;&nbsp;&nbsp; "%USERPROFILE%\.aws\credentials"
        | account id| --account-id| Cloud account id of which account you'd like to delete, this flag is required|
        | provider|--provider| Cloud provider type, you may use AWS, Azure or GCP, AWS by default|
        
#### help
Help about any command
* Usage   
    * `vss help`
#### version
Print the version number of Secure State CLI
* Usage 
    * `vss version`
## Community, discussion, contribution, and support

GitHub's Issue tracker is to be used for managing bug reports, feature requests, and general items that need to be addressed or discussed.

From a non-developer standpoint, Bug Reports or Feature Requests should be opened being as descriptive as possible.

### Code of conduct

Participation in the CloudCoreo community is governed by the [Coreo Code of Conduct](code-of-conduct.md).
