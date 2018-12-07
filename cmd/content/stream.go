package content

//CmdEventSetupUse is a command to setup event stream
const CmdEventSetupUse = "setup"

//CmdEventSetupShort is the short version description for coreo event setup command
const CmdEventSetupShort = "Setup event stream"

//CmdEventSetupLong is the long version description for coreo event setup command
const CmdEventSetupLong = "Run this command to setup event stream. " +
	"It will create a CloudFormation stack with an event rule and SNS topic. " +
	"You will need to run this script for each cloud account. " +
	"Make sure your aws credentials have been configured before run this command."

//CmdEventUse is command for event stream
const CmdEventUse = "event"

//CmdEventShort is the short version description for coreo event command
const CmdEventShort = "Manage event stream"

//CmdEventLong is the long version description for coreo event command
const CmdEventLong = "Manage event stream"

//CmdEventSetupExample is the use case for command event setup
const CmdEventSetupExample = `  coreo event setup
  coreo event setup --aws-profile YOUR_AWS_PROFILE`
