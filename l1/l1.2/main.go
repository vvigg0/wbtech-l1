package main

import (
	"fmt"
	"sync"
)

var arr = []int{2, 4, 6, 8, 10}

type Result struct {
	index int
	value int
}

func getSquaresBase(arr []int) {
	ch := make(chan int)
	for _, val := range arr {
		go calcSquaresBase(ch, val)
		fmt.Println(<-ch)
	}
}
func calcSquaresBase(ch chan int, val int) {
	ch <- val * val
}

//Усложненные варианты
/*
Вариант с указателем
*/
func newArrPointer(arr []int) []int {
	var wg sync.WaitGroup
	newArr := make([]int, len(arr))
	for i, v := range arr {
		wg.Add(1)
		pointer := &newArr[i]
		go squaresArrPointer(pointer, v, &wg)
	}
	wg.Wait()
	return newArr
}
func squaresArrPointer(pointer *int, v int, wg *sync.WaitGroup) {
	defer wg.Done()
	*pointer = v * v
}

/*
Вариант с каналом
*/
func newArrChannel(arr []int) []int {
	var wg sync.WaitGroup
	newArr := make([]int, len(arr))
	ch := make(chan Result, len(arr))
	for i, v := range arr {
		wg.Add(1)
		go squaresArrChannel(ch, i, v, &wg)
	}
	wg.Wait()
	close(ch)
	for result := range ch {
		newArr[result.index] = result.value
	}
	return newArr
}
func squaresArrChannel(ch chan Result, i, v int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- Result{index: i, value: v * v}
}
func main() {
	fmt.Println("Исходный массив")
	fmt.Println(arr)
	fmt.Println("Квадраты чисел")
	getSquaresBase(arr)
	fmt.Println("Новый массив(через указатель)")
	fmt.Println(newArrPointer(arr))
	fmt.Println("Новый массив(через канал)")
	fmt.Println(newArrChannel(arr))
}
