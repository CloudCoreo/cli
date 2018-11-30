package aws

import (
	"fmt"

	"github.com/CloudCoreo/cli/client"
	"github.com/CloudCoreo/cli/pkg/command"
)

type Service struct {
	setup *SetupService
	org   *OrgService
	role  *RoleService
}

type NewServiceInput struct {
	AwsProfile      string
	AwsProfilePath  string
	RoleArn         string
	Policy          string
	RoleSessionName string
	Duration        int64
}

func NewService(awsProfile, awsProfilePath, roleArn string) *Service {
	return &Service{
		setup: NewSetupService(awsProfile, awsProfilePath),
		org:   NewOrgService(awsProfile, awsProfilePath),
		role:  NewRoleService(awsProfile, awsProfilePath),
	}
}

func (s *Service) SetupEventStream(input *client.EventStreamConfig) error {
	return s.setup.SetupEventStream(input)
}

func (s *Service) GetOrgTree(input *command.GetOrgTreeInput) ([]*command.TreeNode, error) {
	return s.org.GetOrganizationTree()
}

func (s *Service) CreateNewRole(input *client.RoleCreationInfo) (arn string, externalID string, err error) {
	return s.role.CreateNewRole(input)
}

func (s *Service) DeleteRole(roleName, policyArn string) {
	err := s.role.DeleteRole(roleName, policyArn)
	if err != nil {
		fmt.Println("Failed to delete role" + roleName + ", " + err.Error())
	}
}
