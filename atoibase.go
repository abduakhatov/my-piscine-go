package piscine

func GetCharIndex(ch rune, base []rune) int {
	for ind, val := range base {
		if val == ch {
			return ind
		}
	}
	return -1
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
	ss := []rune(s)
	bb := []rune(base)
	sSize := get_length(s)
	baseNumb := 1
	for i := 1; i < sSize; i++ {
		baseNumb = baseNumb * size
	}
	res := 0
	for _, val := range ss {
		res += baseNumb * GetCharIndex(val, bb)
		baseNumb /= size
	}
	return res
}
