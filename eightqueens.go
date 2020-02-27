package piscine

import "github.com/01-edu/z01"

// import "fmt"

func PrintSolution(result [8]int) {
	for _, val := range result {
		z01.PrintRune(rune(val) + '0')
	}
}

func PlaceQueen(horiz [8]bool, f_diag [14]bool, b_diag [14]bool, result [8]int, loop int) bool {
	if loop == 8 {
		return false
	}
	for i := 0; i < 8; i++ {
		if !(horiz[i] && ((i+loop <= 7 && b_diag[loop] && f_diag[loop+i]) || (i+loop > 7 && b_diag[loop] && f_diag[loop+i]))) {
			continue
		}
		horiz[loop] = true
		f_diag[loop] = true
		b_diag[loop+i] = true
		result[loop] = i
		if loop == 8 {
			PrintSolution(result)
		}
		PlaceQueen(horiz, f_diag, b_diag, result, loop-1)
	}
	return false
}

func EightQueens() {
	var horiz [8]bool
	var f_diag [14]bool
	var b_diag [14]bool
	var result [8]int
	PlaceQueen(horiz, f_diag, b_diag, result, 0)
}
