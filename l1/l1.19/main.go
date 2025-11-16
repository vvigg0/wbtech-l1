package main

import "fmt"

func ReverseStr(str string) string {
	if len(str) == 0 || len(str) == 1 {
		return str
	}
	r := []rune(str)
	left := 0
	right := len(r) - 1
	for left < right {
		r[left], r[right] = r[right], r[left]
		left++
		right--
	}
	return string(r)
}
func main() {
	fmt.Println(ReverseStr("главрыба"))
}
