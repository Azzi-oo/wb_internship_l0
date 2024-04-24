package main

import (
	"fmt"
)

type Human struct {
	Name    string
	Age     int
	Country string
}

func (h *Human) Introduce() {
	fmt.Printf("Hello, my name is %s. I am %d years old. I am from %s.\n", h.Name, h.Age, h.Country)
}

type Action struct {
	Human
}

func (a *Action) DoAction() {
	fmt.Println("I am doing some action!")
}

func main() {
	action := Action{
		Human: Human{
			Name:    "Azza",
			Age:     30,
			Country: "USA",
		},
	}

	action.Introduce()
	action.DoAction()
}
