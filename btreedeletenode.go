package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == node {
		if root.Left != nil {
			root.Data = root.Left.Data
			node = root.Left
		} else if root.Right != nil {
			root.Data = root.Right.Data
			node = root.Right
		} else {
			return nil
		}
	}
	root.Left = BTreeDeleteNode(root.Left, node)
	root.Right = BTreeDeleteNode(root.Right, node)
	return root
}
