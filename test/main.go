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

func PrintList(l *piscine.NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *piscine.NodeI, data int) *piscine.NodeI {
	n := &piscine.NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {
	var link *piscine.NodeI
	var link2 *piscine.NodeI

	link = listPushBack(link, 3)
	link = listPushBack(link, 5)
	link = listPushBack(link, 7)

	link2 = listPushBack(link2, -2)
	link2 = listPushBack(link2, 9)

	// PrintList(link)
	// PrintList(link2)

	// PrintList(piscine.SortListInsert(link, -2))

	PrintList(piscine.SortedListMerge(link2, link))
}

