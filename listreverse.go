package piscine

import "fmt"

func ListReverse(l *List) {
	newList := &List{}
	tmp := *l
	for ; tmp.Head != nil; tmp.Head = tmp.Head.Next {
		ListPushFront(newList, tmp.Head.Data)
	}
	for ; l.Head != nil; l.Head = l.Head.Next {
		fmt.Println(l.Head.Data, "-")
	}
	l.Head = newList.Head
}
