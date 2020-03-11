package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	f(root.Data)
	BTreeApplyByLevel(root.Left, f)
	BTreeApplyByLevel(root.Right, f)
}
