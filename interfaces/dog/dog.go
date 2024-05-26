package dog

type Dog struct {
	Breed string
}

// Implement the Speak method for Dog
func (d Dog) Speak() string {
	return "Woof! I am a " + d.Breed
}
