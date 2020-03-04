package hackathon

func CollatzCountRec(start int) int {
	if start == 1 {
		return 1
	}
	count := 1
	if start%2 == 0 {
		count = CollatzCountRec(start / 2)
	} else {
		count = CollatzCountRec((start * 3) + 1)
	}
	count++
	return count
}

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	return CollatzCountRec(start)
}
