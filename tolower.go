package piscine

func ToLower(s string) string {
	str := []rune(s)
	for i := 0; i < len(str); i++ {
		c := str[i]
		if 'A' <= c && c <= 'Z' {
			c -= 'A' - 'a'
		}
		str[i] = c
	}
	return string(str)
}
