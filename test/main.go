package main

import (
	"fmt"

	piscine ".."
)

func main() {
	link := &piscine.List{}
	link2 := &piscine.List{}

	// piscine.ListPushBack(link, "three")
	// piscine.ListPushBack(link, 3)
	// piscine.ListPushBack(link, "1")

	piscine.ListPushBack(link, "-2381620042740041718")
	piscine.ListPushBack(link, "-2381620042740041718")
	piscine.ListPushBack(link, "1223605871337147087")
	piscine.ListPushBack(link, "3026360979856967082")

	fmt.Println(piscine.ListLast(link))
	fmt.Println(piscine.ListLast(link2))
}
