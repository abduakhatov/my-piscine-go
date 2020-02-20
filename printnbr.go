package main

import "github.com/01-edu/z01"

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
	_n := n
	length := 0
	for _n*dot > 0 {
		_n /= 10
		length += 1
	}
	n_arr := make([]int, length)
	digit := 0
	for i := length - 1; n*dot >= 1; i-- {
		digit = int((n % 10) * dot)
		n /= 10
		n_arr[i] = digit
	}
	for i := 0; i < length; i++ {
		z01.PrintRune(48 + rune(n_arr[i]))
	}
	z01.PrintRune('\n')
}

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}