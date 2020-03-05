package main

import (
	"fmt"

	piscine ".."
)

const N = 6

func main() {
	arr := make([]string, N)
	arr[0] = "a"
	arr[2] = "b"
	arr[4] = "c"

	for _, v := range arr {
		fmt.Println(v)
	}

	fmt.Println("Size after compacting:", piscine.Compact(&arr))

	for _, v := range arr {
		fmt.Println(v)
	}
}
