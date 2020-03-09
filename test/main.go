package main

import (
	"fmt"
	piscine ".."
)

func main() {
	link := &piscine.List{}

	piscine.ListPushBack(link, "hello")
	piscine.ListPushBack(link, "how are")
	piscine.ListPushBack(link, "you")
	piscine.ListPushBack(link, 1)

	fmt.Println(piscine.ListAt(link.Head, 3).Data)
	fmt.Println(piscine.ListAt(link.Head, 1).Data)
	fmt.Println(piscine.ListAt(link.Head, 7))

	fmt.Println()

	piscine.ListPushBack(link, "I")
	piscine.ListPushBack(link, "1")
	piscine.ListPushBack(link, 2)
	piscine.ListPushBack(link, "something")
	fmt.Println(piscine.ListAt(link.Head, 4).Data)
}
