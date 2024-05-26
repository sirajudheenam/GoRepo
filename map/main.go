package main

import (
	"fmt"
	"log"
	"os"
	// "github.com/sirajudheenam/GoRepo/map"
)

func main() {
	log.SetPrefix("[map] - ")
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Inside main")
	fmt.Println("CWD: ", pwd)
	fmt.Println("IntKeyMap: ", GetKeyIntValString())
	fmt.Println("StringKeyMap: ", GetKeyStringValStringMap())
	CreateMap()
	log.Println("Exiting main")
}