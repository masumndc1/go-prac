package shared

import (
	"fmt"
)

// The variables must start with a capital letter
var Few_words = [...]string{
	"bc",
	"cd",
	"df",
}
var Vowels = [5]string{"a", "e", "i", "o", "u"}

// 1. Using a map literal (with initial data)
var UserAge = map[string]int{
	"Alice": 25,
	"Bob":   30,
}

type Person struct {
	Name string
}

func (p *Person) Greet() {
	fmt.Printf("hello my name is: %s\n", p.Name)
}

type Admin struct {
	Person
	Age int
}
