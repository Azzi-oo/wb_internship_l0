package main

import (
	"fmt"
	"time"
)

func sendData(ch chan<- int) {
	defer close(ch)
	for i := 1; ; i++ {
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	ch := make(chan int)

	go sendData(ch)

	for {
		select {
		case data, ok := <-ch:
			if !ok {
				fmt.Println("Канал закрыт")
				return
			}
			fmt.Printf("Принято значение из канала: %d\n", data)
		case <-time.After(5 * time.Second):
			fmt.Println("Время истекло. Программа завершает работу.")
			return
		}
	}
}
