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
		defer close(ch)
		ticker := time.NewTicker(490 * time.Millisecond)
		defer ticker.Stop()
		i := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Тайм аут (producer)")
				return
			case <-ticker.C:
				ch <- i
				i++
			}
		}
	}()
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("Канал закрыт")
				return
			}
			fmt.Printf("Получено значение %v\n", v)
		case <-ctx.Done():
			fmt.Println("Тайм аут (consumer)")
			return
		}
	}
}
