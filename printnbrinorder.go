package piscine

import "github.com/01-edu/z01"

func SortTable(table []int) []int {
	for i := 0; i < len(table)-1; i++ {
		for j := i + 1; j < len(table); j++ {
			if table[i] > table[j] {
				tmp := table[i]
				table[i] = table[j]
				table[j] = tmp
			}
		}
	}
	return table
}

func ToIntArray(n int) []int {
	if n == 0 {
		return []int{}
	}
	digit := int(n % 10)
	next := int(n / 10)
	return append(ToIntArray(next), digit)
}

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	sorted_arr := SortTable(ToIntArray(n))
	for i := range sorted_arr {
		z01.PrintRune(rune(i) + 49)
	}
}
