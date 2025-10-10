// Source: https://github.com/eliben/code-for-blog/blob/main/2019/gohttpconcurrency/channel-manager-server.go
// This version of the server protects all shared data within a manager
// goroutine that accepts commands using a channel.
//
// Eli Bendersky [http://eli.thegreenplace.net]
// This code is in the public domain.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type CommandType int // CommandType is an alias for int, representing types of operations

const (
	GetCommand = iota // GetCommand == 0; used to fetch a counter value
	SetCommand        // SetCommand == 1; used to set a counter to a value
	IncCommand        // IncCommand == 2; used to increment a counter
)

// Command represents a message sent to the counter manager goroutine.
// Each command includes the operation type, the counter name, and optionally a value.
// replyChan is used to send the result of the operation back to the sender.
type Command struct {
	ty        CommandType // type of command: Get, Set, or Inc
	name      string      // name of the counter to operate on
	val       int         // value for SetCommand
	replyChan chan int    // a channel for sending back results
}

// startCounterManager launches a dedicated goroutine that manages all counter state.
// The goroutine listens for incoming Command objects sent on the returned channel.
// This ensures only one goroutine ever touches the shared counter map.
func startCounterManager(initvals map[string]int) chan<- Command {
	counters := make(map[string]int) // local map holding counter name → value
	for k, v := range initvals {
		// Initialize each counter with the given starting value
		log.Printf("Assigning counters[%v] to a value of %v\n", k, v)
		counters[k] = v
	}

	cmds := make(chan Command) // this channel will be used to send commands to the manager

	// The goroutine that handles all operations on the counters
	go func() {
		for cmd := range cmds { // continuously read commands from the channel
			switch cmd.ty {
			case GetCommand:
				// Lookup the counter value and send it back via replyChan
				if val, ok := counters[cmd.name]; ok {
					fmt.Printf("send the counter value back: [GET] [cmd.replyChan]: %v [val]: %v\n", cmd.replyChan, val)
					cmd.replyChan <- val // send the counter value back
				} else {
					fmt.Printf("counter not found, send -1 back: [GET] [cmd.replyChan]: %v\n", cmd.replyChan)
					cmd.replyChan <- -1 // send -1 if the counter doesn’t exist
				}
			case SetCommand:
				// Set the counter to a specific value
				counters[cmd.name] = cmd.val
				cmd.replyChan <- cmd.val // acknowledge with the value set
				fmt.Printf("send the counter value back [SET]: [cmd.replyChan]: %v [val]: %v\n", cmd.replyChan, cmd.val)
			case IncCommand:
				// Increment the counter if it exists
				if _, ok := counters[cmd.name]; ok {
					counters[cmd.name]++
					cmd.replyChan <- counters[cmd.name] // send back the incremented value
					fmt.Printf("send the counter value back [INC]: [cmd.replyChan]: %v [counters[cmd.name]]: %v [cmd.name]: %v\n", cmd.replyChan, counters[cmd.name], cmd.name)
				} else {
					cmd.replyChan <- -1 // send -1 if not found
				}
			default:
				// Defensive programming: crash if an unknown command type appears
				log.Fatal("unknown command type", cmd.ty)
			}
		}
	}()
	return cmds // return the channel to the caller so it can send commands
}

// Server is a struct holding a command channel used by HTTP handlers.
// Each HTTP request sends a command through this channel to interact
// with the manager goroutine.
type Server struct {
	cmds chan<- Command // write-only channel for sending commands
}

// HTTP handler for the /get endpoint.
// It retrieves a counter value by name and writes the result to the HTTP response.
func (s *Server) get(w http.ResponseWriter, req *http.Request) {
	log.Printf("get %v", req)
	name := req.URL.Query().Get("name") // read ?name=<counter> from the URL
	replyChan := make(chan int)         // channel to receive the result

	// Send a GetCommand to the manager
	s.cmds <- Command{ty: GetCommand, name: name, replyChan: replyChan}

	reply := <-replyChan // wait for the manager’s reply

	if reply >= 0 {
		// If the counter exists, print its name and value
		fmt.Fprintf(w, "%s: %d\n", name, reply)
	} else {
		// Otherwise indicate it was not found
		fmt.Fprintf(w, "%s not found\n", name)
	}
}

// HTTP handler for the /set endpoint.
// It sets a counter to a specified value (?name=<counter>&val=<number>).
func (s *Server) set(w http.ResponseWriter, req *http.Request) {
	log.Printf("set %v", req)
	name := req.URL.Query().Get("name") // get counter name
	val := req.URL.Query().Get("val")   // get value to set (as string)
	intval, err := strconv.Atoi(val)    // convert string to integer
	if err != nil {
		// If the conversion fails, print the error
		fmt.Fprintf(w, "%s\n", err)
	} else {
		replyChan := make(chan int)
		// Send SetCommand to the manager goroutine
		s.cmds <- Command{ty: SetCommand, name: name, val: intval, replyChan: replyChan}
		_ = <-replyChan // wait for acknowledgment (result not used)
		fmt.Fprintf(w, "ok\n")
	}
}

// HTTP handler for the /inc endpoint.
// It increments a counter by 1 (?name=<counter>).
func (s *Server) inc(w http.ResponseWriter, req *http.Request) {
	log.Printf("inc %v", req)
	name := req.URL.Query().Get("name")
	replyChan := make(chan int)
	// Send IncCommand to manager
	s.cmds <- Command{ty: IncCommand, name: name, replyChan: replyChan}

	reply := <-replyChan // wait for reply
	if reply >= 0 {
		fmt.Fprintf(w, "ok\n") // success
	} else {
		fmt.Fprintf(w, "%s not found\n", name) // counter not found
	}
}

// main initializes the server and starts listening for HTTP requests.
func main() {
	// Create a Server that talks to a counter manager initialized with counters i and j
	server := Server{startCounterManager(map[string]int{"i": 0, "j": 0})}

	// Register HTTP handlers for each endpoint
	http.HandleFunc("/get", server.get)
	http.HandleFunc("/set", server.set)
	http.HandleFunc("/inc", server.inc)

	// Default port is 8000; can be overridden by command-line argument
	portnum := 8000
	if len(os.Args) > 1 {
		portnum, _ = strconv.Atoi(os.Args[1])
	}
	log.Printf("Going to listen on port %d\n", portnum)

	// Start the HTTP server and log fatal errors (blocks forever)
	log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(portnum), nil))
}
