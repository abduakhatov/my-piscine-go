package piscine

func Split(str, charset string) []string {
	char_len := get_length(charset)
	str_len := get_length(str)
	runes := []rune(str)
	size := 0
	for ind := 0; ind < str_len-char_len; ind++ {
		if string(runes[ind:ind+char_len]) == charset {
			size++
		}
	}
	resArr := make([]string, size+1)
	i := 0
	start := 0
	ind := 0
	for ; ind < str_len-char_len; ind++ {
		if string(runes[ind:ind+char_len]) == charset {
			resArr[i] = string(runes[start:ind])
			i++
			start = ind + char_len
		}
	}
	resArr[i] = string(runes[start:ind])
	return resArr
}
