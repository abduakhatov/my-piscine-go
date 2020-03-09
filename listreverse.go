package piscine

import "fmt"

func ListReverse(l *List) {
	newList := &List{}
	for ; l.Head != nil; l.Head = l.Head.Next {
		fmt.Println(l.Head.Data)
		ListPushFront(newList, l.Head.Data)
	}
	l.Head = newList.Head
}
