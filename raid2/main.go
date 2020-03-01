package main

import (
	"fmt"
	"os"
)

// IsVertOK Checks Conflicts within vertical line
func IsVertOK(table [9][9]int8, val, i int8) bool {
	for j := int8(1); j <= 9; j++ {
		if table[i][j] == val {
			return false
		}
	}
	return true
}

// IsHorizOK Checks Conflicts within horizontal line
func IsHorizOK(table [9][9]int8, val, j int8) bool {
	for i := int8(1); i <= 9; i++ {
		if table[i][j] == val {
			return false
		}
	}
	return true
}

// IsBoxOK Checks Conflicts within 3x3 Box
func IsBoxOK(table [9][9]int8, val, i, j int8) bool {
	for ii := int8(1); ii <= 3; ii++ {
		for jj := int8(1); jj <= 3; jj++ {
			if table[ii+i][jj+j] == val {
				return false
			}
		}
	}
	return true
}

func main() {
	args := os.Args[1:]
	if len(args) != 9 {
		fmt.Println("+Error\n")
		return
	}
	var candidates [9][9][9]int8
	var table [9][9]int8
	// Fill in the table
	for jRow, b := range args {
		for iCol, a := range b {
			if a == '.' {
				continue
			}
			if '9' < a || a <= '0' {
				fmt.Println("-Error\n")
				return
			}
			table[jRow][iCol] = int8(a) - '0'
		}
		fmt.Println(table[jRow])
	}
	fmt.Println(">>> Candidates")
	// Candidate Filling
	var j int8 = 1
	var k int8 = 1
	for i := int8(1); i <= 9; i++ {
		for j = 1; j <= 9; j++ {
			if table[i][j] != 0 {
				continue
			}
			for k = 1; k <= 9; k++ {
				if IsBoxOK(table, k, i, j) && IsHorizOK(table, k, j) && IsVertOK(table, k, i) {
					candidates[i][j][k-1] = k
				}
			}
		}
	}
}
