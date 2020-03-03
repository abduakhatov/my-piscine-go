package main

import (
	// "fmt"
	"os"
)

func MultplyOverflow(a, b, c int64) bool {
	if a < 0 && c < 0 {
		return  a * b + c < 0
	} else if a > 0 && c > 0 {
		return a * b + c > 0
	}
	return true
	
}

func Atoi(nbr string) int64 {
	var res int64 = 0
	var sign int64 = 1
	if nbr[0] == '-' {
		nbr = nbr[1:]
		sign *= -1
	}
	for _, digit := range nbr {
		tmp := int64(digit - '0')*sign
		if !MultplyOverflow(res, int64(10), tmp) {
			return 0
		}
		res = res*10 + tmp
	}
	return res
}

func main() {
	// fmt.Println("12345")
	// fmt.Println(Atoi("12345"))
	// fmt.Println()
	// fmt.Println("9223372036854775807")
	// fmt.Println(Atoi("9223372036854775807"))
	// fmt.Println()
	// fmt.Println("-9223372036854775809")
	// fmt.Println(Atoi("-9223372036854775809"))
	// fmt.Println()
	// fmt.Println("9223372036854775808")
	// fmt.Println(Atoi("9223372036854775808"))
	// fmt.Println()
	os.Stdout.Write("123")
	os.Stdout.Close()
}
