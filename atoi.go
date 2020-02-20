package piscine

func Atoi(s string) int {
	_s := []byte(s)
	count := 0
	for range _s {
		count += 1
	}
	result := 0
	pow := 1
	tmp := 0
	sign := 1
	adj_chars := false
	for i := count - 1; i >= 0; i-- {
		tmp = (int(_s[i]) - 48)
		if tmp >= 0 && tmp <= 9 {
			result += (tmp * pow)
			pow *= 10
		} else if !adj_chars && ((tmp == -5 || tmp == -3) && i == 0) {
			adj_chars = true
			if tmp == -3 {
				sign = -1
			}
		} else {
			return 0
		}
	}
	return result * sign
}
