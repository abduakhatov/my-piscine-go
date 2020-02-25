package main

import (
	"fmt"
	piscine ".."
)

func main() {
	fmt.Println("DPIHV+Gl#p*Sg", "HV+Gl#p*S" , " 3 - "    ,piscine.Index("DPIHV+Gl#p*Sg", "HV+Gl#p*S"))
	fmt.Println("AeAe!v{E=?Xbv", "Ae!v{E=?Xb" , " 2 - "    ,piscine.Index("AeAe!v{E=?Xbv", "Ae!v{E=?Xb"))
	fmt.Println("Hello!", "l" , " 2 - "    ,piscine.Index("Hello!", "l"))
	fmt.Println("Salut!", "alu" , " 1 - "    ,piscine.Index("Salut!", "alu"))
	fmt.Println("Ola!", "hOl" , " -1 - "    ,piscine.Index("Ola!", "hOl"))
	fmt.Println("", "hOl" , " -1 - "    ,piscine.Index("", "hOl"))
	fmt.Println("Ola!", "" , " 0 - "    ,piscine.Index("Ola!", ""))
	fmt.Println("", "" , " 0 - "    ,piscine.Index("", ""))

	// a := "AeAe!v{E=?Xbv"
	// for ind, val := range a {
	// 	fmt.Println(byte(a[ind]), string(a[ind]), rune(val))
	// }
	// b := "Ae!v{E=?Xb"
	// fmt.Println("\n")
	// for ind, val := range b {
	// 	fmt.Println(byte(b[ind]), string(b[ind]), rune(val))
	// }
}
