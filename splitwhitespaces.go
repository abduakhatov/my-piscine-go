package piscine

func SplitWhiteSpaces(str string) []string {
	runes := []rune(str)
	counter := 1
	chars := 0
	for _, val := range runes {
		if val == '\n' || val == '\t' || val == ' ' {
			counter++
		}
		chars++
	}
	res_arr := make([]string, counter)
	i := 0
	start := 0
	for ind, val := range runes {
		if val == '\n' || val == '\t' || val == ' ' {
			res_arr[i] = string(runes[start:ind])
			start = ind + 1
			i++
		}
	}
	res_arr[i] = string(runes[start:])
	return res_arr
}
