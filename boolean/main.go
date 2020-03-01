package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(str string) {
	for _, val := range []rune(str) {
		z01.PrintRune(val)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	args := os.Args[1:]
	length := 0
	for range args {
		length++
	}
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	if isEven(length) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
