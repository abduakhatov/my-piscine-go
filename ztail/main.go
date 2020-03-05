package main

import (
	piscine ".."
	"fmt"
	"os"
)

func ReadFile(offset int, fileName string) string {
	file, _ := os.Open("file.txt")
	defer file.Close()
	newOffset, _ := file.Seek(-1*int64(offset), 2)

	buffer := []byte{}
	for i := int64(0); i < newOffset; i++ {
		buffer = append(buffer, 0x0)
	}

	n, _ := file.ReadAt(buffer, newOffset)
	return string(buffer[:n])
}

func main() {
	args := os.Args[1:]
	finish := false
	hasOption := false
	hasOffset := false
	offset := 0
	fileName := ""

	for _, arg := range args {
		if arg == "-c" {
			hasOption = true
		} else if hasOption && !hasOffset {
			if isNumeric := piscine.IsNumeric(arg); isNumeric {
				offset = piscine.Atoi(arg)
				hasOffset = true
			} else {
				fmt.Printf("tail: invalid number of bytes: '%v'", arg)
				os.Exit(1)
			}
		} else {
			fileName = arg
		}
		finish = true
	}

	if finish && fileName != "" {
		if hasOption && !hasOffset {
			fmt.Printf("tail: option requires an argument -- 'c'\nTry 'tail --help' for more information.")
			os.Exit(1)
		}
		if _, err := os.Stat(fileName); os.IsNotExist(err) {
			fmt.Printf("tail: cannot open '" + fileName + "' for reading: No such file or directory")
			os.Exit(1)
		}
		result := ReadFile(offset, fileName)
		fmt.Printf(result)
		os.Exit(0)

	} else {
		os.Exit(1)
	}
}

// func readFromFile(file *os.File, offset, size int) ([]byte, error) {
// 	res := make([]byte, size, size)
// 	if _, err := file.ReadAt(res, int64(offset)); err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }

// func ReadFile(offset int, hasOffset bool, fileName string, ) {
// 	file, _ := os.Open(fileName)
// 	defer file.Close()
// 	newOffset, _ := file.Seek(-1 * int64(offset), 2)
// 	buffer := make([]byte, newOffset)
// 	n, _ := file.ReadAt(buffer, newOffset)
// 	// if err != nil && err != io.EOF {
// 	// 	panic(err)
// 	// }
// 	fmt.Println(string(buffer[:n]))
// }
