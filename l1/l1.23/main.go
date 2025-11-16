package main

import "fmt"

func deleteElem[T any](s []T, i int) []T {
	if i < 0 || i >= len(s) {
		panic("index out of range")
	}
	copy(s[i:], s[i+1:])
	var zero T
	s[len(s)-1] = zero
	return s[:len(s)-1]
}
func main() {
	s := []int{}
	for i := 0; i < 10; i++ {
		x := i
		s = append(s, x)
	}
	fmt.Println(s)
	newS := deleteElem(s, 3)
	fmt.Println(newS)
}
