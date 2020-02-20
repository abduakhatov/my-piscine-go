package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	// z01.PrintRune(rune(n+48))
	next_digit(rune(n+48), '0')
	z01.PrintRune('\n')
}

func next_digit(n rune, start rune) {
	if n == '0' {
		z01.PrintRune(',')
		z01.PrintRune(' ')
		return
	}
	for i := start; i <= '9'; i++ {
		z01.PrintRune(rune(i))
		next_digit(n-1, i+1)
	}
}
