package main

import (
	"fmt"
	"log"

	"github.com/sirajudheenam/goRepo/examples/greetings"
	"github.com/sapcc/go-bits/logg"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)

	// for i:=0; i<4; i++ {
	//     fmt.Println("Loop over")
	// }

	message, err := greetings.Hello("Sirajudheen")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
	
	logg.Info("This is information with Arguments %v : %v", 1, 2)

	logg.ShowDebug = true
	logg.Debug("This is a DEBUG msg based on ShowDebug FLAG (current Value is: %v)", logg.ShowDebug)
	logg.Error("This is an Error generated now : %v", "an error occured")
	const msgType string = "MORE INFO"
	logg.Other(msgType, "Additional Information given %s", "The sky is high")
}
