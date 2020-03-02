package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, val := range tab {
		if f(val) {
			count++
		}
	}
	return count
}
