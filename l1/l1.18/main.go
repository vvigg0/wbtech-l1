package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type counter struct {
	val int
	m   sync.Mutex
}

func worker(i int, c *counter, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %v начал работу\n", i)
	// какая-то работа
	fmt.Printf("worker %v Выполняет работу\n", i)
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	c.m.Lock()
	c.val++
	c.m.Unlock()
	fmt.Printf("worker %v заканчивает работу\n", i)
}
func main() {
	var wg sync.WaitGroup
	var c counter
	n := 22 // кол-во горутин "выполняющих работу"
	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(i, &c, &wg)
	}
	wg.Wait()
	fmt.Println(c.val)
}
