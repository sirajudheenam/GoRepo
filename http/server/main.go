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
	w.Write([]byte("Request received!\n"))
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
