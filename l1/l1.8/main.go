package main

import (
	"fmt"
)

func inverseBytes(val int64, poss ...int) (int64, string) {
	for _, pos := range poss {
		val ^= 1 << pos
	}
	return val, fmt.Sprintf("Стало после инверсии %v байта(ов): %b (%d)", poss, val, val)
}
func turnToOne(val int64, poss ...int) (int64, string) {
	for _, pos := range poss {
		val |= 1 << pos
	}
	return val, fmt.Sprintf("Стало после превращения в 1 %v байта(ов): %b (%d)", poss, val, val)
}
func turnToZero(val int64, poss ...int) (int64, string) {
	for _, pos := range poss {
		val &^= 1 << pos
	}
	return val, fmt.Sprintf("Стало после превращения в 0 %v байта(ов): %b (%d)", poss, val, val)
}
func main() {
	var val int64 = 11
	fmt.Printf("Изначально было %b (%d)\n", val, val)
	val, str := inverseBytes(val, 2)
	fmt.Println(str)
	val, str = turnToZero(val, 1, 3, 4)
	fmt.Println(str)
	_, str = turnToOne(val, 2, 3, 4)
	fmt.Println(str)
}
