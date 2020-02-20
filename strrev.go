package piscine

func StrRev(s string) string {
	res := []byte(s)
	for i := 0; i < len(s); i++ {
		res[i] = s[len(s)-i-1]
	}
	return string(res)
}
