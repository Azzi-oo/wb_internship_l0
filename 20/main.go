package main

import (
	"fmt"
	"strings"
)

func reverseWords(sentence string) string {
	words := strings.Fields(sentence)
	reversed := make([]string, len(words))

	for i, j := 0, len(words)-1; i < len(words); i, j = i+1, j-1 {
		reversed[i] = words[j]
	}

	return strings.Join(reversed, " ")
}

func main() {
	input := "snow dog sun"
	fmt.Println("Строка:", input)

	reversed := reverseWords(input)
	fmt.Println("Строка с перевернутыми символами:", reversed)
}
