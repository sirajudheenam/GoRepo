package main

import (
	"fmt"
	"net/http"

	"github.com/sirajudheenam/GoRepo/http/server/pkg/utils"
)

// Example HTTP handler
func handler(w http.ResponseWriter, r *http.Request) {
	reqStr := utils.RequestToString(r)
	fmt.Println(reqStr) // log to console

	switch r.Method {
	case http.MethodGet:
		fmt.Println("GET request received") // log to console
		w.Write([]byte("Hello! This is a GET request.\n"))
	case http.MethodPost:
		fmt.Println("POST request received") // log to console
		w.Write([]byte("Hello! This is a POST request.\n"))
	case http.MethodPut:
		fmt.Println("PUT request received") // log to console
		w.Write([]byte("Hello! This is a PUT request.\n"))
	case http.MethodDelete:
		fmt.Println("DELETE request received") // log to console
		w.Write([]byte("Hello! This is a DELETE request.\n"))
	case http.MethodPatch:
		fmt.Println("PATCH request received") // log to console
		w.Write([]byte("Hello! This is a PATCH request.\n"))
	case http.MethodHead:
		fmt.Println("HEAD request received") // log to console
		w.Write([]byte("Hello! This is a HEAD request.\n"))
	default:
		w.Write([]byte("Hello! This is some other type of request.\n"))
	}
}

// Example HTTP handler - API endpoint
func apiHandler(w http.ResponseWriter, r *http.Request) {
	reqStr := utils.RequestToString(r)
	fmt.Println(reqStr) // log to console
	w.Write([]byte("API Request received!\n"))
}

func main() {
	http.HandleFunc("/api/v1/users", apiHandler)  // API endpoint
	http.HandleFunc("/download", downloadHandler) // Download endpoint
	http.HandleFunc("/", handler)                 // General endpoint catch-all

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
