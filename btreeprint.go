package piscine

import "fmt"

func BtreePrint(tree *TreeNode) {
	if tree == nil {
		return
	}
	fmt.Print("-")
	fmt.Println(tree.Data)
	fmt.Print(" | ")
	BtreePrint(tree.Right)
	fmt.Println(" | ")
	BtreePrint(tree.Left)

}
