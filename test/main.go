package main

import (
	"fmt"

	piscine ".."
)

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Parent != nil {
		if root.Parent.Data <= root.Data && root == root.Parent.Left {
			return false
		} else if root.Parent.Data > root.Data && root == root.Parent.Right {
			return false
		}
	}
	return BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
}

func main() {
	root := &piscine.TreeNode{Data: "04"}
	piscine.BTreeInsertData(root, "07")
	piscine.BTreeInsertData(root, "12")
	piscine.BTreeInsertData(root, "05")
	piscine.BTreeInsertData(root, "10")

	piscine.BTreeInsertData(root, "01")
	piscine.BTreeInsertData(root, "02")
	piscine.BTreeInsertData(root, "03")
	piscine.BTreeInsertData(root, "00")
	root.Left.Left.Data = "123"

	fmt.Println(BTreeIsBinary(root))
}
