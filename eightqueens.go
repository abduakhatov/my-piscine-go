package piscine

import "github.com/01-edu/z01"
// import "fmt"

// func PrintSolution(result []int) {
// 	for _, val := range result {
// 		z01.PrintRune(rune(val)+'0')
// 	}
// }


// func PlaceQueen(horiz []bool, f_diag []bool, b_diag []bool, result []int, loop int) bool {
// 	for i := 0; i < 8; i++ {
// 		if horiz[i] && ((i+loop <= 7 && b_diag[loop] && f_diag[loop+i]) || (i+loop > 7 && b_diag[loop] && f_diag[loop]) ) {
// 			return false
// 		}
// 		horiz[loop] = true
// 		f_diag[loop] = true
// 		b_diag[loop] = true
// 		PlaceQueen(horiz, f_diag, b_diag, result, loop-1)
// 	}
// 	return true
// }

// func EightQueens() {
// 	var horiz [8]bool
// 	var f_diag [14]bool
// 	var b_diag [14]bool
// 	var result [8]int
// 	PlaceQueen(horiz, f_diag, b_diag, result, 0)
// }
