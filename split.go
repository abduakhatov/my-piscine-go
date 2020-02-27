package piscine

// import "fmt"

func Split(str, charset string) []string {
	char_len := get_length(charset)
	str_len := get_length(str)
	// runes := []rune(str)
	size := 0
	for ind := 0; ind <= str_len-char_len; ind++ {
		if string(str[ind:ind+char_len]) == charset {
			size++
		}
	}
	// fmt.Println(size)
	resArr := make([]string, size+1)
	i := 0
	start := 0
	ind := 0
	for ; ind <= str_len-char_len; ind++ {
		// fmt.Println(str_len, "/", ind, "+", char_len, string(runes[ind:ind+char_len]))
		if string(str[ind:ind+char_len]) == charset {
			resArr[i] = string(str[start:ind])
			i++
			start = ind + char_len
		}
		if ind == str_len-char_len {
			// ind += char_len
			resArr[i] = string(str[start:])
		}
	}
	// resArr[i] = string(str[start:ind])
	return resArr
}
