package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func main() {
	counter := Counter{}

	var wg sync.WaitGroup

	numGO := 100

	wg.Add(numGO)

	for i := 0; i < numGO; i++ {
		go func() {
			counter.Increment()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Итоговое значение счетчика: ", counter.count)
}
