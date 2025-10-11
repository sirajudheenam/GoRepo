package main

import "fmt"

// Empty Interface (interface{})
// A universal type — can hold any value (like any in TypeScript or Object in Java).
func describe(i interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", i, i)
}

func main() {
	describe(42)
	describe("hello")
	describe(true)
}
