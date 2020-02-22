package piscine

import "github.com/01-edu/z01"

func ValidComb(comb int) bool {
	for comb != 0 {
		if comb >= 10 && comb%10 <= (comb/10)%10 {
			return false
		}
		comb /= 10
	}
	return true
}

func PrintNext(numb int) {
	if numb == 0 {
		z01.PrintRune('0')
		return
	}
	if numb >= 10 {
		PrintNext(numb / 10)
	}
	digit := '0'
	for i := 0; i < numb%10; i++ {
		digit++
	}
	z01.PrintRune(digit)
}

func PrintCombN(n int) {
	postionLength := 1
	for i := 1; i < n; i++ {
		postionLength *= 10
	}
	st := false
	for i := postionLength / 10; i <= postionLength*9; i++ {
		if !ValidComb(i) {
			continue
		}
		if st {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		if i <= postionLength && postionLength != 1 {
			z01.PrintRune('0')
		}
		PrintNext(i)
		st = true
	}
	z01.PrintRune('\n')
}
