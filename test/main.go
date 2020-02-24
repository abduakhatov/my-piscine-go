package main

import (
	"fmt"
	piscine ".."
)

func main() {
	arg1 := 5
	arg2 := 4
	arg3 := -1
	arg4 := 0
	arg5 := 2147483647
	fmt.Println(arg1, piscine.IsPrime(arg1))
	fmt.Println(arg2, piscine.IsPrime(arg2))
	fmt.Println(arg3, piscine.IsPrime(arg3))
	fmt.Println(arg4, piscine.IsPrime(arg4))
	fmt.Println(arg5, piscine.IsPrime(arg5))
}
