package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, target int) bool {
	index := sort.SearchInts(arr, target)
	return index < len(arr) && arr[index] == target
}

func main() {
	arr := []int{1, 3, 2, 3, 5, 2, 6}

	target := 5

	if binarySearch(arr, target) {
		fmt.Printf("Элемент найден %d", target)
	} else {
		fmt.Printf("Элемент не найден %d", target)
	}

	target = 8
	if binarySearch(arr, target) {
		fmt.Printf("Элемент найден %d", target)
	} else {
		fmt.Printf("Элемент не найден %d", target)
	}

}
