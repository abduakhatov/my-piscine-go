package main

import (
	"fmt"
	piscine ".."
)

func main() {
	arr := []int{1, 2, 3, 1, 2, 3, 4}
	unmatch := piscine.Unmatch(arr)
	fmt.Println(unmatch)


	arr = []int{1, 1, 2, 4, 3, 4, 2, 3, 4}
	unmatch = piscine.Unmatch(arr)
	fmt.Println(unmatch)

}
