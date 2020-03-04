package hackathon

import "fmt"

func _Rot14(str string) string {
	res := ""
	capital := 0
	for _, val := range str {
		if 'A' <= val && val <= 'Z' {
			capital = 'A'
		}
		if 'a' <= val && val <= 'z' {
			if val+14 > 'z' {
				(val + capital + 14) % 26
				// todo finish
			}
		}
		res += val
		capital = 0
	}
}
