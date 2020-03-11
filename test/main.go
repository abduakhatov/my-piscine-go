package main

import (
	"fmt"
	piscine ".."
)

func main() {
	root := &piscine.TreeNode{Data: "40"}
	piscine.BTreeInsertData(root, "07")
	piscine.BTreeInsertData(root, "12")
	piscine.BTreeInsertData(root, "05")
	piscine.BTreeInsertData(root, "10")

	piscine.BTreeInsertData(root, "01")
	piscine.BTreeInsertData(root, "02")
	piscine.BTreeInsertData(root, "03")
	piscine.BTreeInsertData(root, "00")
	root.Left.Left.Data = "123"

	fmt.Println(root.Left.Left.Data)

	// fmt.Println(piscine.BTreeIsBinary(root))
}
