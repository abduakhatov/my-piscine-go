package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	arr := []rune(str)
	for _, b := range arr {
		z01.PrintRune(b)
	}
}
