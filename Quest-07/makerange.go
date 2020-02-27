package piscine

func MakeRange(min, max int) []int {
	if max <= min {
		return []int{}
	}
	size := max-min
	res_arr := make([]int, size)
	for i := min; i < max; i++ {
		res_arr[i-min] = i
	}
	return res_arr
}
