package main

import (
	"fmt"
)

func reverseString(str string) string {
	runes := []rune(str)
	length := len(runes)

	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}
	return string(runes)

}

func main() {
	input := "Hello"
	fmt.Println("Исходная строка:", input)

	reversed := reverseString(input)
	fmt.Println("Перевернутая строка:", reversed)
}
