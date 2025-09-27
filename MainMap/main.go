package main

import (
	"fmt"
	"log"
	"os"
	map "github.com/sirajudheenam/GoRepo/mapPackage"
)

func main() {
	log.SetPrefix("[map] - ")
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Inside main")
	fmt.Println("CWD: ", pwd)
	fmt.Println("IntKeyMap: ", map.GetKeyIntValString())
	fmt.Println("StringKeyMap: ", map.GetKeyStringValStringMap())
	map.CreateMap()
	log.Println("Exiting main")
}
