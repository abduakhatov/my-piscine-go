package main

import (
	"fmt"
	"os"
)

func IsVertOK(table [9][9]byte, val, i byte) bool {
	for j := '1'; j <= '9'; j++ {
		if table[i][j] == val {
			return false 
		}
	}
	return true
}

func IsHorOK(table [9][9]byte, val, j byte) bool {
	for i := '1'; i <= '9'; i++ {
		if table[i][j] == val {
			return false 
		}
	}
	return true
}

func IsBoxOK(table [9][9]byte, val, i, j byte) bool {
	for ii := '1'; ii <= '3'; ii++ {
		for jj := '1'; jj <= '3'; jj++ {
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
		fmt.Println("Error\n")
		return
	}
	var candidates [9][9][9]byte{}
	var table [9][9]byte{}
	// todo fill in the table
	for j_row, b := range args {
		for i_col, a := range []byte(b) {
			if b == '.' {
				continue
			}
			table[j_row][i_col] = a
		}	
	}
	fmt.Println(table)
	return

	// Candidate Filling
	for i := '1'; i <= '9'; i++ {
		for j := '1'; j <= '9'; j++ {
			if table[i][j] == 0 {
				continue
			} 
			for k := '1'; k <= '9'; k++ {
				if IsBoxOK(table, k, i, j) && IsHorOK(table, k, j) && IsVertOK(table, k, i) {
					candidates[i][j][k-1] = k
				}
			}
		}
	}
}