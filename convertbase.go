package piscine

func ToBaseRec(nbr int, base []rune, size int) string {
	if nbr < size {
		return string(base[nbr])
	}
	return ToBaseRec(nbr/size, base, size) + string(base[nbr%size])
}

func ToBase(nbr int, base string) string {
	size := 0
	for range base {
		size++
	}
	return ToBaseRec(nbr, []rune(base), size)
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	number := AtoiBase(nbr, baseFrom)
	return ToBase(number, baseTo)
}
