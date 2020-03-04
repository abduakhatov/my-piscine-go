package piscine

func SortIntTable(table *[]int, length int) {
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if (*table)[i] > (*table)[j] {
				tmp := (*table)[i]
				(*table)[i] = (*table)[j]
				(*table)[j] = tmp
			}
		}
	}
}

func Abort(a, b, c, d, e int) int {
	table := []int{a, b, c, d, e}
	length := 0
	for range table {
		length++
	}
	SortIntTable(&table, length)
	return table[2]
}
