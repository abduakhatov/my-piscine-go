package piscine
import "fmt"
func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	max := root
	if res := BTreeMax(root.Left); res != nil && res.Data > max.Data {
		max = res
	}
	if res := BTreeMax(root.Right); res != nil && res.Data > max.Data {
		max = res
	}
	fmt.Println(max.Data)
	return max
}
