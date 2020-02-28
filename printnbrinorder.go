package piscine

import "github.com/01-edu/z01"

func SortTable(table []int, counter int) []int {
	for i := 0; i < counter-1; i++ {
		for j := i + 1; j < counter; j++ {
			if table[i] > table[j] {
				table[i], table[j] = table[j], table[i]
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
	} else if n == 0 {
		z01.PrintRune('0')
	}
	arr := ToIntArray(n)
	counter := 0
	for range arr {
		counter++
	}
	sorted_arr := SortTable(arr, counter)
	for _, v := range sorted_arr {
		z01.PrintRune(rune(v) + '0')
	}
}
