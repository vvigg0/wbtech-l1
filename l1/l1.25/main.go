package main

import (
	"fmt"
	"time"
)

func sleep(tm time.Duration) {
	timer := time.NewTimer(tm)
	<-timer.C
}
func someFunc(ch chan int) {
	fmt.Println("someFunc | начинаю работу")
	a := 2
	b := 2

	fmt.Println("someFunc | засыпаю")
	sleep(5 * time.Second)

	fmt.Println("someFunc | просыпаюсь")
	sum := a + b
	ch <- sum
}
func main() {
	fmt.Println("main | начинаю работу")
	ch := make(chan int)

	fmt.Println("main | запускаю горутину")
	go someFunc(ch)

	fmt.Println("main | жду результата работы someFunc")
	fmt.Printf("Результат: %v", <-ch)

}
