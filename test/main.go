package main

import (
	"fmt"
	piscine ".."
)

func main() {
	link := &piscine.List{}

	piscine.ListPushBack(link, "1")
	piscine.ListPushBack(link, "2")
	piscine.ListPushBack(link, "3")
	piscine.ListPushBack(link, "5")

	piscine.ListForEach(link, piscine.Add2_node)

	it := link.Head
	for it != nil {
		fmt.Println(it.Data)
		it = it.Next
	}
}
