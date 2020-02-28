package piscine

// func FromBase(s string, base []rune, size int) int {

// }

// func GetCharIndex(ch rune, base []rune) int {

// }

func AtoiBase(s string, base string) int {
	size := 0
	isValid := true
	for _, val := range base {
		if val == '-' || val == '+' {
			isValid = false
			size = 0
			break
		}
		size++
	}
	if size < 2 {
		return 0
	}
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if base[j] == base[i] {
				isValid = false
				break
			}
		}
	}
	if !isValid {
		return 0
	}
	// ss := []rune(s)
	// bb := []rune(s)

	// for ind, val := range ss {
	// 	char_ind := GetCharIndex(val, bb)
	// }

	return 0
}
