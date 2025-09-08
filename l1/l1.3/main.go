package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func startWorkers(ch chan int, n *int, wg *sync.WaitGroup) { // n - количество воркеров
	for i := 0; i < *n; i++ {
		wg.Add(1)
		go worker(ch, i, wg)
	}
}
func worker(ch chan int, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Printf("Воркер %v получил значение: %v\n", id, v)
	}
}

func main() {
	workers := flag.Int("workers", 4, "количество воркеров")
	flag.Parse()
	//n := 10 // ИЗМЕНЕНИЕ КОЛ-ВА ВОРКЕРОВ
	var wg sync.WaitGroup
	ch := make(chan int)
	startWorkers(ch, workers, &wg)
	// постоянная запись в канал
	go func() {
		i := 0
		for {
			ch <- i
			i++
			time.Sleep(100 * time.Millisecond)
		}
	}()
	time.Sleep(10 * time.Second)
	close(ch)
	wg.Wait()
}
