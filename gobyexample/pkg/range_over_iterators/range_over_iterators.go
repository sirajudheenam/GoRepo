package rangeoveriterators

// Starting with version 1.23, Go has added support for iterators, which lets us range over pretty much anything!

import (
	"fmt"
	"iter"
	"slices"
)

// Let’s look at the List type from the previous example again.
//  In that example we had an AllElements method that returned a slice of all elements in the list.
//  With Go iterators, we can do it better - as shown below.

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// All returns an iterator, which in Go is a function with a special signature.

// The iterator function takes another function as a parameter, called yield by convention
// (but the name can be arbitrary). It will call yield for every element we want to iterate over,
// and note yield’s return value for a potential early termination.

func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {

		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

// Iteration doesn’t require an underlying data structure, and doesn’t even have to be finite!
// Here’s a function returning an iterator over Fibonacci numbers: it keeps running as long as yield keeps returning true.

func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func Run() {
	fmt.Println("\nRange Over Iterators")
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	// Since List.All returns an iterator, we can use it in a regular range loop.

	for e := range lst.All() {
		fmt.Println(e)
	}

	// Packages like slices have a number of useful functions to work with iterators.
	//  For example, Collect takes any iterator and collects all its values into a slice.

	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for n := range genFib() {

		// Once the loop hits break or an early return, the yield function passed to the iterator will return false.

		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
