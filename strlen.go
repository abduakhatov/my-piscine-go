package piscine

func StrLen(str string) int {
	counter := 0
	for ind := range str {
		if str[ind] > 31 {
			counter += 1
		}
	}
	return counter
}
