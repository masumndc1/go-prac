package main

// example of variadic parameters
// ...string means variadic parameter
// the function accepts zero or more string arguments.
// Common use: fmt.Println(), fmt.Printf() use variadic parameters.

import "fmt"

/*
func What(args ...string) {
 	// args is a slice: []string
}
*/

func What(args ...string) {
	for i, arg := range args {
		fmt.Println(i, arg)
	}
}

// Multiple parameters (variadic must be last):
func Greet(name string, hobbies ...string) {
	fmt.Printf("%s enjoys: %v\n", name, hobbies)
}

func main() {

	What()                 // 0 args, no output
	What("hello")          // 1 arg
	What("hello", "world") // 2 args
	// Inside the function:
	// Output:
	// 0 hello
	// 1 world
	What("a", "b", "c", "d") // 4 args

	// Unpacking a slice with ...
	fruits := []string{"apple", "banana", "cherry"}
	// Unpacks the slice as separate arguments
	What(fruits...)

	Greet("Alice", "reading", "hiking", "coding")
	// Output: Alice enjoys: [reading hiking coding]

}
