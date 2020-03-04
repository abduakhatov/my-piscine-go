package piscine

func IsNumeric(str string) bool {
	if str == "" {
		return false
	}
	if str[0] == '-' || str[0] == '+' {
		str = str[1:]
	}
	for _, s := range str {
		if s < 48 || s > 57 {
			return false
		}
	}
	return true
}
