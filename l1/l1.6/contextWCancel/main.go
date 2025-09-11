package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan int)

	// продюсер
	go func() {
		defer close(ch)
		for i := 0; i < 6; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
			}
		}
	}()

	// консюмер
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Отмена контекста, завершаем...")
			return
		case v, ok := <-ch:
			if !ok {
				fmt.Println("Канал закрыт")
				return
			}
			fmt.Println("Получили:", v)
		}
	}
}
