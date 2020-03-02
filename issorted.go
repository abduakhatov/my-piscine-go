package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	length := 0
	for range tab {
		length++
	}
	for i := 0; i < length-1; i++ {
		if f(tab[i], tab[i+1]) == 1 {
			return false
		}
	}
	return true
}
