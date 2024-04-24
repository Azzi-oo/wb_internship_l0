package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isUnique("abcd"))
	fmt.Println(isUnique("abCdefAaf"))
	fmt.Println(isUnique("aabcd"))

}

func isUnique(str string) bool {
	str = strings.ToLower(str)
	uniqueMap := map[rune]struct{}{}
	for _, v := range []rune(str) {
		uniqueMap[v] = struct{}{}
	}
	if len([]rune(str)) != len(uniqueMap) {
		return false
	}
	return true
}
