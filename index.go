package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	ss := []rune(toFind)
	bb := []rune(toFind)
	for ind, val := range ss {
		if val == bb[0] {
			return ind
		}
	}
	return -1
}
