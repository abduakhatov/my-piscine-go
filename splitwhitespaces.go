package piscine

func SplitWhiteSpaces(str string) []string {
	runes := []rune(str)
	counter := 0
	for _, val := range runes {
		if val == '\n' || val == '\t' || val == ' ' {
			counter++
		}
	}
	res_arr := make([]string, counter)
	res := ""
	i := 0
	for _, val := range runes {
		if val == '\n' || val == '\t' || val == ' ' {
			res_arr[i] = res
			res = ""
			i++
		} else {
			res += string(val)
		}
	}
	return res_arr
}
