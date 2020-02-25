package piscine

func Join(strs []string, sep string) string {
	res := strs[0]
	for _, val := range strs[1:] {
		res = res + sep + val
	}
	return res
}
