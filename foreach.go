package piscine

func ForEach(f func(int), arr []int) {
	for _, val := range arr {
		f(val)
	}
}
