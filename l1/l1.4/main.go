package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func startWorkers(ch chan int, n int, wg *sync.WaitGroup, ctx context.Context) { // n - количество воркеров
	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(ch, i, wg, ctx)
	}
}
func worker(ch chan int, id int, wg *sync.WaitGroup, ctx context.Context) {
	//сработает как раз перед завершением горутины
	defer wg.Done()
	for {
		select {
		// в случае отмены контекста сработает этот вариант и произойдет завершение горутины
		case <-ctx.Done():
			fmt.Printf("Воркер id-%v завершается...\n", id)
			return
		// основная логика работы горутины
		case v, ok := <-ch:
			if !ok {
				fmt.Printf("Воркер %v: канал закрыт\n", id)
				return
			}
			fmt.Printf("Воркер %v получил значение %v\n", id, v)
		}
	}
}

func main() {
	// контекст с отменой для завершения работы горутин
	ctx, cancel := context.WithCancel(context.Background())
	// канал для сигнала прерывания
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT) // SIGINT - interrupt(CTRL + C)

	var wg sync.WaitGroup
	ch := make(chan int)
	startWorkers(ch, 10, &wg, ctx)
	// постоянная запись в канал
	go func() {
		i := 0
		for {
			select {
			// завершение работы горутины в случае отмены контекста
			case <-ctx.Done():
				close(ch)
				return
			// основная логика работы горутины
			default:
				ch <- i
				i++
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
	fmt.Println("Нажмите CTRL+C для корректного завершения")
	sig := <-sigs
	fmt.Printf("Пришел сигнал - %v,завершаем...\n", sig)
	cancel()

	wg.Wait()
	fmt.Println("Все воркеры завершили работу")
}
