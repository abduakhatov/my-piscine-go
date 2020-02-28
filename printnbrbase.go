package piscine

import (
	"github.com/01-edu/z01"
)

func ParseToBase(nbr int, base []rune, size int, sign int) {
	if nbr*sign < size {
		z01.PrintRune(base[int(nbr*sign)])
		return
	}
	ParseToBase(nbr/size, base, size, sign)
	z01.PrintRune(base[int(nbr%size*sign)])
}

func PrintNbrBase(nbr int, base string) {
	size := 0
	isValid := false
	for _, val := range base {
		size++
		if !(('a' <= val && val <= 'z') || ('A' <= val && val <= 'Z') || ('0' <= val && val <= '9')) { 
			isValid = false
			size = 0
			break
		}
	}
	if size < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	for i := 0; i < size-1; i++ {
		if base[i] != base[i+1] {
			isValid = true
			break
		}
	}
	if !isValid {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	dot := 1
	if nbr < 0 {
		dot *= -1
		z01.PrintRune('-')
	}
	ParseToBase(nbr, []rune(base), size, dot)
}
