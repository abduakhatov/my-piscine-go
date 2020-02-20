package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	dot := 1
	if n == 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	if n < 0 {
		dot *= -1
		z01.PrintRune('-')
	}
	_n := n
	length := 0
	for _n*dot > 0 {
		_n /= 10
		length += 1
	}
	n_arr := []int{}
	digit := 0
	for i := length - 1; n*dot >= 1; i-- {
		digit = int((n % 10) * dot)
		n /= 10
		n_arr = append(n_arr, digit)
	}
	for i := length-1; i >= 0; i-- {
		z01.PrintRune(48 + rune(n_arr[i]))
	}
	z01.PrintRune('\n')
}
