package piscine

// import "fmt"

func TrimAtoi(s string) int {
	isPlus := true
	hadNumber := false
	counter := 0
	resArr := []int{}
	for _, val := range s {
		if '0' <= val && val <= '9' {
			if val == '0' && !hadNumber {
				continue
			}
			resArr = append(resArr, int(val-48))
			hadNumber = true
			counter++
		} else if val == '-' && !hadNumber {
			isPlus = false
		}
	}
	res := 0
	op := 1
	for i := counter - 1; i >= 0; i-- {
		res += resArr[i] * op
		op *= 10
	}
	if !isPlus {
		res *= -1
	}
	return res
}
