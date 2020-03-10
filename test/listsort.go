package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	for i := l; i != nil; i = i.Next {
		for j := l; j != nil; j = j.Next {
			if i.Data > j.Data {
				i.Next, j.Next = j.Next, i
			}
		}
	}
	return &l
}
