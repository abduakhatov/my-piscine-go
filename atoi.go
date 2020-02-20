package main

import "fmt"

func Atoi(s string) int {
	_s := []byte(s)
	count := 0
	for range _s {
		count += 1
	}
	result := 0
	pow := 1
	tmp := 0
	sign := 1
	adj_chars := false
	for i := count - 1; i >= 0; i-- {
		tmp = (int(_s[i]) - 48)
		if tmp >= 0 && tmp <= 9 {
			result += (tmp * pow)
			pow *= 10
		} else if !adj_chars && (tmp == -5 || tmp == -3) {
			adj_chars = true
			if tmp == -3 {
				sign = -1
			}
		} else {
			return 0
		}
	}
	return result * sign
}

func main() {
	s := "12345"
	s2 := "0000000012345"
	s3 := "012 345"
	s4 := "Hello World!"
	s5 := "+1234"
	s6 := "-1234"
	s7 := "++1234"
	s8 := "--1234"

	n := Atoi(s)
	n2 := Atoi(s2)
	n3 := Atoi(s3)
	n4 := Atoi(s4)
	n5 := Atoi(s5)
	n6 := Atoi(s6)
	n7 := Atoi(s7)
	n8 := Atoi(s8)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n6)
	fmt.Println(n7)
	fmt.Println(n8)
}
