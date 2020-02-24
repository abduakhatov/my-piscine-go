package main

import (
	"fmt"
	piscine ".."
)

func main() {
	arg := 0
	fmt.Println(piscine.RecursiveFactorial(arg))
	arg = 2
	fmt.Println(piscine.RecursiveFactorial(arg))
	arg = 3
	fmt.Println(piscine.RecursiveFactorial(arg))
	arg = 4
	fmt.Println(piscine.RecursiveFactorial(arg))
	arg = 5
	fmt.Println(piscine.RecursiveFactorial(-8778634010981231238))

}
