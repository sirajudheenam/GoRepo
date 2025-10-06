package main

import (
	"fmt"
	"log"
	"os"

	myMap "github.com/sirajudheenam/GoRepo/mapPackage"
)

func main() {
	log.SetPrefix("[map] - ")
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Inside main")
	fmt.Println("CWD: ", pwd)
	fmt.Println("IntKeyMap: ", myMap.GetKeyIntValString())
	fmt.Println("StringKeyMap: ", myMap.GetKeyStringValStringMap())
	myMap.CreateMap()
	fmt.Println(myMap.PrintYaml())
	log.Println("Exiting main")
}
