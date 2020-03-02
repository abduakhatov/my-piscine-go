package main

import (
	"fmt"
	piscine ".."
)

func f(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}


func main() {
	tab1 := []int{0, 1, 2, 3, 4, 5}
	tab2 := []int{0, 2, 1, 3}

	result1 := piscine.IsSorted(f, tab1)
	result2 := piscine.IsSorted(f, tab2)

	fmt.Println(result1)
	fmt.Println(result2)
}
