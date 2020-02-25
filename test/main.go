package main

import (
	"fmt"
	piscine ".."
)

func main() {

	fmt.Println(piscine.IsPrintable("010203"))
	fmt.Println(piscine.IsPrintable("01,02,03"))
	fmt.Println(piscine.IsPrintable("Hello"))
	fmt.Println(piscine.IsPrintable("Hello\n"))
	fmt.Println(piscine.IsPrintable(""))
}