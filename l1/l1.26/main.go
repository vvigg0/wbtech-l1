package main

import (
	"fmt"
	"unicode"
)

func isUnique(str string) bool {
	m := make(map[rune]struct{})
	for _, l := range str {
		lowL := unicode.ToLower(l)
		if _, exist := m[lowL]; exist {
			return false
		}
		m[lowL] = struct{}{}
	}
	return true
}
func main() {
	str := "AbCdeF"
	fmt.Println(isUnique(str))

}
