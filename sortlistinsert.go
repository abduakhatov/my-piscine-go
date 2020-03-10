package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	n := &NodeI{Data: data_ref}
	n.Next = l
	l = n
	l = ListSort(l)
	return l
}
