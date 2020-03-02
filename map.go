package piscine

func Map(f func(int) bool, arr []int) []bool {
	counter := 0
	for range arr {
		counter++
	}
	res_arr := make([]bool, counter)
	for i, val := range arr {
		res_arr[i] = f(val)
	}
	return res_arr
}
