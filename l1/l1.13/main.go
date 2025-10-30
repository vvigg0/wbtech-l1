package main

import "fmt"

func reset() (int, int) {
	a := 8
	b := 5
	return a, b
}
func main() {
	// XOR обмен
	a, b := reset()
	fmt.Printf("Первый вариант:\na=%v b=%v\n", a, b)
	a ^= b
	b ^= a
	a ^= b
	fmt.Printf("a=%v b=%v\n", a, b)
	// Множественное присваивание
	a, b = reset() //сброс к a=8 b=4
	fmt.Printf("Второй вариант:\na=%v b=%v\n", a, b)
	a, b = b, a
	fmt.Printf("a=%v b=%v\n", a, b)
	// Сложение и вычитание
	a, b = reset()
	fmt.Printf("Третий вариант:\na=%v b=%v\n", a, b)
	a = a - b
	b = a + b
	a = b - a
	fmt.Printf("a=%v b=%v\n", a, b)
}
