package piscine

// import "fmt"

func IsPrintable(str string) bool {
	if str == "" {
		return false
	}
	for _, s := range str {
		// fmt.Println(">", s, string(s))
		if !(s >= 32) {
			return false
		}
	}
	return true
}
