package main

import (
	"fmt"
	piscine ".."
)

func main() {
	arg := 0
	fmt.Println(piscine.RecursivePower(2, arg))
	arg = 2
	fmt.Println(piscine.RecursivePower(2, arg))
	arg = 3
	fmt.Println(piscine.RecursivePower(2, arg))
	arg = 4
	fmt.Println(piscine.RecursivePower(2, arg))
	arg = -5
	fmt.Println(piscine.RecursivePower(2, arg))

}
