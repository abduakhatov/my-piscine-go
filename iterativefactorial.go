package piscine

import "math"

func IterativeFactorial(nb int) int {
	max := math.MaxInt32
	if nb < 1 || nb > max {
		return 0
	}
	result := 1
	for i := 1; i <= nb; i++ {
		if result > max {
			return 0
		}
		result *= i
	}
	return result
}
