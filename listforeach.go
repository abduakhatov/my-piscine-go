package piscine

// import "fmt"

func ListForEach(l *List, f func(*NodeL)) {
	tmp := *l.Head
	l.Head = nil
	for ; tmp.Next != nil; tmp = *tmp.Next {
		f(&tmp)
		ListPushBack(l, tmp.Data)
	}
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
