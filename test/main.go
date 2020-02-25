package main

import (
	"fmt"
	piscine ".."
)

func main() {
	fmt.Println("Hello!", "l" , " 2 - "    ,piscine.Index("Hello!", "l"))
	fmt.Println("Salut!", "alu" , " 1 - "    ,piscine.Index("Salut!", "alu"))
	fmt.Println("Ola!", "hOl" , " -1 - "    ,piscine.Index("Ola!", "hOl"))
	fmt.Println("", "hOl" , " -1 - "    ,piscine.Index("", "hOl"))
	fmt.Println("Ola!", "" , " 0 - "    ,piscine.Index("Ola!", ""))
	fmt.Println("", "" , " 0 - "    ,piscine.Index("", ""))
}
