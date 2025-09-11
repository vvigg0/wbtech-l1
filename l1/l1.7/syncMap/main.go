package main

import (
	"fmt"
	"sync"
)

func worker(m *sync.Map, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		m.Store(id, i)
		fmt.Printf("Записали [%v:%v]\n", id, i)
	}
}
func main() {
	var wg sync.WaitGroup
	var m sync.Map
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go worker(&m, i, &wg)
	}
	wg.Wait()
}
