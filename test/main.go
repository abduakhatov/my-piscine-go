// package main

// import (
// 	"fmt"
// 	piscine ".."
// )

// func PrintTree(root *piscine.TreeNode) {
// 	if root == nil {
// 		return
// 	}
// 	fmt.Println(root.Data)
// 	PrintTree(root.Left)
// 	PrintTree(root.Right)
// }


// func main() {
// 	root := &piscine.TreeNode{Data: "05"}
// 	piscine.BTreeInsertData(root, "07")
// 	piscine.BTreeInsertData(root, "12")
// 	piscine.BTreeInsertData(root, "06")
// 	piscine.BTreeInsertData(root, "10")

// 	piscine.BTreeInsertData(root, "02")
// 	piscine.BTreeInsertData(root, "03")
// 	piscine.BTreeInsertData(root, "04")
// 	piscine.BTreeInsertData(root, "00")
// 	root.Left.Left.Data = "123"

// 	fmt.Println(piscine.BTreeIsBinary(root))
// }

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
<<<<<<< HEAD
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
=======
	root := &piscine.TreeNode{Data: "4"}
	piscine.BTreeInsertData(root, "1")
	piscine.BTreeInsertData(root, "7")
	piscine.BTreeInsertData(root, "5")

	// root := &piscine.TreeNode{Data: "04"}
	// piscine.BTreeInsertData(root, "07")
	// piscine.BTreeInsertData(root, "12")
	// piscine.BTreeInsertData(root, "05")
	// piscine.BTreeInsertData(root, "10")

	// piscine.BTreeInsertData(root, "01")
	// piscine.BTreeInsertData(root, "02")
	// piscine.BTreeInsertData(root, "03")
	
	node := piscine.BTreeSearchItem(root, "4")
	fmt.Println("Before delete:")
	piscine.BTreeApplyInorder(root, fmt.Println)
	root = piscine.BTreeDeleteNode(root, node)
	fmt.Println("After delete:")
	piscine.BTreeApplyInorder(root, fmt.Println)
>>>>>>> 553ff245236c5e6ff0632568a5c09f92411a77c3
}
