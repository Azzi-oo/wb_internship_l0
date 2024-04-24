package main

import (
	"fmt"
)

func operations(a, b int) {
	if a > (1<<20) && b > (1<<20) {
		addition := a + b
		fmt.Println("Sum:", addition)

		substraction := a - b
		fmt.Println("Sub:", substraction)

		multiplication := a * b
		fmt.Println("Mul:", multiplication)

		if b != 0 {
			division := a / b
			fmt.Println("Div:", division)
		} else {
			fmt.Println("Div on zero!")
		}
	} else {
		fmt.Println("Меньше 2^20")
	}
}

func main() {
	a := 1048576
	b := 1048577
	operations(a, b)
}
