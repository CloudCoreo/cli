package aws

import (
	"fmt"

	"github.com/CloudCoreo/cli/client"
)

// Service contains three aws service groups
type Service struct {
	setup  *SetupService
	role   *RoleService
	remove *RemoveService
}

// NewServiceInput contains the info for creating a new Service
type NewServiceInput struct {
	AwsProfile          string
	AwsProfilePath      string
	Policy              string
	RoleSessionName     string
	Duration            int64
	IgnoreMissingTrails bool
}

// NewService returns a new aws service group
func NewService(input *NewServiceInput) *Service {
	return &Service{
		setup:  NewSetupService(input),
		role:   NewRoleService(input),
		remove: NewRemoveService(input),
	}
}

// SetupEventStream calls the SetupEventStream function in SetupService
func (s *Service) SetupEventStream(input *client.EventStreamConfig) error {
	return s.setup.SetupEventStream(input)
}

// CreateNewRole calls the CreateNewRole function in RoleService
func (s *Service) CreateNewRole(input *client.RoleCreationInfo) (arn string, externalID string, err error) {
	return s.role.CreateNewRole(input)
}

// DeleteRole calls the DeleteRole function in RoleService
func (s *Service) DeleteRole(roleName string) {
	err := s.role.DeleteRole(roleName)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Deleted role successfully!")
	}
}

//RemoveEventStream perform the same function as event stream removal script
func (s *Service) RemoveEventStream(input *client.EventRemoveConfig) error {
	return s.remove.RemoveEventStream(input)
}
