package person
type Person struct {
	Name string
}

// Implement the Speak method for Person
func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}
