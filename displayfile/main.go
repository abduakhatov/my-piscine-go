package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	counter := 0
	fileName := "quest8.txt"
	for _, val := range args {
		if val != fileName {
			fmt.Println("File name missing")
			return
		}
		counter++
	}
	fmt.Println((counter))
	if counter > 1 {
		fmt.Println("Too many arguments")
		return
	} else if counter == 0 {
		fmt.Println("File name missing")
		return
	}
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error")
	}
	arr := make([]byte, 24)
	file.Read(arr)
	fmt.Println(string(arr))
	fmt.Println()
}
