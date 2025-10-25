package main

import "fmt"

func GenSet(data []string) []string {
	m := make(map[string]struct{}) // мапа, которая будет хранить уникальные значения
	for _, v := range data {
		m[v] = struct{}{}
	}
	// dataMap = map[cat:{} dog:{} tree:{}]
	result := make([]string, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}
func main() {
	var data []string = []string{"cat", "cat", "dog", "cat", "tree"}
	set := GenSet(data)
	fmt.Println(set)
}
