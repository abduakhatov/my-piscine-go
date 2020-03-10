package piscine

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Tail == nil && data != nil {
		l.Tail = n
	}
	n.Next = l.Head
	l.Head = n
}
