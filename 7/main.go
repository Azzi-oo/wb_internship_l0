package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	data := make(map[string]int)
	mutex := &sync.Mutex{}

	writeData := func(key string, value int) {
		mutex.Lock()
		data[key] = value
		mutex.Unlock()
	}

	go writeData("key1", 100)
	go writeData("key2", 200)
	go writeData("key3", 300)

	fmt.Println("Ждем завершения всех горутин...")
	time.Sleep(1 * time.Second)

	fmt.Println("Результат записи в map:")
	fmt.Println(data)
}
