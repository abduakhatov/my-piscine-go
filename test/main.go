package main

import (
	piscine ".."
	"fmt"
)

func PrintElem(node *piscine.NodeL) {
	fmt.Println(node.Data)
}

func StringToInt(node *piscine.NodeL) {
	node.Data = 2
}

func PrintList(l *piscine.List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, "->")
		it = it.Next
	}
	fmt.Print("nil","\n")
}

func main() {
	link := &piscine.List{}

	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, "hello")
	piscine.ListPushBack(link, 3)
	piscine.ListPushBack(link, "there")
	piscine.ListPushBack(link, 23)
	piscine.ListPushBack(link, "!")
	piscine.ListPushBack(link, 54)

	PrintList(link)

	fmt.Println("--------function applied--------")
	piscine.ListForEachIf(link, PrintElem, piscine.IsPositive_node)

	piscine.ListForEachIf(link, StringToInt, piscine.IsNotNumeric_node)

	fmt.Println("--------function applied--------")
	PrintList(link)

	fmt.Println()
}
