package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Parent != nil {
		if root.Parent.Data > root.Data && root == root.Parent.Left {
			return false
		} else if root.Parent.Data <= root.Data && root == root.Parent.Right {
			return false
		}
	}
	return BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
}
