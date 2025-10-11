package main

import "fmt"

func main() {
	var i interface{} = "hello"
	var j interface{} = map[string]int{"one": 1, "two": 2}
	var k interface{} = 42

	// Type switch
	switch v := i.(type) {
	case string:
		fmt.Println("string:", v)
	case int:
		fmt.Println("int:", v)
	case map[string]int:
		fmt.Println("map:", v)
	default:
		fmt.Println("unknown type")
	}

	s, ok := i.(string)
	if ok {
		fmt.Println(i, "It's a string:", s)
	}

	mapType, ok := j.(map[string]int)
	if ok {
		fmt.Println(j, ": It's a map:", mapType)
	}

	intType, ok := k.(int)
	if ok {
		fmt.Println(k, "It's an int:", intType)
	}
}
