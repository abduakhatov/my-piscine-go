package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	max := &TreeNode{Data: root.Data}
	left := BTreeMax(root.Left)
	right := BTreeMax(root.Right)

	if left != nil && BtreeAtoi(left.Data) < BtreeAtoi(root.Data) {
		max = &TreeNode{Data: left.Data}
	}
	if right != nil && BtreeAtoi(right.Data) < BtreeAtoi(root.Data) {
		max = &TreeNode{Data: right.Data}
	}
	return max
}
