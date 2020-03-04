package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// ./student/displayfile/quest8.txt
	args := os.Args[1:]
	count := 0
	for range args {
		count++
	}
	if count == 0 {
		fmt.Println("File name missing")
		return
	} else if count > 1 {
		fmt.Println("Too many arguments")
		return
	}
	fileName := args[0]

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println()
		return
	}
	fmt.Println(string(content))
}
