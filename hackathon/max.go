package hackathon

func Max(arr []int) int {
	length := 0
	for range arr {
		length++
	}
	SortIntTable(&arr)
	return true //arr[length-1]

}
