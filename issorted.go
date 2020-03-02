package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	length := 0
	for range tab {
		length++
	}
	cap := length
	sort := 0
	for i := 0; i < length-1; i++ {
		isSorted := f(tab[i], tab[i+1])
		if isSorted > 0 {
			sort++
		} else if isSorted < 0 {
			sort--
		} else {
			cap--
		}
	}
	cap--

	if sort == cap || sort == -1*cap {
		return true
	}
	return false
}
