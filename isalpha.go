package piscine

func IsAlpha(str string) bool {
	hasAlpha := false
	for _, s := range str {
		if (s >= 48 && s <= 57) || (s >= 65 && s <= 90) || (s >= 97 && s <= 122) {
			if s >= 65 && s <= 90 {
				hasAlpha = true
			}
		} else {
			return false
		}
	}
	return hasAlpha
}
