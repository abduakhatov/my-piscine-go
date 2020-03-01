package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	counter := 0
	file := ""
	for range args {
		counter++

	}
	if counter > 1 {
		fmt.Println("Too many arguments")
	}

	file, err := os.Open(file)
	if err != nil {
		fmt.Println("Error")
	}
	var arr [24]byte
	file.Read(arr)
	fmt.Println(string(arr))

}
