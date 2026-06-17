package main

import (
	"fmt"
	"math/rand"
)

func main() {
	simple()
	fmt.Println("I'm on go")
	var count = 0
	for count < 10 {
		var num = rand.Intn(10)
		fmt.Println(num)
		count++
	}
	fmt.Println("Type check")
	type_check()
	fmt.Println("example of method")

	user1 := Human{
		Name: "Alice",
		Age:  29,
	}

	user2 := Human{
		Name: "bob",
		Age:  10,
	}

	name1, age1 := user1.UpdateName(true)
	fmt.Println(name1, age1)
	name2, age2 := user2.UpdateName(false)
	fmt.Println(name2, age2)

	var admin1 User
	admin1 = User{Name: "Alice"}

	// 4. Call the method directly on the variable
	admin1.Greet()

	admin2 := Person{Name: "John"}
	admin2.greet()
}
