package main

import (
	// "fmt"
	"os"
)

func PrintConsole(str string) {
	os.Stdout.WriteString(str)
	os.Stdout.WriteString("\n")
	os.Stdout.Close()
}

func NbrToStrRec(n, dot int64) string {
	if 10 > n && n > -1*10 {
		return string('0' + n*dot)
	}
	return NbrToStrRec(n/10, dot) + string('0'+(n%10)*dot)
}

func NbrToStr(n int64) string {
	dot := int64(1)
	res := ""
	if n == 0 {
		return "0"
	}
	if n < 0 {
		dot *= -1
		res += "-"
	}
	return res + NbrToStrRec(n, dot)
}

func AtoiOverflow(a, b, c int64) bool {
	if a < 0 && c < 0 {
		return a*b+c < 0
	} else if a > 0 && c > 0 {
		return a*b+c > 0
	}
	return true
}

func MultiplyOverflow(a, b, c int64) bool {
	prod := a*b + c
	return (prod/b)-c == a
}

func PlusOverflow(a, b int64) bool {
	if a < 0 && b < 0 {
		return a+b < 0
	} else if a > 0 && b > 0 {
		return a+b > 0
	}
	return true
}

func MinusOverflow(a, b int64) bool {
	if a < 0 && b < 0 {
		if b <= a {
			return a-b >= 0
		}
		return a-b < 0
	} else if a > 0 && b > 0 {
		if a <= b {
			return a-b <= 0
		}
		return a-b > 0
	}
	return true
}

func Atoi(nbr string) (int64, bool) {
	var res int64 = 0
	var sign int64 = 1
	if nbr[0] == '-' {
		nbr = nbr[1:]
		sign *= -1
	} else if nbr[0] == '+' {
		nbr = nbr[1:]
	}
	for _, digit := range nbr {
		tmp := int64(digit-'0') * sign
		if !AtoiOverflow(res, int64(10), tmp) {
			return 0, false
		}
		res = res*10 + tmp
	}
	return res, true
}

func Plus(a, b string) {
	aa, aBool := Atoi(a)
	if !aBool {
		PrintConsole("0")
		return
	}
	bb, bBool := Atoi(b)
	if !bBool {
		PrintConsole("0")
		return
	}
	if !PlusOverflow(aa, bb) {
		PrintConsole("0")
		return
	}
	PrintConsole(NbrToStr(aa + bb))
}

func Deduct(a, b string) {
	aa, aBool := Atoi(a)
	if !aBool {
		PrintConsole("0")
		return
	}
	bb, bBool := Atoi(b)
	if !bBool {
		PrintConsole("0")
		return
	}
	if !MinusOverflow(aa, bb) {
		PrintConsole("-0")
		return
	}
	PrintConsole(NbrToStr(aa - bb))
}

func Devide(a, b string) {
	bb, bBool := Atoi(b)
	if !bBool {
		PrintConsole("0")
		return
	}
	if bb == 0 {
		PrintConsole("No division by 0")
		return
	}
	aa, aBool := Atoi(a)
	if !aBool {
		PrintConsole("0")
		return
	}
	PrintConsole(NbrToStr(aa / bb))
}

func Multiply(a, b string) {
	aa, aBool := Atoi(a)
	if !aBool {
		PrintConsole("0")
		return
	}
	bb, bBool := Atoi(b)
	if !bBool {
		PrintConsole("0")
		return
	}
	if !MultiplyOverflow(aa, bb, 0) {
		PrintConsole("0")
		return
	}
	PrintConsole(NbrToStr(aa * bb))
}

func Mod(a, b string) {
	bb, bBool := Atoi(b)
	if !bBool {
		PrintConsole("0")
		return
	}
	if bb == 0 {
		PrintConsole("No modulo by 0")
		return
	}
	aa, aBool := Atoi(a)
	if !aBool {
		PrintConsole("0")
		return
	}
	PrintConsole(NbrToStr(aa % bb))
}

func IsNumeric(str string) bool {
	if str == "" {
		return false
	}
	if str[0] == '-' || str[0] == '+' {
		str = str[1:]
	}
	for _, s := range str {
		if s < 48 || s > 57 {
			return false
		}
	}
	return true
}

func main() {
	args := os.Args[1:]
	argsCount := 0
	for range args {
		argsCount++
	}
	if argsCount != 3 {
		return
	}
	if !(IsNumeric(args[0]) && IsNumeric(args[2])) {
		PrintConsole("0")
	}
	funcsArr := []func(string, string){Plus, Deduct, Devide, Multiply, Mod}
	operators := []string{"+", "-", "/", "*", "%"}
	for i, val := range operators {
		if val == args[1] {
			funcsArr[i](args[0], args[2])
			return
		}
	}
	PrintConsole("0")
}
