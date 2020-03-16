package main

import (
	// "github.com/01-edu/z01"
	// "os"
	"fmt"
)

type NodeAddL struct {
	Next *NodeAddL
	Num  int
}

func pushBack(node *NodeAddL, val int) *NodeAddL {
	if node == nil {
		node = &NodeAddL{Num: val}
		return node
	}
	n := node
	for node.Next != nil {
		node = node.Next
	}
	node.Next = &NodeAddL{Num: val}
	return n
}

func Sortll(node *NodeAddL) *NodeAddL {
	tmp := node
	for i := node; i != nil; i = i.Next {
		for j := i; j != nil; j = j.Next {
			if i.Num < j.Num {
				i.Num, j.Num = j.Num, i.Num
			}
		}
	}
	return tmp
}

func main() {
	num1 := &NodeAddL{Num: 5}
	num1 = pushBack(num1, 1)
	num1 = pushBack(num1, 3)
	num1 = pushBack(num1, 1)
	num1 = pushBack(num1, 3)

	result := Sortll(num1)
	for tmp := result; tmp != nil; tmp = tmp.Next {
		fmt.Print(tmp.Num)
		if tmp.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}
