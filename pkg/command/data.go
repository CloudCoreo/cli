package command

import (
	"github.com/CloudCoreo/cli/client"
)

type SetupEventStreamInput struct {
	AwsProfile     string
	AwsProfilePath string
	Config         *client.EventStreamConfig
}
