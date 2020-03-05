package piscine

func SortByMax(table []int) {
	for i := 0; i < len(table)-1; i++ {
		for j := i + 1; j < len(table); j++ {
			if table[i] < table[j] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
}

func Max(arr []int) int {
	SortByMax(arr)
	return arr[0]
}
