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
	root := &piscine.TreeNode{Data: "05"}
	piscine.BTreeInsertData(root, "07")
	piscine.BTreeInsertData(root, "12")
	piscine.BTreeInsertData(root, "06")
	piscine.BTreeInsertData(root, "10")

	piscine.BTreeInsertData(root, "02")
	piscine.BTreeInsertData(root, "03")
	piscine.BTreeInsertData(root, "04")
	piscine.BTreeInsertData(root, "00")
	root.Left.Left.Data = "123"

	fmt.Println(piscine.BTreeIsBinary(root))
}

// package main

// import (
// 	"fmt"
// 	piscine ".."
// )

// func main() {
// 	root := &piscine.TreeNode{Data: "4"}
// 	piscine.BTreeInsertData(root, "1")
// 	piscine.BTreeInsertData(root, "7")
// 	piscine.BTreeInsertData(root, "23")
// 	max := piscine.BTreeMax(root)
// 	fmt.Println(max.Data)
// }
