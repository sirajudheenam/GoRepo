package main

import (
	"fmt"
	"net/http"
	"time"
)

func LimitConcurrentRequestsMiddleware(maxRequests int) func(http.Handler) http.Handler {
	return func(inner http.Handler) http.Handler {
		if maxRequests == 0 {
			return inner
		}

		semaphore := make(chan struct{}, maxRequests)
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			semaphore <- struct{}{}        // acquire
			defer func() { <-semaphore }() // release
			inner.ServeHTTP(w, r)
		})
	}
}

func slowHandler(w http.ResponseWriter, r *http.Request) {

	time.Sleep(3 * time.Second)
	fmt.Fprintf(w, "Processed: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Request from: %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "%+v", r)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", slowHandler)

	limited := LimitConcurrentRequestsMiddleware(3)(mux)
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", limited)
}
