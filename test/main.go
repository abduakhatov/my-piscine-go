package main

import (
	"fmt"
	piscine ".."
)

func main() {
	arr := []int{1, 2, 3, 4, 1, 2, 3}
	unmatch := piscine.Unmatch(arr)
	fmt.Println(unmatch)
}
