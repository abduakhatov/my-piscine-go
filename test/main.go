package main

import (
	"fmt"
	piscine ".."
)

func main() {
	middle := piscine.Abort(2, 3, 8, 4, 7)
	fmt.Println(middle)
}
