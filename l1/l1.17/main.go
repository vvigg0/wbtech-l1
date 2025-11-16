package main

import (
	"fmt"
	"sort"
)

func bSearchRecurs(arr []int, v int) int {
	if len(arr) == 0 {
		return -1
	}
	mid := len(arr) / 2
	if v < arr[mid] {
		return bSearchRecurs(arr[:mid], v)
	} else if v > arr[mid] {
		i := bSearchRecurs(arr[mid+1:], v)
		if i == -1 {
			return -1
		}
		return mid + 1 + i

	}
	return mid
}
func bSearchIterate(arr []int, v int) int {
	first, last := 0, len(arr)-1
	for first <= last {
		mid := first + (last-first)/2
		if arr[mid] == v {
			return mid
		}
		if arr[mid] < v {
			first = mid + 1
		} else {
			last = mid - 1
		}
	}
	return -1
}
func main() {
	val := 11
	arr := []int{3, 8, 3, 2, 5, 6, 4, 2, 2, 10, 4, 6, 6, 9, 1, 6, 6, 1, 1, 7}
	sort.Ints(arr)
	i := bSearchRecurs(arr, val)
	if i < 0 {
		fmt.Println("Значение не найдено")
	} else {
		fmt.Printf("Значение %v найдено по индексу %v\n", val, i)
	}
	i = bSearchIterate(arr, val)
	if i < 0 {
		fmt.Println("Значение не найдено")
	} else {
		fmt.Printf("Значение %v найдено по индексу %v", val, i)
	}
}
