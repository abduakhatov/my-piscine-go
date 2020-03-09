package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	if l.Head == nil {
		l.Head = &NodeL{Data: data}
		return
	}
	n := l.Head
	for n.Next != nil {
		n = n.Next
	}
	n.Next = &NodeL{Data: data}
}
