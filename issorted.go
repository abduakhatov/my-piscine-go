package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	length := 0
	for range tab {
		length++
	}
	sort := 0
	for i := 0; i < length-1; i++ {
		isSorted := f(tab[i], tab[i+1])
		if isSorted > 0 {
			sort++
		} else if isSorted < 0 {
			sort--
		} else {
			length--
		}
	}
	length--
	if sort == length || sort == -1*length {
		return true
	}
	return false
}
