package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		i := 0
		defer close(ch)
		for {
			if i == 5 {
				fmt.Println("Число равно 5,остановка горутины...")
				runtime.Goexit()
				//остановка
			}
			ch <- i
			i++
			time.Sleep(500 * time.Millisecond)
		}
	}()
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("Канал закрыт\nПрограмма завершена")
			return
		}
		fmt.Println("Получили ", v)
	}
}
