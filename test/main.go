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

func main() {
	root := &piscine.TreeNode{Data: "04"}
	piscine.BTreeInsertData(root, "07")
	piscine.BTreeInsertData(root, "12")
	piscine.BTreeInsertData(root, "05")
	piscine.BTreeInsertData(root, "10")

	piscine.BTreeInsertData(root, "01")
	piscine.BTreeInsertData(root, "02")
	piscine.BTreeInsertData(root, "03")
	piscine.BTreeApplyByLevel(root, fmt.Println)
}

// package main

// import "fmt"

// type node struct {
//     value string
//     left  *node
//     right *node
// }

// func insert(n *node, v string) *node {
//     if n == nil {
//         return &node{v, nil, nil}
//     } else if v <= n.value {
//         n.left = insert(n.left, v)
//     } else {
//         n.right = insert(n.right, v)
//     }
//     return n
// }


// /* breadth first traversal */
// func breadth(n *node) {
//     if n != nil {
// 		s := []*node{n}

// 		fmt.Print(s[0])
// 		fmt.Print(s[1])
// 		fmt.Print(s[2])

//         for len(s) > 0 {
//             current_node := s[0]
//             fmt.Printf(current_node.value + " ")
//             s = s[1:]
//             if current_node.left != nil {
//                 s = append(s, current_node.left)
//             }
//             if current_node.right != nil {
//                 s = append(s, current_node.right)
//             }
//         }
//     }
// }

// func main() {
//     var root *node
//     root = insert(root, "F")
//     root = insert(root, "D")
//     root = insert(root, "H")
//     root = insert(root, "B")
//     root = insert(root, "E")
//     root = insert(root, "G")
//     root = insert(root, "J")
//     fmt.Println("\nBFS traversal")
//     breadth(root)
// }