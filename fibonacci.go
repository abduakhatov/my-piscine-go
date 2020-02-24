package piscine

func Fibonacci(index int) int {
	return SubFibonacci(index, 0, 1)
}

func SubFibonacci(index, st, nd int) int {
	if index <= 1 {
		return nd
	}
	return SubFibonacci(index-1, nd, st+nd)
}
