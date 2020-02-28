package piscine

func ConcatParams(args []string) string {
	res := ""
	for _, val := range args {
		res += "\n" + val
	}
	return res[1:]
}
