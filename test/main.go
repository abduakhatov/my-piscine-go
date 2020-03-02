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
	tab3 := []int{-859847, 187500, 734084, 571713, 760402, 62627, -871811, 165954}

	result1 := piscine.IsSorted(f, tab1)
	result2 := piscine.IsSorted(f, tab2)
	result3 := piscine.IsSorted(f, tab3)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}

// IsSorted((func(int, int) int)(0x5545f0), 
// []int{-859847, 187500, 734084, 571713, 760402, 62627, -871811, 165954}) == true
//  instead of false
// FAIL