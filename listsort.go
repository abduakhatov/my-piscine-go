package piscine

// import "fmt"

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	for i := l; i != nil; i = i.Next {
		for j := i.Next; j != nil; j = j.Next {
			// fmt.Println(i.Data, j.Data)
			if i.Data > j.Data {
				i.Data, j.Data = j.Data, i.Data
			}
		}
	}
	return l
}
