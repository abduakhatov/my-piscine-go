package main

import (
	"fmt"
	piscine ".."
)

func main() {
	str := "HelloHAyou?"
	fmt.Println(piscine.Split(str, "HA"))
}
