package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

// Command and CommandType remain the same
type CommandType int

const (
	GetCommand = iota
	SetCommand
	IncCommand
)

type Command struct {
	ty        CommandType
	name      string
	val       int
	replyChan chan int
	reqID     string // request ID for logging/tracing
}

// Helper function to generate random 16-byte hex request IDs
func newRequestID() string {
	b := make([]byte, 8)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

// startCounterManager now logs request IDs for visibility
func startCounterManager(ctx context.Context, initvals map[string]int) chan<- Command {
	counters := make(map[string]int)
	for k, v := range initvals {
		counters[k] = v
	}

	cmds := make(chan Command)

	go func() {
		defer log.Println("Counter manager stopped")

		for {
			select {
			case <-ctx.Done():
				log.Println("Manager received cancel signal:", ctx.Err())
				return
			case cmd, ok := <-cmds:
				if !ok {
					return
				}

				// Log each operation with request ID
				log.Printf("[req:%s] handling %v for '%s'\n", cmd.reqID, cmd.ty, cmd.name)

				switch cmd.ty {
				case GetCommand:
					val, ok := counters[cmd.name]
					if ok {
						cmd.replyChan <- val
					} else {
						cmd.replyChan <- -1
					}
				case SetCommand:
					counters[cmd.name] = cmd.val
					cmd.replyChan <- cmd.val
				case IncCommand:
					if _, ok := counters[cmd.name]; ok {
						counters[cmd.name]++
						cmd.replyChan <- counters[cmd.name]
					} else {
						cmd.replyChan <- -1
					}
				}
			}
		}
	}()

	return cmds
}

// Server type unchanged
type Server struct {
	cmds chan<- Command
}

// Middleware-like wrapper to assign Request ID and inject into context
func withRequestID(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		reqID := newRequestID()
		ctx := context.WithValue(req.Context(), "reqID:", reqID)
		log.Printf("[req:%s] %s %s", reqID, req.Method, req.URL.Path)
		next(w, req.WithContext(ctx))
	}
}

// Helper to retrieve the request ID from context
func getRequestID(ctx context.Context) string {
	if v := ctx.Value("reqID:"); v != nil {
		return v.(string)
	}
	return "unknown"
}

func (s *Server) get(w http.ResponseWriter, req *http.Request) {
	reqID := getRequestID(req.Context())
	name := req.URL.Query().Get("name")
	replyChan := make(chan int)

	s.cmds <- Command{ty: GetCommand, name: name, replyChan: replyChan, reqID: reqID}
	reply := <-replyChan

	if reply >= 0 {
		fmt.Fprintf(w, "[req:%s] %s: %d\n", reqID, name, reply)
	} else {
		fmt.Fprintf(w, "[req:%s] %s not found\n", reqID, name)
	}
}

func (s *Server) set(w http.ResponseWriter, req *http.Request) {
	reqID := getRequestID(req.Context())
	name := req.URL.Query().Get("name")
	val := req.URL.Query().Get("val")
	intval, err := strconv.Atoi(val)
	if err != nil {
		fmt.Fprintf(w, "[req:%s] %s\n", reqID, err)
		return
	}

	replyChan := make(chan int)
	s.cmds <- Command{ty: SetCommand, name: name, val: intval, replyChan: replyChan, reqID: reqID}
	<-replyChan
	fmt.Fprintf(w, "[req:%s] ok\n", reqID)
}

func (s *Server) inc(w http.ResponseWriter, req *http.Request) {
	reqID := getRequestID(req.Context())
	name := req.URL.Query().Get("name")
	replyChan := make(chan int)

	s.cmds <- Command{ty: IncCommand, name: name, replyChan: replyChan, reqID: reqID}
	reply := <-replyChan

	if reply >= 0 {
		fmt.Fprintf(w, "[req:%s] ok\n", reqID)
	} else {
		fmt.Fprintf(w, "[req:%s] %s not found\n", reqID, name)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server := Server{startCounterManager(ctx, map[string]int{"i": 0, "j": 0})}

	// Wrap each handler with the Request ID generator
	http.HandleFunc("/get", withRequestID(server.get))
	http.HandleFunc("/set", withRequestID(server.set))
	http.HandleFunc("/inc", withRequestID(server.inc))

	srv := &http.Server{Addr: ":8000"}

	go func() {
		log.Println("Server started on :8000")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Graceful shutdown logic
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println("Shutting down...")

	cancel()
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited cleanly.")
}
