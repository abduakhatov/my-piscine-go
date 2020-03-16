package piscine

/*
type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}
*/

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	// visit
	// left
	// right
	if root == nil {
		return
	}
	f(root.Data)
	BTreeApplyPreorder(root.Left, f)
	BTreeApplyPreorder(root.Right, f)
}
