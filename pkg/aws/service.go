package aws

import (
	"fmt"

	"github.com/CloudCoreo/cli/client"
	"github.com/CloudCoreo/cli/pkg/command"
)

// Service contains three aws service groups
type Service struct {
	setup *SetupService
	org   *OrgService
	role  *RoleService
}

// NewServiceInput contains the info for creating a new Service
type NewServiceInput struct {
	AwsProfile       string
	AwsProfilePath   string
	RoleArn          string
	Policy           string
	RoleSessionName  string
	Duration         int64
	IgnoreCloudTrail bool
}

// NewService returns a new aws service group
func NewService(input *NewServiceInput) *Service {
	return &Service{
		setup: NewSetupService(input),
		org:   NewOrgService(input),
		role:  NewRoleService(input),
	}
}

// SetupEventStream calls the SetupEventStream function in SetupService
func (s *Service) SetupEventStream(input *client.EventStreamConfig) error {
	return s.setup.SetupEventStream(input)
}

// GetOrgTree calls the GetOrganizationTree function in OrgService
func (s *Service) GetOrgTree() ([]*command.TreeNode, error) {
	return s.org.GetOrganizationTree()
}

// CreateNewRole calls the CreateNewRole function in RoleService
func (s *Service) CreateNewRole(input *client.RoleCreationInfo) (arn string, externalID string, err error) {
	return s.role.CreateNewRole(input)
}

// DeleteRole calls the DeleteRole function in RoleService
func (s *Service) DeleteRole(roleName, policyArn string) {
	err := s.role.DeleteRole(roleName, policyArn)
	if err != nil {
		fmt.Println("Failed to delete role" + roleName + ", " + err.Error())
	} else {
		fmt.Println("Deleted role successfully!")
	}
}
