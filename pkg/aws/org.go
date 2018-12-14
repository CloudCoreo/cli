package aws

import (
	"errors"
	"fmt"

	"github.com/CloudCoreo/cli/pkg/command"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/organizations"
	"github.com/aws/aws-sdk-go/service/sts"
)

// OrgService which connects to AWS organizations
type OrgService struct {
	orgService *organizations.Organizations
}

// NewOrgService returns a new orgservice with the credentials provided
func NewOrgService(input *NewServiceInput) (awsService *OrgService) {
	var sess *session.Session
	if input.AwsProfile != "" {
		sess = session.Must(session.NewSession(&aws.Config{Credentials: credentials.NewSharedCredentials(input.AwsProfilePath, input.AwsProfile)}))
	} else {
		sess = session.Must(session.NewSession())
	}

	if input.RoleArn != "" {
		stsSvc := sts.New(sess)
		stsInput := &sts.AssumeRoleInput{
			DurationSeconds: aws.Int64(input.Duration),
			Policy:          aws.String(input.Policy),
			RoleArn:         aws.String(input.RoleArn),
			RoleSessionName: aws.String(input.RoleSessionName),
		}

		sresult, serr := stsSvc.AssumeRole(stsInput)
		if serr != nil {
			fmt.Println("Unable to assume role" + serr.Error())
		}

		res, err := NewOrgServiceWithCreds(sresult.Credentials)
		if err != nil {
			fmt.Println("Unable to initialize AWS Service -", err)
			return nil
		}

		return res
	}

	svc := organizations.New(sess)
	return &OrgService{
		orgService: svc,
	}
}

// NewOrgServiceWithCreds initializes a connection to the AWS organization service
func NewOrgServiceWithCreds(creds *sts.Credentials) (awsService *OrgService, err error) {
	sess := session.Must(session.NewSession())

	// sess := session.Must(session.NewSessionWithOptions(session.Options{

	// 	SharedConfigState: session.SharedConfigEnable,
	// }))
	provider := NewAssumeRoleCredentialsProvider(creds)
	svc := organizations.New(sess, &aws.Config{Credentials: credentials.NewCredentials(provider)})
	return &OrgService{svc}, nil
}

// DescribeOrganization provides the details of an organization
func (svc *OrgService) DescribeOrganization() (response *command.Organization, err error) {
	orgInput := &organizations.DescribeOrganizationInput{}
	orgResp, orgErr := svc.orgService.DescribeOrganization(orgInput)
	if orgErr != nil {
		return nil, orgErr
	}
	// fmt.Println(orgResp)

	accInput := &organizations.DescribeAccountInput{
		AccountId: aws.String(*orgResp.Organization.MasterAccountId),
	}
	accResp, accErr := svc.orgService.DescribeAccount(accInput)
	if accErr != nil {
		return nil, accErr
	}

	masterAccount := command.OrgNode{
		ID:         *accResp.Account.Id,
		Name:       *accResp.Account.Name,
		Type:       "ACCOUNT",
		Properties: make(map[string]string),
	}
	masterAccount.Properties["Email"] = *accResp.Account.Email
	masterAccount.Properties["AccountType"] = "master"

	response = &command.Organization{
		ID:            *orgResp.Organization.Id,
		MasterAccount: &masterAccount,
		Properties:    make(map[string]string),
	}
	return response, nil
}

// GetRoots provides a list of all the root accounts in a particular organization
func (svc *OrgService) GetRoots() (response []*command.OrgNode, err error) {
	listInput := &organizations.ListRootsInput{}
	resp, err := svc.orgService.ListRoots(listInput)
	if err != nil {
		fmt.Println("Failed to get list of accounts -", err)
	}
	fmt.Println(resp)
	for _, element := range resp.Roots {
		tmpNode := &command.OrgNode{
			ID:         *element.Id,
			Name:       *element.Name,
			Type:       "ROOT",
			Properties: make(map[string]string),
		}
		response = append(response, tmpNode)
	}
	return
}

