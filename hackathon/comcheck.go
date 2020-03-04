package hackathon

import (
	"fmt"
	"os"
)

func ComCheck() {
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
