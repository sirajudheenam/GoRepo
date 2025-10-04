package sorting

import (
	"fmt"
	"slices"
) // Go’s slices package implements sorting for builtins and user-defined types. We’ll look at sorting for builtins first.

func Run() {
	fmt.Println("\nSorting: ")

	// Sorting functions are generic, and work for any ordered built-in type. For a list of ordered types, see cmp.Ordered.

	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	// An example of sorting ints.

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)

	// We can also use the slices package to check if a slice is already in sorted order.

	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}

// $ go run sorting.go
// Strings: [a b c]
// Ints:    [2 4 7]
// Sorted:  true
