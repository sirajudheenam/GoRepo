package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const serverPort = 3333

func main() {
	// make http server
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("server: %s /\n", r.Method)
			if r.Method == http.MethodGet {
				fmt.Printf("server: got GET request\n")
				w.Header().Add("Custom-Header", "custom-OF-value")

				id := r.URL.Query().Get("id")
				if id != "" {
				fmt.Printf("server: query id: %s\n", id )
				} else {
					fmt.Printf("server: query id: %s\n", "no id found" )
				}
				fmt.Printf("server: content-type: %s\n", r.Header.Get("content-type"))
				fmt.Printf("server: headers:\n")
				for headerName, headerValue := range r.Header {
					fmt.Printf("\t%s = %s\n", headerName, strings.Join(headerValue, ", "))
				}

			}
			if r.Method == http.MethodPost {
				fmt.Printf("server: got POST request\n")
				fmt.Printf("server: query id: %s\n", r.URL.Query().Get("id"))
				fmt.Printf("server: content-type: %s\n", r.Header.Get("content-type"))
				fmt.Printf("server: headers:\n")
				for headerName, headerValue := range r.Header {
					fmt.Printf("\t%s = %s\n", headerName, strings.Join(headerValue, ", "))
				}

				reqBody, err := io.ReadAll(r.Body)
				if err != nil {
					fmt.Printf("server: could not read request body: %s\n", err)
				}
				fmt.Printf("server: request body: %s\n", reqBody)

			}
		})
		mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("server: %s /health\n", r.Method)
			w.WriteHeader(http.StatusOK)
		})
		server := http.Server{
			Addr:    fmt.Sprintf(":%d", serverPort),
			Handler: mux,
		}
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				fmt.Printf("error running http server: %s\n", err)
			}
		}
	}()

	time.Sleep(1000 * time.Millisecond)

	// Make a GET request to the server
	requestURL := fmt.Sprintf("http://localhost:%d", serverPort)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body: %s\n", resBody)
	// print the custom header
	fmt.Printf("client: HEADER (Custom-Header) Value: %s\n", res.Header.Get("Custom-Header"))

	// Make a POST request to the server
	jsonBody := []byte(`{"client_message": "hello, server!"}`)
	bodyReader := bytes.NewReader(jsonBody)
   
	requestURL = fmt.Sprintf("http://localhost:%d?id=1234", serverPort)
	req, err = http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		fmt.Printf("client: could not create POST request: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("client: POST request created")
	fmt.Println("POST REQUEST : ", req)

}