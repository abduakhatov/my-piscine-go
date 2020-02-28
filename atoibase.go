package piscine

func FromBase(s string, base string) int {
	
}

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
	return 0
}
