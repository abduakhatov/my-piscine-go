package main

import (
	"fmt"
	piscine ".."
)

func main() {

	fmt.Println(piscine.IsNumeric("010203"))
	fmt.Println(piscine.IsNumeric("01,02,03"))
	fmt.Println(piscine.IsAlpha(""))
}