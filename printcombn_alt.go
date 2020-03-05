package piscine

// import "github.com/01-edu/z01"

// func printNum(n int) { // n=12
// 	if n == 0 {
// 		z01.PrintRune('0')
// 		return
// 	}
// 	if n >= 10 {
// 		printNum(n / 10)
// 	}
// 	w := '0'
// 	for i := 0; i < n%10; i++ {
// 		w++
// 	}
// 	z01.PrintRune(w)
// }

// func isValid(n int) bool {
// 	for n != 0 {
// 		if n >= 10 && n%10 <= (n/10)%10 {
// 			return false
// 		}
// 		n /= 10
// 	}
// 	return true
// }

// func isLastNum(numb, n int) bool {
// 	maxLen := 1
// 	for i := 1; i < n; i++ { // maxLen=10
// 		maxLen *= 10
// 	}
// 	if numb/maxLen == 10-n { // 0 == 8
// 		return true
// 	}
// 	return false
// }

// // PrintCombN_alt prints all possible combinations of n different digits in ascending order.
// func PrintCombN_alt(n int) { // n=2; maxLen=10
// 	maxLen := 1
// 	for i := 1; i < n; i++ {
// 		maxLen *= 10
// 	}
// 	for i := maxLen / 10; i <= maxLen*9; i++ { // i=12, i <= 90
// 		if isValid(i) {
// 			if i <= maxLen && maxLen != 1 {
// 				z01.PrintRune('0')
// 			}
// 			printNum(i)
// 			if !isLastNum(i, n) {
// 				z01.PrintRune(',')
// 				z01.PrintRune(' ')
// 			}
// 		}
// 	}
// 	z01.PrintRune('\n')
// }

// func main() {
// 	PrintCombN_alt(2)
// }
