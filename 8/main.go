package main

import (
	"fmt"
)

func setBit(num int64, i int, bit int) int64 {
	mask := int64(1) << i
	if bit == 1 {
		num |= mask
	} else if bit == 0 {
		num &^= mask
	}
	return num
}

func main() {
	var num int64 = 10
	i := 2
	bit := 1
	result := setBit(num, i, bit)
	fmt.Println("Результат:", result)
}
