package piscine

func SortWordArr(array []string) {
	length := 0
	for range array {
		length++
	}
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if array[i] > array[j] {
				tmp := array[i]
				array[i] = array[j]
				array[j] = tmp
			}
		}
	}
}
