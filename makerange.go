package piscine

func MakeRange(min, max int) []int {
	var arr []int
	if max <= min {
		return arr
	}
	arr = make([]int, max-min)
	for i := min; i < max; i++ {
		arr[i-min] = i
	}
	return arr
}