// GetChildren provides all the children for a particular account or OU
func (svc *OrgService) GetChildren(id string) (response []*command.OrgNode, err error) {
	response = make([]*command.OrgNode, 0)
	ouInput := &organizations.ListChildrenInput{
		ChildType: aws.String("ORGANIZATIONAL_UNIT"),
		ParentId:  aws.String(id),
	}
	if ouInput == nil {
		return nil, errors.New("Unsupported type provided")
	}
	ouOutput, ouErr := svc.getChildrenPaginated(ouInput)
	if ouErr != nil {
		return nil, ouErr
	}
	fmt.Printf("Number of OU's under id %s - %d\n", id, len(ouOutput))
	for _, ouElement := range ouOutput {
		ouChild, ouChildErr := svc.DescribeGroup(*ouElement.Id)
		if ouChildErr != nil {
			fmt.Printf("Error - %s\nSkipping ou - %s...", ouChildErr, *ouElement.Id)
		}
		response = append(response, ouChild)
	}

	accInput := &organizations.ListChildrenInput{
		ChildType: aws.String("ACCOUNT"),
		ParentId:  aws.String(id),
	}
	if accInput == nil {
		return nil, errors.New("Unsupported type provided")
	}
	accOutput, accErr := svc.getChildrenPaginated(accInput)
	if accErr != nil {
		return nil, accErr
	}
	fmt.Printf("Number of accounts under id %s - %d\n", id, len(accOutput))
	for _, accElement := range accOutput {
		accChild, accChildErr := svc.DescribeAccount(*accElement.Id)
		if accChildErr != nil {
			fmt.Printf("Error - %s\nSkipping account - %s...", accChildErr, *accElement.Id)
		}
		response = append(response, accChild)
	}
	return response, nil
}

// DescribeGroup given an id for an ou provides additional information about that ou
func (svc *OrgService) DescribeGroup(id string) (response *command.OrgNode, err error) {
	ouInput := &organizations.DescribeOrganizationalUnitInput{
		OrganizationalUnitId: aws.String(id),
	}
	ouResp, ouErr := svc.orgService.DescribeOrganizationalUnit(ouInput)
	if ouErr != nil {
		return nil, ouErr
	}
	fmt.Println(ouResp)
	response = &command.OrgNode{
		ID:         *ouResp.OrganizationalUnit.Id,
		Name:       *ouResp.OrganizationalUnit.Name,
		Type:       "ORGANIZATIONAL_UNIT",
		Properties: make(map[string]string),
	}
	response.Properties["Arn"] = *ouResp.OrganizationalUnit.Arn

	return response, nil
}

// DescribeAccount given an id for an account provides additional information for that account
func (svc *OrgService) DescribeAccount(id string) (response *command.OrgNode, err error) {
	accInput := &organizations.DescribeAccountInput{
		AccountId: aws.String(id),
	}
	accResp, accErr := svc.orgService.DescribeAccount(accInput)
	if accErr != nil {
		return nil, accErr
	}
	fmt.Println(accResp)
	response = &command.OrgNode{
		ID:         *accResp.Account.Id,
		Name:       *accResp.Account.Name,
		Type:       "ACCOUNT",
		Properties: make(map[string]string),
	}
	response.Properties["Email"] = *accResp.Account.Email
	response.Properties["AccountType"] = "standard"
	response.Properties["Arn"] = *accResp.Account.Arn
	response.Properties["Status"] = *accResp.Account.Status
	response.Properties["JoinedMethod"] = *accResp.Account.JoinedMethod
	response.Properties["JoinedTimeStamp"] = accResp.Account.JoinedTimestamp.String()
	return response, nil
}

func (svc *OrgService) getChildrenPaginated(input *organizations.ListChildrenInput) (pages []organizations.Child, err error) {
	pages = []organizations.Child{}
	numPages, gotToEnd := 0, false
	err = svc.orgService.ListChildrenPages(input, func(p *organizations.ListChildrenOutput, last bool) bool {
		numPages++
		for _, t := range p.Children {
			pages = append(pages, *t)
		}
		if last {
			if gotToEnd {
				fmt.Println("goToEnd happened multiple times")
			}
			gotToEnd = true
		}
		return true
	})
	return pages, err
}
