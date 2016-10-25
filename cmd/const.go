package cmd


const (
	ACCESS_KEY = "API_KEY"
	SECRET_KEY = "SECRET_KEY"
	TEAM_ID = "TEAM_ID"
	DEFAULT_FOLDER = "/.cloudcoreo"
	DEFAULT_FILE = "profiles.yaml"
	NONE = "None"

	// root command
	CMD_COREO_USE = "cloudcoreo"
	CMD_COREO_SHORT = "A brief description of your application"
	CMD_COREO_LONG = `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`

	// Configure command
	CMD_CONFIG_USE = "configure"
	CMD_CONFIG_SHORT = "create a new configuration"
	CMD_CONFIG_LONG = `Configure  Coreo  CLI  options.`
	CMD_CONFIG_PROMPT_API_KEY = "Enter CloudCoreo api key[%s]: "
	CMD_CONFIG_PROMPT_SECRET_KEY = "Enter CloudCoreo secret key[%s]: "
	CMD_CONFIG_PROMPT_TEAM_ID = "Enter CloudCoreo team ID[%s]: "
)