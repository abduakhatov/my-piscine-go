package piscine

func LastRune(s string) rune {
	runes := []rune(s)

	counter := 0
	for range runes {
		counter += 1
	}

	return runes[counter-1]
}
