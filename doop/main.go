package main

import (
	// "fmt"
	piscine ".."
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

func MultplyOverflow(a, b, c int64) bool {
	if a < 0 && c < 0 {
		return a*b+c < 0
	} else if a > 0 && c > 0 {
		return a*b+c > 0
	} else if c == 0 {
		return a*b > 0
	}
	return true
}

func PlusOverflow(a, b int64) bool {
	// if a < 0 && c < 0 {
	// 	return  a * b + c < 0
	// } else if a > 0 && c > 0 {
	// 	return a * b + c > 0
	// }
	return true
}

func MinusOverflow(a, b int64) bool {
	// if a < 0 && c < 0 {
	// 	return  a * b + c < 0
	// } else if a > 0 && c > 0 {
	// 	return a * b + c > 0
	// }
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
		if !MultplyOverflow(res, int64(10), tmp) {
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

func Minus(a, b string) {

}

func Devide(a, b string) {
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
	if !MultplyOverflow(aa, bb, 0) {
		PrintConsole("0")
		return
	}
	PrintConsole(NbrToStr(aa + bb))
}

func Mod(a, b string) {
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
	PrintConsole(NbrToStr(aa % bb))
}

func main() {
	args := os.Args[1:]
	argsCount := 0
	for range args {
		argsCount++
	}
	if argsCount != 3 {
		PrintConsole("0")
		return
	}
	if !(piscine.IsNumeric(args[0]) || piscine.IsNumeric(args[1])) {
		PrintConsole("0")
	}
	funcsArr := []func(string, string){Plus, Minus, Devide, Multiply, Mod}
	operators := []string{"+", "-", "/", "*", "%"}
	for i, val := range operators {
		if val == args[1] {
			funcsArr[i](args[0], args[2])
			break
		}
	}
}
