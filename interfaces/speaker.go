package main

import (
	"fmt"

	"github.com/sirajudheenam/GoRepo/interfaces/dog"
	"github.com/sirajudheenam/GoRepo/interfaces/person"
)

func main() {
	// Create a Person and a Dog
	p := person.Person{Name: "Sirajudheen"}
	d := dog.Dog{Breed: "Labrador"}

	// Call the Speak method on the Person and Dog
	fmt.Println(p.Speak())
	fmt.Println(d.Speak())
}