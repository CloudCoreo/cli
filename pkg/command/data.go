package command

import (
	"github.com/CloudCoreo/cli/client"
)

//SetupEventStreamInput is the input for event stream setup
type SetupEventStreamInput struct {
	AwsProfile     string
	AwsProfilePath string
	Config         *client.EventStreamConfig
}
