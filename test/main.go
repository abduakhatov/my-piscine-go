package main

import (
	"fmt"
	piscine ".."
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	result := piscine.Map(piscine.IsPrime, arr)
	fmt.Println(result)
}
