/* above is exact equivalent to c
void UpdateName (User *r, char *newName) {
	r->Name = newName;
}*/

/* anomaly of a function in go
  1      2      3          4               5
func (r *User) MethodName(param string) return_type { ... }
func (r *User) MethodName(param1 string, param2 int) (return_type1, return_type2) { ... }

func: The keyword signaling a function definition is starting.
(r *User) (The Pointer Receiver): This tells Go: "This isn't a standalone
  function. This is a method that belongs to the User struct.
  The r acts exactly like self in Python or this in C++.
  The * means we are passing a C-style memory pointer so we can modify the
  original struct's data.
MethodName: The name of the function (capitalized here so it is public!).
  (param string): The normal incoming input parameters.
return_type: The output data type (left blank if it returns nothing,
  like void in C).
*/

/*
	pointer way

// 1. Pointer Receiver: Can modify the outside world
// this means nothing but attaching a method to a struct
// GO SYNTAX

	func (u *User) UpdateAge(newAge int) {
	    u.Age = newAge // Modifies the original object directly!
	} vs

// passing the value way
// This change WILL BE LOST when the function ends
// GO SYNTAX

	func (u User) DisplayAge() {
	    fmt.Println(u.Age) // Reads data safely.
	    // If you tried to change u.Age here, it would only change inside this copy!
	}

Use the Pointer way (*User) by default. In real-world Go development, about 80%
to 90% of methods use pointer receivers because they allow you to modify data
and avoid the performance cost of copying memory.
*/

package main

import (
	"fmt"
)

type Human struct {
	Name string
	Age  int
}

func (r *Human) UpdateName(Active bool) (string, int) {
	if Active {
		r.Age++
		fmt.Println("active user")
	}
	return r.Name, r.Age
}
