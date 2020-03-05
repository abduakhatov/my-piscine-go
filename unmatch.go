package piscine

func Unmatch(arr []int) int {
	length := 0
	for range arr {
		length++
	}
	for i := 0; i < length; i++ {
		hasPair := false
		for j := 0; j < length; j++ {
			if arr[i] == arr[j] && i != j {
				hasPair = true
				break
			}
		}
		if !hasPair {
			return arr[i]
		}
	}
	return -1
}
