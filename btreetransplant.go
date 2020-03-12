package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == node {
		root = rplc
	} else {
		root.Left = BTreeTransplant(root.Left, node, rplc)
		root.Right = BTreeTransplant(root.Right, node, rplc)
	}
	return root
}
