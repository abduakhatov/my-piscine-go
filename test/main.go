package main

import piscine ".."

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	piscine.ForEach(piscine.PrintNbr, arr)
}
