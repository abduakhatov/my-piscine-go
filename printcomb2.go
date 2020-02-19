package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	_min := '0'
	_max := '9'
	for i := _min; i <= _max; i++ {
		for j := _min; j <= _max; j++ {
			for k := _min; k <= _max; k++ {
				for l := _min; l <= _max; l++ {
					if i > k {
						continue
					}
					if i != k && j > l {
						continue
					}
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(k)
					z01.PrintRune(l)
					if i == _max && j == i-1 && k == _max && l == _max {
						continue
					}
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
