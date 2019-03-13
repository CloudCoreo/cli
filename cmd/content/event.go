package content

//CmdEventSetupUse is a command to setup event stream
const CmdEventSetupUse = "setup"

//CmdEventSetupShort is the short version description for vss event setup command
const CmdEventSetupShort = "Setup event stream"

//CmdEventSetupLong is the long version description for vss event setup command
const CmdEventSetupLong = "Run this command to setup event stream. " +
	"It will create a CloudFormation stack with an event rule and SNS topic. " +
	"You will need to run this script for each cloud account. " +
	"Make sure your aws credentials have been configured before run this command."

//CmdEventUse is command for event stream
const CmdEventUse = "event"

//CmdEventShort is the short version description for vss event command
const CmdEventShort = "Manage event stream"

//CmdEventLong is the long version description for vss event command
const CmdEventLong = "Manage event stream"

//CmdEventSetupExample is the use case for command event setup
const CmdEventSetupExample = `  vss event setup
  vss event setup --aws-profile YOUR_AWS_PROFILE --cloud-id YOUR_CLOUD_ID`

//CmdEventRemoveUse is the command name for command event remove
const CmdEventRemoveUse = "remove"

//CmdEventRemoveShort is the short version description for vss event remove command
const CmdEventRemoveShort = "Remove event stream"

//CmdEventRemoveLong is the long version description for vss event remove command
const CmdEventRemoveLong = "Run this command to remove event stream." +
	"You will need to run this script for each cloud account." +
	"Make sure your aws credentials have been configured before run this command."

//CmdEventRemoveExample is the use case for command event remove
const CmdEventRemoveExample = `vss event remove
  vss event remove --aws-profile YOUR_AWS_PROFILE --cloud-id YOUR_CLOUD_ID`

const CmdEventAuthFile = "auth-file"

const CmdEventAuthFileDescription = "auth file for azure authentication"

const CmdEventRegion = "region"

const CmdEventRegionDescription = "The region in which you'd like to create Azure resource group in"
