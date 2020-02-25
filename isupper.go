package piscine

func IsUpper(str string) bool {
	for _, s := range str {
		if s < 65 || s > 90 {
			return false
		}
	}
	return true
}
