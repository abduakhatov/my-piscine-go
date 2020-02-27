package piscine

import (
	"github.com/01-edu/z01"
)

func PrintWordsTables(table []string) {
	for _, a := range table {
		for _, b := range []rune(a) {
			z01.PrintRune(b)
		}
		z01.PrintRune('\n')
	}
}
