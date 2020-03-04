package piscine

func Rot14(str string) string {
	res := []rune(str)
	for i, val := range str {
		if !(('A' <= val && val <= 'Z') || ('a' <= val && val <= 'z')) {
			continue
		}
		cap := rune(0)
		if 'A' <= val && val <= 'Z' {
			cap = 'a' - 'A'
		}
		if val+cap+14 > 'z' {
			res[i] = 'a' + (val + cap + 14 - 'z') - cap - 1
			continue
		}
		res[i] = val + 14
	}
	return string(res)
}
