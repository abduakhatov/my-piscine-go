package piscine

import (
	"github.com/01-edu/z01"
)

func ParseToBase(nbr int, base []rune, size int, sign int) {
	if nbr < size && nbr > -1*size {
		z01.PrintRune(base[nbr*sign])
		return
	}
	ParseToBase(nbr/size, base, size, sign)
	z01.PrintRune(base[nbr%size*sign])
}

func PrintNbrBase(nbr int, base string) {
	size := 0
	isValid := true
	for _, val := range base {
		if val == '-' || val == '+' {
			isValid = false
			size = 0
			break
		}
		size++
	}
	if size < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if base[j] == base[i] {
				isValid = false
				break
			}
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
