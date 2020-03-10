package piscine

import "fmt"

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	for i := n2; i != nil; i = i.Next {
		fmt.Println(i.Data)
	}
	fmt.Println()
	for i := n1; i != nil; i = i.Next {
		// fmt.Println(i.Data)
		n1 = SortListInsert(n2, i.Data)
	}
	return n1
}
