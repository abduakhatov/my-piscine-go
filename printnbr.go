package piscine

import "github.com/01-edu/z01"

func print_next(n, dot int) {
	if n == 0 {
		return
	}
	digit := int((n % 10) * dot)
	_n := int(n / 10)
	print_next(_n, dot)
	z01.PrintRune(48 + rune(digit))
}

func PrintNbr(n int) {
	dot := 1
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		dot *= -1
		z01.PrintRune('-')
	}
	print_next(n, dot)
}
