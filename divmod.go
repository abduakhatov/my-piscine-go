package piscine

func DivMod(a int, b int, div *int, mod *int) {
	res := a / b
	*div = res
	*mod = a % b
}
