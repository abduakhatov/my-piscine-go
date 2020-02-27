package piscine

func AppendRange(min, max int) []int {
	var arr_res []int
	if max <= min {
		return arr_res
	}
	for i := min; i < max; i++ {
		arr_res = append(arr_res, i)
	}
	return arr_res
}
