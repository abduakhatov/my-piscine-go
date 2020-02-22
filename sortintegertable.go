package piscine

func SortIntegerTable(table []int) {
	for i := 0; i < len(table)-1; i++ {
		for j := i + 1; j < len(table); j++ {
			if table[i] > table[j] {
				tmp := table[i]
				table[i] = table[j]
				table[j] = tmp
			}
		}
	}
	// fmt.Println(table)
}
