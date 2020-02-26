package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arg := os.Args[1:]
	charCase := 0
	if len(arg) > 0 && arg[0] == "--upper" {
		charCase = -32
		arg = arg[1:]
	}
	for _, a := range arg {
		digit := 0
		for _, b := range a {
			digit = digit*10 + int(b-'0')
		}
		if 1 <= digit && digit <= 26 {
			z01.PrintRune('a' + rune(digit+charCase) - 1)
		} else {
			z01.PrintRune(' ')
		}
	}
}
