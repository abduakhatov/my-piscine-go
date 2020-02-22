package piscine

import "github.com/01-edu/z01"

func get_length(previous string) int {
	counter := 0
	for previous != "" {
		previous = previous[1:]
		counter += 1
	}
	return counter
}

func next_digit(start rune, previous string, n int) string {
	result := ""
	if n == 0 {
		previous = ", " + previous
		return previous
	}

	for i := start; i <= '9'; i++ {
		_prev := previous + string(i)
		_i := i + 1
		_len := n - 1
		result += next_digit(_i, _prev, _len)
	}
	return result
}

func PrintCombN_cheating(n int) {
	result := next_digit('0', "", n)
	result = result[2:]
	length := get_length(result)
	for j := 0; j < length; j++ {
		z01.PrintRune(rune(result[j]))
	}
	z01.PrintRune('\n')
}
