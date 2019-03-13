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
			node, isRoot := AddChildNode(element, root)
			if node == nil || isRoot == true {
				fmt.Printf("Unexpected failure while adding ou id - %s to parent", element.ID)
				return
			}
			svc.buildOrgTree(node, element.ID)
		case "ACCOUNT":
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

// GetOrganizationTree returns an array of treenode roots
func (svc *OrgService) GetOrganizationTree() ([]*command.TreeNode, error) {
	err := svc.init()
	if err != nil {
		return nil, err
	}
	// Collect information about the organization and master account
	org, orgErr := svc.DescribeOrganization()
	if orgErr != nil {
		return nil, orgErr
	}
	fmt.Println(org)

	roots, rootsErr := svc.GetRoots()
	if rootsErr != nil {
		return nil, rootsErr
	}

	res := make([]*command.TreeNode, 0)
	for _, element := range roots {
		root, isRoot := AddChildNode(element, nil)
		if root == nil || isRoot == false {
			return nil, errors.New("Unable to add organization root to tree")
		}
		// Begin creating our organization tree by adding accounts and ou's to the root
		svc.buildOrgTree(root, element.ID)
		// PrintTree(root)
		res = append(res, root)
	}

	return res, nil
}
