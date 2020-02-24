package piscine

func SubRecursiveFactorial(nb, index int) int {
	if index > 21 {
		return 0
	} else if nb == 1 {
		return 1
	}
	return SubRecursiveFactorial(nb-1, index+1) * nb
}

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	return SubRecursiveFactorial(nb, 1)
}
