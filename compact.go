package piscine

func Compact(ptr *[]string) int {
	res := *ptr
	count := 0
	for _, val := range *ptr {
		if val != "" {
			res[count] = val
			count++
		}
	}
	*ptr = res[0:count]
	return count
}
