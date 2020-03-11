package piscine
// import "fmt"
func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	
	left := BTreeMax(root.Left)
	if left == nil {
		left = root
	}
	if right := BTreeMax(root.Right); right != nil && right.Data > left.Data {
		left = right
	}
	return left
}
