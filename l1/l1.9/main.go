package main

import (
	"context"
	"fmt"
	"time"
)

func genNums(ctx context.Context, genCh chan int, nums []int) {
	defer close(genCh)
	for _, x := range nums {
		select {
		case <-ctx.Done():
			fmt.Println("генератор: контекст отменен -> выхожу")
			return
		case genCh <- x:
			fmt.Println("генератор: Сгенерировал число", x)
		}
	}
}
func procsNums(ctx context.Context, genCh, procsCh chan int) {
	defer close(procsCh)
	for {
		select {
		case v, ok := <-genCh:
			if !ok {
				fmt.Println("обработчик: канал закрыт -> выхожу")
				return
			}
			procsCh <- v * 2
		case <-ctx.Done():
			fmt.Println("обработчик: контекст отменен -> выхожу")
			return
		}

	}
}
func main() {
	nums := []int{1, 5, 9, 10, 15}
	ctx, cancel := context.WithCancel(context.Background())
	genCh := make(chan int)
	procsCh := make(chan int)
	go genNums(ctx, genCh, nums)
	go procsNums(ctx, genCh, procsCh)
	go func() {
		time.Sleep(1 * time.Microsecond)
		fmt.Println("main: отменяю контекст")
		cancel()
	}()
	for v := range procsCh {
		fmt.Println("Получили обработанное число:", v)
	}
	fmt.Println("Программа завершилась")
}
