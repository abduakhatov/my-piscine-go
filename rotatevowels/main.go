package main

import (
	"github.com/01-edu/z01"
	"os"
	// "fmt"
)

func main() {
	args := os.Args[1:]
	str := ""
	for _, val := range args {
		str = str + " " + val
	}
	indexes := []int{}
	length := 0
	hasNext := false
	for ind, val := range str {
		if val == 'a' || val == 'e' || val == 'i' || val == 'o' || val == 'u' || val == 'A' || val == 'E' || val == 'I' || val == 'O' || val == 'U' {
			indexes = append(indexes, ind-1)
			length++
		}
		hasNext = true
	}
	if !hasNext {
		z01.PrintRune('\n')
		return
	}
	str = str[1:]
	strRune := []rune(str)
	for i := 0; i < length/2 && length != 0; i++ {
		strRune[indexes[i]], strRune[indexes[length-i-1]] = strRune[indexes[length-i-1]], strRune[indexes[i]]
	}
	for _, val := range strRune {
		z01.PrintRune(val)
	}
	z01.PrintRune('\n')
}
