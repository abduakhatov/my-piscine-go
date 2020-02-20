package piscine

func UltimateDivMod(a *int, b *int) {
	*a = *a / *b
	*b = *a % *b
}
