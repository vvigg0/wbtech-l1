package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex //не RW потому что рассматриваем только запись

func worker(m map[int]int, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		mutex.Lock()
		m[id] = i
		mutex.Unlock()
		fmt.Printf("Записали [%v:%v]\n", id, i)
	}
}
func main() {
	var wg sync.WaitGroup
	m := make(map[int]int)
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go worker(m, i, &wg)
	}
	wg.Wait()
}
