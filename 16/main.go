package main

import "fmt"

func quicksort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	var less, greater, equal []int

	for _, num := range arr {
		switch {
		case num < pivot:
			less = append(less, num)
		case num == pivot:
			equal = append(equal, num)
		case num > pivot:
			greater = append(greater, num)
		}
	}

	return append(append(quicksort(less), equal...), quicksort(greater)...)
}

func main() {
	arr := []int{3, 4, 5, 6, 7, 8, 7}
	fmt.Println(quicksort(arr))
}
