package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan struct{})
	timer := time.After(3 * time.Second)
	ch := make(chan int)
	go func() {
		i := 0
		defer close(ch)
		for {
			select {
			case <-stop:
				fmt.Println("Пришел сигнал завершения")
				return
			default:
				ch <- i
				i++
				time.Sleep(1 * time.Second)
			}
		}
	}()
	for {
		select {
		case <-timer:
			fmt.Println("Отправляем сигнал")
			close(stop)
			timer = nil
		case v, ok := <-ch:
			if !ok {
				fmt.Println("Канал закрыт")
				return
			}
			fmt.Println("Получили ", v)
		}

	}
}
