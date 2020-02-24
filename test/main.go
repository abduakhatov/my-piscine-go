package main

import (
	"fmt"
	piscine ".."
)

func main() {
	arg := 1
	fmt.Println(piscine.IterativeFactorial(arg))
	arg = 2
	fmt.Println(piscine.IterativeFactorial(arg))
	arg = 3
	fmt.Println(piscine.IterativeFactorial(arg))
	arg = 4
	fmt.Println(piscine.IterativeFactorial(arg))
	arg = 5
	fmt.Println(piscine.IterativeFactorial(-8778634010981231238))

}
