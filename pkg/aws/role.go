package aws

import (
	"github.com/CloudCoreo/cli/client"
	"github.com/pkg/errors"

	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

// RoleService interacts with aws role
type RoleService struct {
	AwsProfilePath string
	AwsProfile     string
}

// NewRoleService returns a new RoleService
func NewRoleService(input *NewServiceInput) *RoleService {
	return &RoleService{
		AwsProfile:     input.AwsProfile,
		AwsProfilePath: input.AwsProfilePath,
	}
}

func (c *RoleService) createAssumeRolePolicyDocument(awsAccount string, externalID string) string {
	return `{
	"Version": "2012-10-17",
	"Statement": [
		{
			"Effect": "Allow",
			"Principal": {
				"AWS": "arn:aws:iam::` + awsAccount + `:root"
			},
			"Action": "sts:AssumeRole",
			"Condition": {
				"StringEquals": {
					"sts:ExternalId": "` + externalID + `"
				}
			}
		}
	]
}`
}

// CreateNewRole created a role with specified policy attached
func (c *RoleService) CreateNewRole(input *client.RoleCreationInfo) (arn string, externalID string, err error) {
	sess, err := c.newSession()
	svc := iam.New(sess)

	// Create a new session for iam
	result, err := c.createNewAwsRole(input.AwsAccount, input.ExternalID, input.RoleName, svc)
	if err != nil {
		return "", "", err
	}
	roleArn := result.Role.Arn
	_, err = c.attachRolePolicy(svc, input.Policy, input.RoleName)
	if err != nil {
		return "", "", err
	}

	return *roleArn, input.ExternalID, nil
}

func (c *RoleService) createNewAwsRole(awsAccount, externalID, roleName string, svc *iam.IAM) (*iam.CreateRoleOutput, error) {
	input := &iam.CreateRoleInput{
		AssumeRolePolicyDocument: aws.String(c.createAssumeRolePolicyDocument(awsAccount, externalID)),
		Path:     aws.String("/"),
		RoleName: aws.String(roleName),
	}

	result, err := svc.CreateRole(input)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *RoleService) attachRolePolicy(svc *iam.IAM, policyArn, roleName string) (*iam.AttachRolePolicyOutput, error) {
	input := &iam.AttachRolePolicyInput{
		PolicyArn: aws.String(policyArn),
		RoleName:  aws.String(roleName),
	}

	result, err := svc.AttachRolePolicy(input)
	return result, err
}

func (c *RoleService) newSession() (*session.Session, error) {
	var sess *session.Session
	var err error
	if c.AwsProfile != "" {
		sess, err = session.NewSession(&aws.Config{Credentials: credentials.NewSharedCredentials(c.AwsProfilePath, c.AwsProfile)})
		if err != nil {
			return nil, err
		}
	} else {
		sess, err = session.NewSession()
		if err != nil {
			return nil, err
		}
	}
	return sess, nil
}

func (c *RoleService) DetachPolicy(roleName, policyArn string) error {
	sess, err := c.newSession()

	if err != nil {
		return err
	}

	svc := iam.New(sess)
	detachPolicyInput := &iam.DetachRolePolicyInput{
		PolicyArn: aws.String(policyArn),
		RoleName:  aws.String(roleName),
	}
	_, err = svc.DetachRolePolicy(detachPolicyInput)
	if err != nil {
		return errors.New("Detach role policy " + policyArn + "for " + roleName + " failed, " + err.Error())
	}

	return nil
}

// DeleteRole will remove the role created before if the cloud account add fails
func (c *RoleService) DeleteRole(roleName string) error {
	sess, err := c.newSession()

	if err != nil {
		return err
	}

	svc := iam.New(sess)

	policies, err := c.getManagedRolePolicies(svc, roleName)
	for _, policy := range policies {
		policyArn := *(policy.PolicyArn)
		detachPolicyInput := &iam.DetachRolePolicyInput{
			PolicyArn: aws.String(policyArn),
			RoleName:  aws.String(roleName),
		}
		_, err = svc.DetachRolePolicy(detachPolicyInput)
		if err != nil {
			return errors.New("Detach role policy " + policyArn + "for " + roleName + " failed, " + err.Error())
		}
	}
	/*
		detachPolicyInput := &iam.DetachRolePolicyInput{
			PolicyArn: aws.String(policyArn),
			RoleName:  aws.String(roleName),
		}
		_, err = svc.DetachRolePolicy(detachPolicyInput)
		if err != nil {
			return errors.New("Detach role policy " + policyArn + "for " + roleName + " failed, " + err.Error())
		}*/

	deleteRoleInput := &iam.DeleteRoleInput{
		RoleName: aws.String(roleName),
	}
	_, err = svc.DeleteRole(deleteRoleInput)
	if err != nil {
		return errors.New("Delete role " + roleName + " failed, " + err.Error())
	}
	return nil
}

func (c *RoleService) checkRolePolicy(roleName, policy string) (bool, error) {
	sess, err := c.newSession()

	if err != nil {
		return false, err
	}

	svc := iam.New(sess)
	input := &iam.ListAttachedRolePoliciesInput{}
	input.SetRoleName(roleName)
	res, err := svc.ListAttachedRolePolicies(input)
	if err != nil {
		return false, err
	}
	for i := range res.AttachedPolicies {
		if *res.AttachedPolicies[i].PolicyName == policy {
			return true, nil
		}
	}
	return false, nil
}

func (c *RoleService) getManagedRolePolicies(svc *iam.IAM, roleName string) ([]*iam.AttachedPolicy, error) {
	res := make([]*iam.AttachedPolicy, 0)

	input := &iam.ListAttachedRolePoliciesInput{
		RoleName: &roleName,
	}
	err := svc.ListAttachedRolePoliciesPages(input, func(output *iam.ListAttachedRolePoliciesOutput, last bool) bool {
		res = append(res, output.AttachedPolicies...)
		return true
	})
	return res, err
}
