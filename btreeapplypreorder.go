package piscine

/*
type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}
*/

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	// left
	// visit
	// right
	if root == nil {
		return
	}
	f(root.Data)
	BTreeApplyPostorder(root.Left, f)
	BTreeApplyPostorder(root.Right, f)
}
