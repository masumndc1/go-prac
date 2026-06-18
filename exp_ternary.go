package main

import "fmt"

func main() {
	// var x interface{} = "hello"
	var x interface{} = 12
	if _, ok := x.(string); !ok {
		y := x
		fmt.Println(y)
	}

	// Unsafe way (panics if wrong type):
	// str := x.(string)

	// Safe way (checks first):
	// if _, ok := x.(string); !ok {
	if _, ok := x.(string); !ok {
		fmt.Println("x is not a string!")
		fmt.Println(ok)
	}

	// another way
	// str, ok := x.(string)
	// if !ok {
	// 	fmt.Println("x is not a string")
	// } else {
	// 	fmt.Printf("x is string and value: %s\n", str)
	// }
}
