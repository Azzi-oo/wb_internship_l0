package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	results := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(numbers))

	for _, num := range numbers {
		go func(n int) {
			defer wg.Done()
			square := n * n
			results <- square
		}(num)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	// Суммируем квадраты чисел из канала
	sum := 0
	for square := range results {
		sum += square
	}

	fmt.Println("Сумма квадратов чисел:", sum)
}
