package piscine

func AlphaCount(str string) int {
	counter := 0
	for _, val := range str {
		if (val >= 65 && val <= 90) || (val >= 97 && val <= 122) {
			counter += 1
		}
	}
	return counter
}
