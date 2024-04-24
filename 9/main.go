package main

import "fmt"

func main() {
	nums := make(chan int)
	double_nums := make(chan int)

	go func() {
		arr := []int{1, 2, 3, 4, 5}
		for _, num := range arr {
			nums <- num
		}
		close(nums)
	}()

	go func() {
		for num := range nums {
			double_nums <- num * 2
		}
		close(double_nums)
	}()

	for doubledNum := range double_nums {
		fmt.Println(doubledNum)
	}
}
