package main

import "fmt"

func removeIndex(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	indexToRemove := 2

	numbers = removeIndex(numbers, indexToRemove)

	fmt.Println(numbers)
}
