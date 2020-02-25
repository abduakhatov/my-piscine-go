package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	bb := []rune(toFind)
	for ind, val := range s {
		if val == bb[0] {
			return ind
		}
	}
	return -1
}
