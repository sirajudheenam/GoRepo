package main

import "fmt"

// This example demonstrates an interface with multiple methods in Go.
// Any type that implements all the methods of the interface satisfies the interface.
// Here we define a Shape interface with two methods: Area and Perimeter.
// We then create a Rectangle type that implements both methods.

// Define an interface with multiple methods

type Shape interface {
	Area() float64
	Perimeter() float64
}

// Implement the interface for a Rectangle type
type Rectangle struct {
	Width, Height float64
}

// Implementing the Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Implementing the Perimeter method for Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Function that accepts any type that implements Shape
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	rect := Rectangle{Width: 5, Height: 3}
	printShapeInfo(rect)
}
