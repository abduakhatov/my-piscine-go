package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	for i := n2; i != nil; i = i.Next {
		n1 = SortListInsert(n1, i.Data)
	}
	return n1
}
