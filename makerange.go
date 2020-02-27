package piscine

func MakeRange(min, max int) []int {
	if max <= min {
		return []int{}
	}
	arr := make([]int, max-min)
	for i := min; i < max; i++ {
		arr[i-min] = i
	}
	return arr
}
