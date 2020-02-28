package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func PrintSolution(result [8]int) {
	for _, val := range result {
		z01.PrintRune(rune(val) + '0')
	}
}

func PlaceQueen(horiz [8]bool, f_diag [15]bool, b_diag [15]bool, result [8]int, loop int) bool {
	if loop == 8 {
		return false
	}
	for i := 0; i < 8; i++ {
		if horiz[i] || b_diag[loop+i] || f_diag[7-loop+i] {
			continue
		}
		horiz[loop] = true
		f_diag[7-i+loop] = true
		b_diag[i+loop] = true
		result[loop] = i + 1
		if loop == 8 {
			PrintSolution(result)
		}
		PlaceQueen(horiz, f_diag, b_diag, result, loop+1)
	}
	return false
}

func EightQueens() {
	fmt.Println("start")
	var horiz [8]bool
	var f_diag [15]bool
	var b_diag [15]bool
	var result [8]int
	PlaceQueen(horiz, f_diag, b_diag, result, 0)
}
