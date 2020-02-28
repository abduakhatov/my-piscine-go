package piscine

import (
	"github.com/01-edu/z01"
)

func ParseToBase(nbr int, base []rune, size int, sign int) {
	if nbr < size {
		z01.PrintRune(base[int(nbr*sign)])
		return
	}
	ParseToBase(nbr/size, base, size, sign)
	z01.PrintRune(base[int(nbr%size*sign)])
}

func PrintNbrBase(nbr int, base string) {
	size := 0
	// isDifferent := false
	for range base {
		size++
	}
	if size < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	// for i := 0; i < size-1; i++ {
	// }
	dot := 1
	if nbr < 0 {
		dot *= -1
		z01.PrintRune('-')
	}
	ParseToBase(nbr, []rune(base), size, dot)
}
