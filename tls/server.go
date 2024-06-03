package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, TLS! Guten Tag !! \n")
}

func main() {
	// Define the TLS configuration
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// Create a new HTTP server
	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}

	http.HandleFunc("/hello", helloHandler)

	// Start the server with TLS
	log.Println("Starting server on https://localhost:8443")
	err := server.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
