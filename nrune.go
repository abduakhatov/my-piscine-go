package piscine

func NRune(s string, n int) rune {
	for ind, val := range s {
		if ind == n-1 {
			return rune(val)
		}
	}
	return 0
}
