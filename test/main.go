package main

import (
	"fmt"
	piscine ".."
)

func main() {
	s := "12345"
	s2 := "0000000012345"
	s3 := "012 345"
	s4 := "Hello World!"

	n := piscine.BasicAtoi2(s)
	n2 := piscine.BasicAtoi2(s2)
	n3 := piscine.BasicAtoi2(s3)
	n4 := piscine.BasicAtoi2(s4)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)

}
