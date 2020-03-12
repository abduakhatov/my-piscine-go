package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left != nil && root.Left == node {
		root.Left = node
		root.Left.Parent = root

	} else if root.Right != nil && root.Right == node {
		root.Right = node
		root.Right.Parent = root
	} else {
		root.Left = BTreeTransplant(root.Left, node, rplc)
		root.Right = BTreeTransplant(root.Right, node, rplc)
	}
	return root
}
