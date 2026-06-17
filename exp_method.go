package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p Person) greet() {
	fmt.Printf("hello my name is %s\n", p.Name)
}
