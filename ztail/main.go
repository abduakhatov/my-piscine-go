package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// func PrintResult(str string) {
// 	for _, val := range str {
// 		z01.PrintRune(val)
// 	}
// }

// func MyReadFile(fileName string) string {
// 	content, err := ioutil.ReadFile(fileName)
// 	if err != nil {
// 		return "error"
// 	}
// 	return string(content)
// }

// func main() {
// 	args := os.Args[1:]
// 	// abc := "1"
// 	finish := false
// 	for _, fileName := range args {
// 		if _, err := os.Stat(fileName); err != nil {
// 			PrintResult("open " + fileName + ": no such file or directory\n")
// 			return
// 		}
// 		PrintResult(MyReadFile(fileName))
// 		finish = true
// 	}
// 	if !finish {
// 		reader := io.TeeReader(os.Stdin, os.Stdout)
// 		ioutil.ReadAll(reader)
// 		os.Stdin.Close()
// 		os.Stdout.Close()
// 	}
// }

// func readFromFile(file *os.File, offset, size int) ([]byte, error) {
// 	res := make([]byte, size, size)
// 	if _, err := file.ReadAt(res, int64(offset)); err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }

func main() {
	// res := make([]byte, 100, size)
	file, _ := os.Open("file.txt")
	var res []byte
	var res2 []byte
	if _, err := file.Read(res); err != nil {

	}
	if _, err := file.ReadAt(res2, int64(100)); err != nil {

	}
	fmt.Println(res)
	fmt.Println(file)
	fmt.Println(res)
	data, _ := ioutil.ReadAll(file)
	fmt.Println(string(data))

}
