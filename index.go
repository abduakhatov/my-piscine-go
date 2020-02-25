package piscine

func Length(str string) int {
	counter := 0
	for range str {
		counter += 1
	}
	return counter
}

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	ss := []rune(s)
	bb := []rune(toFind)
	length := Length(toFind)
	ind := -1
	counter := 0
	for i := 0; i < Length(s); i++ {
		if ss[i] == bb[counter] {
			if counter == 0 {
				ind = i
			}
			counter++
			if counter == length {
				return ind
			}
		} else {
			if counter > 0 {
				i--
			}
			counter = 0
			ind = -1
		}
	}
	return ind
}
