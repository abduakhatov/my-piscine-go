package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	_min := '0'
	_max := '9'
	var st int = 0
	for i := _min; i <= _max; i++ {
		for j := _min; j <= _max; j++ {
			for k := _min; k <= _max; k++ {
				if i < j && j < k {
					if st > 0 {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					st = 1
				}
			}
		}
	}
	z01.PrintRune('\n')
}
