package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const numWorkers = 3

func worker(id int, wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()
	for {
		data, ok := <-ch
		if !ok {
			fmt.Printf("Worker %d завершает работу\n", id)
			return
		}
		fmt.Printf("Worker %d принял данные: %d\n", id, data)
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, &wg, ch)
	}

	go func() {
		defer close(ch)
		for i := 1; ; i++ {
			ch <- i
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	fmt.Println("Ожидание завершения всех воркеров...")
	wg.Wait()
	fmt.Println("Все воркеры завершили работу")
}
