package aws

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/CloudCoreo/cli/pkg/command"
)

// AddChildNode creates and add a child node given a parent
func AddChildNode(nodeInfo *command.OrgNode, parent *command.TreeNode) (node *command.TreeNode, isRoot bool) {
	node = &command.TreeNode{
		Info:     nodeInfo,
		Parent:   parent,
		Children: make([]*command.TreeNode, 0),
	}
	if parent == nil {
		isRoot = true
	} else {
		parent.Children = append(parent.Children, node)
		isRoot = false
	}
	return
}

// PrintTree given the root node prints the entire tree
func PrintTree(root *command.TreeNode) {
	if root == nil {
		return
	}

	dummy := &command.TreeNode{}
	queue := make([]*command.TreeNode, 0)
	queue = append(queue, root)
	queue = append(queue, dummy)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == dummy {
			fmt.Println()
			continue
		} else {
			fmt.Print(node.Info.Type, "-", node.Info.ID, "    ")
		}
		for _, child := range node.Children {
			queue = append(queue, child)
		}
		queue = append(queue, dummy)
	}
}

func (svc *OrgService) buildOrgTree(root *command.TreeNode, rootID string) {
	if root == nil {
		return
	}
	children, rErr := svc.GetChildren(rootID)
	if rErr != nil {
		fmt.Printf("Failed to get children for id - %s; err - %s", rootID, rErr)
		return
	}

	for _, element := range children {
		switch element.Type {
		case "ORGANIZATIONAL_UNIT":
			// ouResp, ouErr := svc.DescribeGroup(element.ID)
			// if ouErr != nil {
			// 	fmt.Printf("Failed to get information about ou - %s; err - %s", element.ID, ouErr)
			// }
			node, isRoot := AddChildNode(element, root)
			if node == nil || isRoot == true {
				fmt.Printf("Unexpected failure while adding ou id - %s to parent", element.ID)
				return
			}
			svc.buildOrgTree(node, element.ID)
		case "ACCOUNT":
			// accResp, accErr := svc.DescribeAccount(element.ID)
			// if accErr != nil {
			// 	fmt.Printf("Failed to get information about accound - %s; err - %s", element.ID, accErr)
			// }
			node, isRoot := AddChildNode(element, root)
			if node == nil || isRoot == true {
				fmt.Printf("Unexpected failure while adding node id - %s to parent", element.ID)
				return
			}
		default:
			fmt.Println("Unknown type found while building tree")
			return
		}
	}
}

func (svc *OrgService) GetOrganizationTree() ([]*command.TreeNode, error) {
	/*
		stsSvc := sts.New(session.New())
		stsInput := &sts.AssumeRoleInput{
			DurationSeconds: aws.Int64(3600),
			Policy:          aws.String("{\"Version\":\"2012-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Action\":\"organizations:*\",\"Resource\":\"*\"}]}"),
			RoleArn:         aws.String("arn:aws:iam::116462199383:role/VMW_Rosetta_Role"),
			RoleSessionName: aws.String("AwsOrganizationsDemo"),
		}

		sresult, serr := stsSvc.AssumeRole(stsInput)
		if serr != nil {
			fmt.Println("Unable to assume role")
		}


		svc, err := NewOrgServiceWithCreds(sresult.Credentials)
		if err != nil {
			fmt.Println("Unable to initialize AWS Service -", err)
			return nil, err
		}
	*/

	// Collect information about the organization and master account
	org, orgErr := svc.DescribeOrganization()
	if orgErr != nil {
		fmt.Println("Failed to get org information -", orgErr)
		return nil, orgErr
	}
	fmt.Println(org)

	roots, rootsErr := svc.GetRoots()
	if rootsErr != nil {
		fmt.Println("Failed to get roots for the organization -", rootsErr)
		return nil, rootsErr
	}

	res := make([]*command.TreeNode, 0)
	for _, element := range roots {
		root, isRoot := AddChildNode(element, nil)
		if root == nil || isRoot == false {
			// fmt.Println("Unable to add organization root to tree")
			return nil, errors.New("Unable to add organization root to tree")
		}
		// Begin creating our organization tree by adding accounts and ou's to the root
		svc.buildOrgTree(root, element.ID)
		PrintTree(root)
		res = append(res, root)
	}

	return res, nil
}
