package piscine

func UltimateDivMod(a *int, b *int) {
	_a := *a / *b
	_b := *a % *b
	*a = _a
	*b = _b
}
