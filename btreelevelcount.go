package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	resLeft := BTreeLevelCount(root.Left) + 1
	resRight := BTreeLevelCount(root.Right) + 1
	if resLeft > resRight {
		resRight = resLeft
	}
	return resRight
}
