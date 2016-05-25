package main

import (
	"./greeting"
)

func main() {

	salutations := greeting.Salutations{
		{"Bob", "Hello"},
		{"Joe", "Hi"},
		{"Mary", "What is up?"},
	}

	salutations[0].Rename("Alex")

	salutations.Greet(greeting.CreatePrintFunction("!!!"), true, 6)

	//salutations.Greet(greeting.CreatePrintFunction("!!!"), true, 6)
}
