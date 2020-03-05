package piscine

func Unmatch(arr []int) int {
	for i, iVal := range arr {
		count := 1
		for j, jVal := range arr {
			if iVal == jVal && i != j {
				count++
			}
		}
		if count%2 == 1 {
			return iVal
		}
	}
	return -1
}
