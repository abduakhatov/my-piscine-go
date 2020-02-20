package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	arr := []byte(str)
	for _, b := range arr {
		z01.PrintRune(rune(b))
	}
}
