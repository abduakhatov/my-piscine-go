package piscine

func NRune(s string, n int) rune {
	ss := []rune(s)
	for ind, val := range ss {
		if ind == n-1 {
			return val
		}
	}
	return 0
}
