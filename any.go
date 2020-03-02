package piscine

func Any(f func(string) bool, arr []string) bool {
	isNumeric := false
	for _, val := range arr {
		if f(val) {
			isNumeric = true
		}
	}
	return isNumeric
}
