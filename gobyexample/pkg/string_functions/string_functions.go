package stringfunctions

// The standard library’s strings package provides many useful string-related functions.
// Here are some examples to give you a sense of the package.

import (
	"fmt"
	s "strings"
)

// We alias fmt.Println to a shorter name as we’ll use it a lot below.

var p = fmt.Println

func Run() {

	p("\nString Functions: ")

	// Here’s a sample of the functions available in strings. Since these are functions from the package,
	// not methods on the string object itself, we need to pass the string in question as the first argument to the function.
	// You can find more functions in the strings package docs.

	p("[test] Contains [es]:  ", s.Contains("test", "es"))
	p("Count:    'test' has [", s.Count("test", "t"), "] 't' (s)")
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
}

// $ go run string-functions.go
// Contains:   true
// Count:      2
// HasPrefix:  true
// HasSuffix:  true
// Index:      1
// Join:       a-b
// Repeat:     aaaaa
// Replace:    f00
// Replace:    f0o
// Split:      [a b c d e]
// ToLower:    test
// ToUpper:    TEST
