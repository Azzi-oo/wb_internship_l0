package main

import "fmt"

func intersection(set1, set2 []int) []int {
	set1Map := make(map[int]bool)

	for _, num := range set1 {
		set1Map[num] = true
	}
	result := make([]int, 0)

	for _, num := range set2 {
		if set1Map[num] {
			result = append(result, num)
		}
	}
	return result
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}

	intersect := intersection(set1, set2)
	fmt.Println("Intersection:", intersect)
}
