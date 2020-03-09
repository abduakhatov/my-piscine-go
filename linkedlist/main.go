package main

import "fmt"

type node struct {
	data int
	next *node
}

func insert(head *node, data int) *node {
	fmt.Println("ins:", &head, head)
	n := &node{data: data}
	fmt.Println("n:", n, &n, *n)
	if head != nil {
		n.next = head
	}
	fmt.Println("n2:", n, &n, *n)
	return n
}

func PrintList(head *node) {
	for head != nil {
		fmt.Print(head.data, " -> ")
		head = head.next
	}
	fmt.Println(nil)
}

func main() {
	var link *node
	fmt.Println(&link)
	link = insert(link, 1)
	fmt.Println("\n", &link, *link )
	link = insert(link, 2)
	fmt.Println("\n", &link, *link )
	link = insert(link, 3)
	fmt.Println("\n", &link, *link )

	PrintList(link)
}