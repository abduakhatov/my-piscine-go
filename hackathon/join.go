package hackathon

func Join(strs []string, sep string) string {
	res := ""
	for _, val := range strs {
		res += val + sep
	}
	return res[1:]

}
