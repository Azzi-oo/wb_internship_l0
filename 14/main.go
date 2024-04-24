package main

import (
	"fmt"
	"reflect"
)

func getType(variable interface{}) string {
	switch reflect.TypeOf(variable).Kind() {
	case reflect.Int:
		return "int"
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "bool"
	case reflect.Chan:
		return "channel"
	default:
		return "unknown"
	}
}

func main() {
	var integer int = 10
	var str string = "Hello"
	var boolean bool = true
	var ch chan int

	fmt.Println("Тип переменной int: ", getType(integer))
	fmt.Println("Тип перемнной str: ", getType(str))
	fmt.Println("Тип переменной boolean: ", getType(boolean))
	fmt.Println("Тип переменной ch: ", getType(ch))
}
