package piscine

func Capitalize(s string) string {
	str := []rune(s)
	isReset := true
	for ind, val := range str {
		if isReset {
			if ('A' <= val && val <= 'Z') || ('0' <= val && val <= '9') {
				isReset = false
				continue
			} else if 'a' <= val && val <= 'z' {
				str[ind] -= 'a' - 'A'
				isReset = false
			}
		} else if !isReset && 'A' <= val && val <= 'Z' {
			str[ind] -= 'A' - 'a'
		} else if !(('a' <= val && val <= 'z') || ('A' <= val && val <= 'Z') || ('0' <= val && val <= '9')) {
			isReset = true
		}
	}
	return string(str)
}
