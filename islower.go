package piscine

func IsLower(str string) bool {
	for _, s := range str {
		if s < 97 || s > 122 {
			return false
		}
	}
	return true
}
