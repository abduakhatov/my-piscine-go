package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	min := '0'
	max := '9'
	for i := min; i <= max; i++ {
		for j := min; j <= max; j++ {
			z01.PrintRune(i)
			z01.PrintRune(j)
			if i == max && j == max {
				continue
			} 
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
