package piscine

func Sqrt(nb int) int {
	if nb == 1 {
		return 1
	}
	for i := nb / 2; i >= 1; i-- {
		tmp := i * i
		if tmp == nb {
			return i
		}
	}
	return 0
}
