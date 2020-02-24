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
	arg6 := 1
	fmt.Println(arg1, piscine.FindNextPrime(arg1))
	fmt.Println(arg2, piscine.FindNextPrime(arg2))
	fmt.Println(arg3, piscine.FindNextPrime(arg3))
	fmt.Println(arg4, piscine.FindNextPrime(arg4))
	fmt.Println(arg5, piscine.FindNextPrime(arg5))
	fmt.Println(arg6, piscine.FindNextPrime(arg6))
}
