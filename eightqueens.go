package piscine

import (
	"github.com/01-edu/z01"
)

func PrintSolution(result [8]int) {
	for _, val := range result {
		z01.PrintRune(rune(val) + '0')
	}
	z01.PrintRune('\n')
}

// PlaceQueen is
func PlaceQueen(horiz [8]bool, fDiag [15]bool, bDiag [15]bool, result [8]int, col int) {
	if col == 8 {
		return
	}
	for row := 0; row <= 7; row++ {
		if horiz[row] || bDiag[row+col] || fDiag[7-row+col] {
			continue
		}
		horizTmp := horiz
		fDiagTmp := fDiag
		bDiagTmp := bDiag
		resulTmp := result

		horizTmp[row] = true
		fDiagTmp[7-row+col] = true
		bDiagTmp[row+col] = true
		resulTmp[col] = row + 1
		if col == 7 {
			PrintSolution(resulTmp)
			return
		}
		PlaceQueen(horizTmp, fDiagTmp, bDiagTmp, resulTmp, col+1)
	}
}

func EightQueens() {
	var horiz [8]bool
	var fDiag [15]bool
	var bDiag [15]bool
	var table [8]int
	PlaceQueen(horiz, fDiag, bDiag, table, 0)
}
