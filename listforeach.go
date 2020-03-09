package piscine

// import "fmt"

func ListForEach(l *List, f func(*NodeL)) {
	// tmp := *l.Head
	// l.Head = nil
	tmp := &List{}
	for ; l.Head != nil; l.Head = l.Head.Next {
		f(l.Head)
		ListPushBack(tmp, l.Head.Data)
	}
	l.Head = tmp.Head
}

func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
}
