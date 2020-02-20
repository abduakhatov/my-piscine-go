package main

import "fmt"

func BasicAtoi2(s string) int {
	_s := []byte(s)
	count := 0
	for range _s {
		count += 1
	}
	result := 0
	pow := 1
	tmp := 0
	for i := count - 1; i >= 0; i-- {
		tmp = (int(_s[i]) - 48)
		if tmp < 0 || tmp > 9 {
			return 0
		}
		result += (tmp * pow)
		pow *= 10
	}
	return result
}
