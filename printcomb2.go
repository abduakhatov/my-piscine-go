package package

import "github.com/01-edu/z01"

func PrintComb2() {
	_min := '0'
	_max := '9'
	var st int = 0
	l := '0'
	for i := _min; i <= _max; i++ {
		for j := _min; j <= _max; j++ {
			for k := i; k <= _max; k++ {
				l = j
				for ; l <= _max; l++ {
					// if i == k && (j > l || j == l) {
					// 	continue
					// }
					if st > 0 {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(k)
					z01.PrintRune(l)
					st = 1
				}
			}
		}
	}
	z01.PrintRune('\n')
}
