package piscine

func BasicJoin(strs []string) string {
	res := ""
	for _, val := range strs {
		res = res + val
	}
	return res
}
