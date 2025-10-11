package main

import "fmt"

// Define an interface
type Speaker interface {
	Speak() string
}

// Implement the interface
type Dog struct{}

// Implement the Speak method for Dog
func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

// Implement the Speak method for Cat
func (c Cat) Speak() string {
	return "Meow!"
}

// Function that accepts any type that implements Speaker
func makeItSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

// ✅ Both Dog and Cat automatically satisfy the Speaker interface since they both implement Speak() method.
// No explicit declaration of intent is needed.

// Implicit Implementation:
// No need to declare implements. Go automatically detects matching methods.

// Dynamic Typing
// An interface value can hold any type that satisfies it.

// Nil Interface Values
// Both the type and value can be nil.

// Empty Interface (interface{})
// A universal type — can hold any value (like any in TypeScript or Object in Java).

func main() {
	// Create instances of Dog and Cat
	d := Dog{}
	c := Cat{}

	// Pass them to the function
	makeItSpeak(d)
	makeItSpeak(c)
}
