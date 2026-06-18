package main

import (
	"fmt"
	"go-prac/shared"
	"math/rand"
)

func main() {
	fmt.Println("I'm on the go")
	var count = 0
	for count < 10 {
		var num = rand.Intn(10)
		fmt.Println(num)
		count++
	}
	fmt.Println("end of random generator")

	user1 := shared.Admin{
		Person: shared.Person{
			Name: "alice",
		},
		Age: 29,
	}

	user2 := shared.Admin{
		Person: shared.Person{
			Name: "bob",
		},
		Age: 10,
	}

	user1.Greet()
	user2.Greet()

	/*
		name1, age1 := user1.UpdateName(true)
		fmt.Println(name1, age1)
		name2, age2 := user2.UpdateName(false)
		fmt.Println(name2, age2)

		var admin1 shared.User
		admin1 = User{Name: "Alice"}

		// 4. Call the method directly on the variable

		admin2 := shared.Person{Name: "John"}
		admin2.greet()
	*/
}
