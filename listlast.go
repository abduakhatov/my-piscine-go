package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	node := l.Head
	for l.Head != nil {
		node = l.Head
		l.Head = l.Head.Next

	}
	return node.Data
}
