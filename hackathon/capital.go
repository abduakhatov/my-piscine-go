package hackathon

func SortIntTable(table *[]int) {
	for i := 0; i < len(*table)-1; i++ {
		for j := i + 1; j < len(*table); j++ {
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
	SortIntTable(table)
	return table[2]
}
