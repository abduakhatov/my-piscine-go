package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	matches := []string{"01", "galaxy", "galaxy 01"}
	for _, val := range args {
		for _, jVal := range matches {
			if val == jVal {
				fmt.Println("Alert!!!")
				return
			}
		}
	}
}
