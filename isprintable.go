package piscine

func IsPrintable(str string) bool {
	if str == "" {
		return false
	}
	for _, s := range str {
		if !((s >= 65 && s <= 90) || (s >= 97 && s <= 122)) {
			return false
		}
	}
	return true
}
