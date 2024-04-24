package main

import (
	"context"
	"fmt"
	"time"
)

const N = 5

func fibo(ctx context.Context) {
	x, y := 0, 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println(x)
			return
		default:
			x, y = y, x+y
			time.Sleep(100 * time.Millisecond)
		}

	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go fibo(ctx)

	time.Sleep(time.Millisecond * N * 100)

	cancel()

	time.Sleep(time.Millisecond * 100)
}
