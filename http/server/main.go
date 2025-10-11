package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func requestToString(r *http.Request) string {
	var buf bytes.Buffer

	// Method and URL
	buf.WriteString(fmt.Sprintf("Method: %s\n", r.Method))
	buf.WriteString(fmt.Sprintf("URL: %s\n", r.URL.String()))
	buf.WriteString(fmt.Sprintf("Proto: %s\n", r.Proto))
	buf.WriteString(fmt.Sprintf("RemoteAddr: %s\n", r.RemoteAddr))

	// Query parameters
	if len(r.URL.Query()) > 0 {
		buf.WriteString("\nQuery Parameters:\n")
		for key, values := range r.URL.Query() {
			for _, v := range values {
				buf.WriteString(fmt.Sprintf("  %s: %s\n", key, v))
			}
		}
	}

	// Headers
	buf.WriteString("\nHeaders:\n")
	// Authorization Header can be decoded with: echo "encodedString" | base64 -D
	for key, values := range r.Header {
		for _, v := range values {
			buf.WriteString(fmt.Sprintf("  %s: %s\n", key, v))
		}
	}

	// Body (read and restore)
	if r.Body != nil {
		bodyBytes, _ := io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // restore body for later use

		if len(bodyBytes) > 0 {
			buf.WriteString("\nBody:\n")
			buf.WriteString(string(bodyBytes))
			buf.WriteString("\n")
		} else {
			buf.WriteString("\nBody: (empty)\n")
		}
	}

	return buf.String()
}

// Example HTTP handler
func handler(w http.ResponseWriter, r *http.Request) {
	reqStr := requestToString(r)
	fmt.Println(reqStr) // log to console
	w.Write([]byte("Request received!\n"))
}

// Example HTTP handler - API endpoint
func apiHandler(w http.ResponseWriter, r *http.Request) {
	reqStr := requestToString(r)
	fmt.Println(reqStr) // log to console
	w.Write([]byte("API Request received!\n"))
}

func main() {
	http.HandleFunc("/api/v1/users", apiHandler) // API endpoint
	http.HandleFunc("/", handler)                // General endpoint catch-all

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
