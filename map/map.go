package main

import "fmt"

func GetKeyIntValString() map[int]string {
	return map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
}
func GetKeyStringValStringMap() map[string]string {

	return map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
	}
}

func CreateMap() {
	fmt.Println("Inside CreateMap")
	m := make(map[string][]string)
	m["one"] = []string{"1", "2", "3"}
	m["two"] = []string{"4", "5", "6"}
	m["three"] = []string{"7", "8", "9"}
	fmt.Println("Map: ", m)

	// Accessing the map
	fmt.Println("Accessing the map")
	fmt.Println("m[\"one\"]: ", m["one"])
	fmt.Println("m[\"two\"]: ", m["two"])
	fmt.Println("m[\"three\"]: ", m["three"])
	
}

yamlContent := `apiVersion: v1
kind: Pod
metadata:
  name: mypod
  labels: [app: myapp]
spec:
  containers:
	- name: mycontainer
	  image: myimage
	  ports:
	  	- containerPort: 80
	- name: mycontainer2
	  image: myimage2
	  ports:
	  	- containerPort: 8080
`