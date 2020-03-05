package main

import (
	"fmt"
	piscine ".."
)

func main() {
	nbits := piscine.ActiveBits(255)
	fmt.Println(nbits)
}
