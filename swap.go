package piscine

func Swap(a *int, b *int) {
	_a := *a
	*a = *b
	*b = _a
}
