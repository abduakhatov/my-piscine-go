package main

import (
	"fmt"
	piscine ".."
)

func main() {
	str := "HelloHAhowHAareHAyou?HA"
	fmt.Println(piscine.Split(str, "HA"))
}
