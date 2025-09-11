package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan int)

	// продюсер
	go func() {
		defer close(ch)
		i := 0
		t := time.NewTicker(500 * time.Millisecond)
		defer t.Stop()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("продюсер: отмена контекста -> выхожу")
				return
			case <-t.C:
				ch <- i
				i++
			}
		}
	}()

	// консюмер (когда увидит 3 - отменит контекст)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("консюмер: канал закрыт -> выхожу")
				return
			}
			fmt.Println("консюмер: получили", v)
			if v == 3 {
				fmt.Println("консюмер: делаю отмену контекста")
				cancel()
			}
		case <-ctx.Done():
			for v := range ch {
				fmt.Println("консюмер: остатки - Получил", v)
			}
			fmt.Println("консюмер: контекст отменен -> выхожу")
			return
		}

	}
	fmt.Println("Готово")
}
