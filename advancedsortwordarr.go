package piscine

func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	length := 0
	for range array {
		length++
	}
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if f(array[i], array[j]) > 0 {
				tmp := array[i]
				array[i] = array[j]
				array[j] = tmp
			}
		}
	}
}
