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

// TreeNode is the fundamental element of an org tree
type TreeNode struct {
	Info     *OrgNode
	Parent   *TreeNode
	Children []*TreeNode
}

// OrgNode of account
type OrgNode struct {
	ID         string
	Name       string
	Type       string
	Properties map[string]string
}

// Organization struct
type Organization struct {
	ID            string
	MasterAccount *OrgNode
	Properties    map[string]string
}

type GetOrgTreeInput struct {
	AwsProfile     string
	AwsProfilePath string
	RoleArn        string
}
