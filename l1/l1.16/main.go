package main

import "fmt"

func QSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var left, middle, right []int
	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v == pivot {
			middle = append(middle, v)
		} else {
			right = append(right, v)
		}
	}
	result := make([]int, 0, len(arr))
	result = append(result, QSort(left)...)
	result = append(result, middle...)
	result = append(result, QSort(right)...)
	return result
}
func main() {
	fmt.Println(QSort([]int{3, 8, 3, 2, 5, 6, 4, 2, 2, 10, 4, 6, 5, 9, 1, 6, 6, 1, 1, 7}))
	// result = []int{1 1 1 2 2 2 3 3 4 4 5 5 6 6 6 6 7 8 9 10}
}
