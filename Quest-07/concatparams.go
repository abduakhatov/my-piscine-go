package piscine

func ConcatParams(args []string) string {
	res := ""
	for _, val := range args {
		res = res + "\n" + val
	}
	res = res[1:]
	return res
}