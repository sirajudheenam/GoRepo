package main

import (
	"fmt"

	"github.com/sirajudheenam/GoRepo/interfaces/exampleSpeak/dog"
	"github.com/sirajudheenam/GoRepo/interfaces/exampleSpeak/person"
)

func main() {
	// Create a Person and a Dog
	p := person.Person{Name: "Sam"}
	d := dog.Dog{Breed: "Labrador"}

	// Call the Speak method on the Person and Dog
	fmt.Println(p.Speak())
	fmt.Println(d.Speak())
}
