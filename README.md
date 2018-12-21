# CloudCoreo CLI

[![Build Status](https://travis-ci.org/CloudCoreo/cli.svg?branch=master)](https://travis-ci.org/CloudCoreo/cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/CloudCoreo/cli)](https://goreportcard.com/report/github.com/CloudCoreo/cli)

CLI is a tool for managing CloudCoreo resources. 

Use CLI to...

- Add/remove teams, cloud accounts, Git keys and API tokens
- Intelligently manage your composites manifest files
- Manage releases of CloudCoreo composites/plans
- Create reproducible plans of your CloudCoreo applications

## Install

**DISCLAIMER: These are PRE-RELEASE binaries -- use at your own peril for now**

### OSX

Download `coreo` from [https://github.com/CloudCoreo/cli/releases/download/v0.0.22/coreo_darwin_amd64](https://github.com/CloudCoreo/cli/releases/download/v0.0.22/coreo_darwin_amd64)

```sh
 mkdir coreo && cd coreo
 wget -q -O coreo https://github.com/CloudCoreo/cli/releases/download/v0.0.22/coreo_darwin_amd64
 chmod +x coreo
 export PATH=$PATH:${PWD}   # Add current dir where coreo has been downloaded to
 coreo
```

### Linux

Download `coreo` from [https://github.com/CloudCoreo/cli/releases/download/v0.0.22/coreo_linux_amd64](https://github.com/CloudCoreo/cli/releases/download/v0.0.22/coreo_linux_amd64)

```sh
 mkdir coreo && cd coreo
 wget -q -O coreo https://github.com/CloudCoreo/cli/releases/download/v0.0.22/coreo_linux_amd64
 chmod +x coreo
 export PATH=$PATH:${PWD}   # Add current dir where coreo has been downloaded to
 coreo
```

### Windows

Download `coreo.exe` from [https://github.com/CloudCoreo/cli/releases/download/v0.0.22/coreo_windows_amd64.exe](https://github.com/CloudCoreo/cli/releases/download/v0.0.22/coreo_windows_amd64.exe)

```
C:\Users\Username\Downloads> rename coreo_windows_amd64.exe coreo.exe
C:\Users\Username\Downloads> coreo.exe
```

### Building from source

Build instructions are as follows (see [install golang](https://docs.minio.io/docs/how-to-install-golang) for setting up a working golang environment):

```sh
 go get -d github.com/CloudCoreo/cli
 cd $GOPATH/src/github.com/CloudCoreo/cli/cmd
 go build -o $GOPATH/bin/coreo
 coreo
```
## Getting started
Get your access keys on [VMware Secure State](https://app.cloudcoreo.com/main/settings/cli).
Go to the “Settings” menu and select the "Command-Line Interface”, then click the “Create Access Key” button. Your default team id is right below the access key id window.  

You may need to configure your access key and default team id the first time using CLI but you can also skip this step and pass these to CLI using flags `--api-key, --api-secret, --team-id`. You may set up configuration using:
	`coreo  configure`
	
And then type you access key information and default team id as well. You may check you current configuration settings using
	`coreo configure list`

## Usage
```sh
coreo <command> [--help] [--verbose] [--json] [<args>]
```

The most commonly used coreo commands are:


|Command         |Usage      | Sub-commands|
| --------   | :-------------:| :-------------:|
|cloud     | Manage your cloud accounts                    | add, delete, list, scan, show|
|configure | Configure CLI options. You may also view your current configuration using 'list' subcommand| list|
|team      | Manage your team                              | add, list, show|
|result    | Get violation results                         | rule, object|
|token     | Manage your api tokens                        | delete, list, show|
|completion| Generate bash autocompletions script|
|event     | Manage event stream                           | setup|
|help      | Help about any command|
|version   | Print the version number of Coreo CLI|
-------------      
 
## Configurable variables
|Variable | Option | Environment Variable | Description |
| ------ | ------ | :--------:| :-------- |
|api-key | --api-key| |Coreo API Key, will read api-key in configure file by default| 
|api-secret |--api-secret | |Coreo API Secret, will read api-secret in configure file by default|
|endpoint| --endpoint |$CC_API_ENDPOINT| Coreo API endpoint, default https://app.cloudcoreo.com/api |
|help    | --help, -h| | Get user manual for command
|home    | --home | $COREO_HOME | Location of your Coreo config. Overrides $COREO_HOME.
|json    |--json | | Output in json format
|profile | --profile | $COREO_PROFILE | Coreo profile to use. Overrides $COREO_PROFILE, default "default" |
|team-id | --team-id | | Coreo team id, will read team-id in configure file by default|
|verbose | --verbose | | Enable verbose output

The values passing by flags will override environment variables.  
Flags for specific commands are listed in Docs section.

## Example
You may use CLI to do scriptable onboarding with two commands:
```sh
 coreo cloud add --name YOUR_NEW_ACCOUNT_NAME --role NAME_FOR_NEW_ROLE [--aws-profile PROFILE_NAME] [–aws-profile-path PROFILE_PATH] [--policy-arn YOUR_POLICY_ARN] [--team-id YOUR_TEAM_ID] 
 coreo event setup --cloud-id YOUR_CLOUD_ID [--aws-profile PROFILE_NAME] [--aws-profile-path PROFILE_PATH] [--team-id YOUR_TEAM_ID]  
```
If the team-id flag is omitted, CLI will use the default team id in configuration. If default team id is not set, an error will be returned.
## Docs

Get started with [Coreo commands](docs/coreo/coreo.md), setup for [Coreo bash completion](docs/bash-completion.md)

#### cloud
Manage Cloud Accounts
* add
    * Usage
        * `coreo cloud add --team-id YOUR_TEAM_ID --name YOUR_NEW_ACCOUNT_NAME --role NAME_FOR_NEW_ROLE [flags]`  
        * `coreo cloud add --team-id YOUR_TEAM_ID --name YOUR_NEW_ACCOUNT_NAME --arn YOUR_ROLE_ARN --external-id EXTERNAL_ID_OF_YOUR_ROLE [flags]`
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
        |draft| --draft| Will add a draft account with this flag|
        |Environment| --env| Environment label for the cloud account to add, must be one of these: Production, Staging, Development, Test"|
        |email|--email|The email address of account owner|
        |username|--username| The username of account owner|
        
    * You need to either use your own role or let CLI create one for you. 
        * To use your own role, you need to pass the role arn and external id to CLI. 
        * To make CLI create one for you, you need to pass the role name to CLI
        
* delete
    * Usage
        * `coreo cloud delete --cloud-id YOUR_CLOUD_ID [flags]`
    * Flags  
    
        |Variable | Option | Description |
        | ------ | ------ | :-------- |
        | cloud id| --cloud-id| VMware Secure State cloud id of which account you'd like to delete, this flag is required|
* list
    * Usage
        *  `coreo cloud list [flags]`
* scan
    * Usage
        * `coreo cloud scan [flags]`
    * Flags
    
        |Variable | Option | Description |
        | ------ | ------ | :-------- |
        | aws profile | --aws-profile |  Aws shared credential file. If empty default provider chain will be used to look for credentials with the following order. <br> <br> 1. Environment variables.<br>2. Shared credentials file. <br>3. If your application is running on an Amazon EC2 instance, IAM role for Amazon EC2.
        |aws profile path| --aws-profile-path| The file path of aws profile. If empty will look for AWS_SHARED_CREDENTIALS_FILE env variable. If the env value is empty will default to current user's home directory. <br> <br> Linux/OSX: &nbsp; "$HOME/.aws/credentials"<br> Windows: &nbsp;&nbsp;&nbsp; "%USERPROFILE%\.aws\credentials"
* show
    * Usage
        * `coreo cloud show --cloud-id YOUR_CLOUD_ID [flags]`
    * Flags
    
        |Variable | Option | Description |
        | ------ | ------ | :-------- |
        | cloud id| --cloud-id| VMware Secure State cloud id of which account you'd like to show information for, this flag is required|
        
#### configure
Configure CLI options
* Usage
    * `coreo configure [flags]` &nbsp; :configure CLI options
    * `coreo configure list` &nbsp; : list current configuration
* Examples
    * `coreo configure`
    * `coreo configure --api-key COREO_API_KEY --api-secret COREO_API_SECRET --team-id COREO_TEAM_ID`
    * `coreo configure list`
    
#### team
Manage Teams
* add
    * Usage
        * `coreo team add -n YOUR_NEW_TEAM_NAME -d YOUR_TEAM_DESCRIPTION [flags]`
    * Flags
        
        |Variable | Option | Description |
        | ------ | ------ | :-------- |
        |name| -n, --name | Provide team name |
        |description| -d, --description | Provide team description|
    * Flag name and description are required for team add command
* list  
    Get all the teams under user's account
    * Usage
        * `coreo team list [flags]`
* show  
    Show info of one team
    * Usage 
        * `coreo team show [flags]`        

#### result
Show violation results
* object
    * Usage
        * `coreo result object [flags]`
    * Flags
        
         |Variable | Option | Description |
         | ------ | ------ | :-------- |
         | cloud id| --cloud-id| VMware Secure State cloud id of which account you'd like to show violation for, this flag is optional|
         | severity | --severity | The severity level you'd like to show in violation results |
    * By default you will get all violation objects under your account, three flag filters are provided: team-id, cloud-id and severity
    * Examples
         * `coreo result object --severity "High|Medium"`    
         * `coreo result object --severity "High|Low"`  
         * `coreo result object --cloud-id YOUR_CLOUD_ID --severity "Low"`
* rule
    * Usage
        * `coreo result rule [flags]`
    * Flags
    
        |Variable | Option | Description |
        | ------ | ------ | :-------- |
        | cloud id| --cloud-id| VMware Secure State cloud id of which account you'd like to show violation for, this flag is optional|
        | severity | --severity | The severity level you'd like to show in violation results |
    * By default you will get all violation rules under your account, three flag filters are provided: team-id, cloud-id and severity
    * Examples
        * `coreo result rule --severity "High|Medium"`
        * `coreo result rule --severity "High|Low"`
        * `coreo result rule --cloud-id YOUR_SECURITY_STATE_CLOUD_ACCOUNT_ID --severity "Low"`
        
#### token
Manage API Tokens
* delete
    * Usage
        * `coreo token delete --token-id YOUR_TOKEN_ID [flags]`
    * Flags
    
        |Variable | Option | Description |
        | :------: | :------: | :--------: |
        | token id| --token-id| Secure State token id, this flag is required|
                
* list
    * Usage 
        * `coreo token list [flags]`
        
* show
    * Usage
        * `coreo token show --token-id YOUR_TOKEN_ID [flags]`
    * Flags
    
        |Variable | Option | Description |
        | :------: | :------: | :--------: |
        | token id| --token-id| Secure State token id, this flag is required|
        
#### completion
Generate bash auto-completions script
* Usage
    * `coreo completion [flags]`
    
#### event
Manage event stream
* setup
    * Usage 
        * `coreo event setup --team-id YOUR_TEAM_ID --cloud-id YOUR_CLOUD_ID [flags]`
    * Flags
    
        |Variable | Option | Description |
        | ------ | ------ | :-------- |
        | aws profile | --aws-profile |  Aws shared credential file. If empty default provider chain will be used to look for credentials with the following order. <br> <br> 1. Environment variables.<br>2. Shared credentials file. <br>3. If your application is running on an Amazon EC2 instance, IAM role for Amazon EC2.
        |aws profile path| --aws-profile-path| The file path of aws profile. If empty will look for AWS_SHARED_CREDENTIALS_FILE env variable. If the env value is empty will default to current user's home directory. <br> <br> Linux/OSX: &nbsp; "$HOME/.aws/credentials"<br> Windows: &nbsp;&nbsp;&nbsp; "%USERPROFILE%\.aws\credentials"
        | cloud id| --cloud-id| VMware Secure State cloud id of which account you'd like to show information for, this flag is required|
        |ignore-missing-trails|--ignore-missing-trails| With this flag, CLI will skip regions of which CloudTrail in not enables and continue on other regions.|

#### help
Help about any command
* Usage   
    * `coreo help`
#### version
Print the version number of Coreo CLI
* Usage 
    * `coreo version`
## Community, discussion, contribution, and support

GitHub's Issue tracker is to be used for managing bug reports, feature requests, and general items that need to be addressed or discussed.

From a non-developer standpoint, Bug Reports or Feature Requests should be opened being as descriptive as possible.

### Code of conduct

Participation in the CloudCoreo community is governed by the [Coreo Code of Conduct](code-of-conduct.md).
