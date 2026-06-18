package main

// import (
// 	"fmt"
// 	"go-prac/shared"
// )

// feeling of python and
// include of c.
import "fmt"
import "go-prac/shared"

// A standard global function (C-style block) written in main.go
func greetPerson(p *shared.Admin) {
	fmt.Printf(
		`Hello, my name is %s
		And my age is %d`, p.Name, p.Age)
}

func main() {
	// Create an instance using the shared package struct
	user := shared.Admin{
		Age: 40,
		Person: shared.Person{
			Name: "bob",
		},
	}

	// Call it like a regular function, passing the address pointer explicitly
	greetPerson(&user)
}
