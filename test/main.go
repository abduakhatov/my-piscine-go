package main

import (
	"fmt"
	piscine ".."
)

func PrintTree(root *piscine.TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Data)
	PrintTree(root.Left)
	PrintTree(root.Right)
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
	PrintTree(root)
	// fmt.Println(root.Left.Left.Data)
	// fmt.Println(piscine.BTreeIsBinary(root))
}
