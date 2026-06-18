package main

import "fmt"

func main() {
	score := 85
	age := 22
	hasID := true
	isVIP := false

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B") // Output: Grade: B
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// Example combining AND (&&) and OR (||)
	if (age >= 18 && hasID) || isVIP {
		fmt.Println("Access granted.") // Output: Access granted.
	} else {
		fmt.Println("Access denied.")
	}
}

/*
// Syntaxes look like: value, ok := interfaceVariable.(TargetType)
var i any = "Hello"

s, ok := i.(string) // Checks if 'i' holds a string
if ok {
    fmt.Println(s) // Safely use 's' as a normal string
}
*/
