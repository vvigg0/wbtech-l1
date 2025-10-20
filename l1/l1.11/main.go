package main

import "fmt"

func intersect(slices ...[]int) []int {
	if len(slices) == 0 {
		return nil
	}
	freq := make(map[int]int)
	for _, s := range slices {
		dataSlice := make(map[int]struct{})
		for _, v := range s {
			dataSlice[v] = struct{}{}
		}
		for v := range dataSlice {
			freq[v]++
		}
	}
	res := make([]int, 0)
	need := len(slices)
	for k, v := range freq {
		if v == need {
			res = append(res, k)
		}
	}
	return res
}
func main() {
	a := []int{1, 2, 3, 1, 1}
	b := []int{2, 3, 4}
	c := []int{2, 3, 4, 5, 6}
	set := intersect(a, b, c)
	fmt.Println(set)
}
