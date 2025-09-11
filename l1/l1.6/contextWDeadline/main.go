package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()
	ch := make(chan int)
	go func() {
		i := 0
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Дедлайн!(producer)")
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
		case <-ctx.Done():
			fmt.Println("Дедлайн!(consumer)")
			return
		case v, ok := <-ch:
			if !ok {
				fmt.Println("Канал закрыт")
			}
			fmt.Println("Получили ", v)
		}
	}
}
