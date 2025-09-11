package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ch := make(chan int)
	go func() {
		i := 0
		defer close(ch)
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Таймаут!(producer)")
				return
			case <-ticker.C:
				ch <- i
				i++
			}
		}
	}()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Таймаут!(consumer)")
			return
		case v, ok := <-ch:
			if !ok {
				fmt.Println("Канал закрыт")
				return
			}
			fmt.Println("Получили ", v)
		}
	}
}
