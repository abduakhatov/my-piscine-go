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
	s5 := "+1234"
	s6 := "-1234"
	s7 := "++1234"
	s8 := "--1234"
	s9 := "123-"

	n := piscine.Atoi(s)
	n2 := piscine.Atoi(s2)
	n3 := piscine.Atoi(s3)
	n4 := piscine.Atoi(s4)
	n5 := piscine.Atoi(s5)
	n6 := piscine.Atoi(s6)
	n7 := piscine.Atoi(s7)
	n8 := piscine.Atoi(s8)
	n9 := piscine.Atoi(s9)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n6)
	fmt.Println(n7)
	fmt.Println(n8)
	fmt.Println(n9)
}