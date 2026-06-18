package main

import (
	"fmt"
	"go-prac/shared"
)

func main() {

	// function name must be in capital in shared/vars.go.
	// otherwise it will give following error
	// admin.greet undefined (type shared.Person has no field
	// or method greet, but does have method Greet)
	// admin.greet()
	var admin = shared.Admin{
		Age: 32,
		Person: shared.Person{
			Name: "bob",
		},
	}

	admin.Greet()
	fmt.Printf("Admins name %s\n", admin.Person.Name)
	fmt.Printf("Admins age %d\n", admin.Age)
}
