package piscine

func IsNumeric(str string) bool {
	// if str == "" {
	// 	return false
	// }
	for _, s := range str {
		if s < 48 && s > 57 {
			return false
		}
	}
	return true
}
