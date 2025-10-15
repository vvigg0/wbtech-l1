package main

import (
	"fmt"
	"sort"
)

func Grouping(temps []float64) map[int][]float64 {
	groups := make(map[int][]float64)
	for _, temp := range temps {
		key := int(temp/10) * 10 // к примеру 15.5| 15.5/10=1.55 -> int(1.55)=1 -> 1*10=10
		groups[key] = append(groups[key], temp)
	}
	return groups
}
func main() {
	temps := []float64{-15.5, -14.9, -10.1, -9.9, -0.01, 0.01, 9.99, 10.01, 19.99, 20.01, 25.5, 44.4}
	g := Grouping(temps)
	keys := make([]int, 0, len(g))
	for key := range g {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, key := range keys {
		fmt.Printf("%d: %v\n", key, g[key])
	}
}
