# CloudCore CLI

### Build
```
go build
```

###  Commands
```
CloudCoreo CLI.

Usage:
  coreo [command]

Available Commands:
  cloud       Manage Coreo Cloud Accounts
  composite   Manage Coreo Composites
  configure   Manage Coreo Configuration
  git-key     Manage Coreo Git keys
  plan        Manage Coreo Plans
  team        Manage Coreo Teams
  token       Manage Coreo Tokens
  version     Print the version number of Coreo CLI

Flags:
      --api-key string      Coreo API Key (default "None")
      --api-secret string   Coreo API Secret (default "None")
      --config string       Config file (default is $HOME/.cloudcoreo/profiles.yaml)
      --json                Output in json format
      --profile string      Coreo profile to use. (default "default")
      --team-id string      Coreo team id (default "None")
      --verbose             Enable verbose output

Use "coreo [command] --help" for more information about a command.
```