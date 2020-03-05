package hackathon

func Unmatch2(arr []int) int {
	length := 0
	for range arr {
		length++
	}
	for i := 0; i < length; i++ {
		hasPair := false
		for j := 1; j < length; j++ {
			if arr[i] == arr[j] {
				hasPair = true
				break
			}
		}
		if !hasPair {
			return arr[i]
		}
	}

}
