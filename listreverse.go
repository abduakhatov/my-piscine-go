package piscine

func ListReverse(l *List) {
	newList := &List{}
	for ; l.Head != nil; l.Head = l.Head.Next {
		ListPushFront(newList, l.Head.Data)
	}
	l.Head = newList.Head
}
