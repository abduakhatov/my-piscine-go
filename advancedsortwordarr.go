package piscine

func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	for i := 0; i < len(*array)-1; i++ {
		for j := i + 1; j < len(*array); j++ {
			if f((*array)[i], (*array)[j]) > 0 {
				tmp := (*array)[i]
				(*array)[i] = (*array)[j]
				(*array)[j] = tmp
			}
		}
	}
}
