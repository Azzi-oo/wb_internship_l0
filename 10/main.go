package main

import (
	"fmt"
	"math"
	"sort"
)

func groupTemp(temp []float64, step float64) map[int][]float64 {
	temperaturesGroup := make(map[int][]float64)

	for _, temp := range temp {
		groupKey := int(math.Floor(temp/step)) * int(step)
		temperaturesGroup[groupKey] = append(temperaturesGroup[groupKey], temp)
	}

	return temperaturesGroup
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	step := 10.0

	groups := groupTemp(temperatures, step)

	var keys []int
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("%d: %v\n", k, groups[k])
	}
}
