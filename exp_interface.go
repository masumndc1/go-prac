package main

import "fmt"

// Define an interface
type Animal interface {
	Speak() string
}

// Dog implements Animal
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + " says Woof!\n"
}

// Cat implements Animal
type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return c.Name + " tells Meow!\n"
}

// Function that accepts ANY type implementing Animal interface
func animalSound(a Animal) {
	fmt.Printf(a.Speak())
}

func main() {
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}

	animalSound(dog)
	animalSound(cat)

}

/*
The key points:

1. Interface definition — Animal requires a Speak() method
2. Implicit implementation — Dog and Cat both implement Animal automatically (no implements keyword)
3. Polymorphism — AnimalSound() accepts any type that implements Animal

Why this matters: You can add new animals (Bird, Cow, etc.) without changing AnimalSound().
It just works if they implement Speak().

output:
Buddy says Woof!
Whiskers says Meow!

*/
