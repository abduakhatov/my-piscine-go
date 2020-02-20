package piscine

func StrRev(s string) string {
	length := 0
	for range s {
		length += 1
	}
	res := []byte(s)
	for ind := range s {
		res[ind] = s[length-ind-1]
	}
	return string(res)
}
