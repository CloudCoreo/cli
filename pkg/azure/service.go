package azure

import (
	"github.com/CloudCoreo/cli/client"
	"github.com/CloudCoreo/cli/pkg/command"
)

type NewServiceInput struct {
	AuthFile string
	Region   string
}

type Service struct {
	setup  *SetupService
	remove *RemoveService
}

// NewService returns a new Azure service group
func NewService(input *NewServiceInput) *Service {
	return &Service{
		setup:  NewSetupService(input),
		remove: NewRemoveService(input),
	}
}

// SetupEventStream calls the SetupEventStream function in SetupService
func (s *Service) SetupEventStream(input *client.EventStreamConfig) error {
	return s.setup.SetupEventStream(input)
}

// GetOrgTree calls the GetOrganizationTree function in OrgService
func (s *Service) GetOrgTree() ([]*command.TreeNode, error) {
	return nil, nil
}

// CreateNewRole calls the CreateNewRole function in RoleService
func (s *Service) CreateNewRole(input *client.RoleCreationInfo) (arn string, externalID string, err error) {
	return "", "", nil
}

// DeleteRole calls the DeleteRole function in RoleService
func (s *Service) DeleteRole(roleName string) {

}

//RemoveEventStream perform the same function as event stream removal script
func (s *Service) RemoveEventStream(input *client.EventRemoveConfig) error {
	return s.remove.RemoveEventStream(input)
}
