// inspiration: https://www.digitalocean.com/community/tutorials/how-to-make-http-requests-in-go
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


func MakeServer() {

	// make http server
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("[server][GET]: %s /\n", r.Method)
			if r.Method == http.MethodGet {
				w.Header().Add("X-Custom-HEADER", "BILLA")
				w.Header().Add("X-Happy-TAG", "VERY-HAPPY")

				id := r.URL.Query().Get("id")
				if id != "" {
				fmt.Printf("[server][GET]: query id: %s\n", id )
				} else {
					fmt.Printf("[server][GET]: query id: %s\n", "no id found" )
				}
				fmt.Printf("[server][GET]: content-type: %s\n", r.Header.Get("content-type"))
				fmt.Printf("[server][GET]: headers:\n")
				for headerName, headerValue := range r.Header {
					fmt.Printf("\t%s = %s\n", headerName, strings.Join(headerValue, ", "))
				}
				
				w.Write([]byte("Hello, client!..  How are you?.."))
			}
			if r.Method == http.MethodPost {
				fmt.Printf("[server][POST]: got POST request\n")
				fmt.Printf("[server][POST]: query id: %s\n", r.URL.Query().Get("id"))
				fmt.Printf("[server][POST]: content-type: %s\n", r.Header.Get("content-type"))
				w.Header().Add("X-First-Name", "Sirajudheen")
				w.Header().Add("X-Last-Name", "Mohamed Ali")
				fmt.Printf("[server][POST]: headers:\n")
				for headerName, headerValue := range r.Header {
					fmt.Printf("\t%s = %s\n", headerName, strings.Join(headerValue, ", "))
				}

				// The caller must close the response body when finished with it
				defer r.Body.Close()

				reqBody, err := io.ReadAll(r.Body)
				if err != nil {
					fmt.Printf("[server][POST]: could not read request body: %s\n", err)
				}
				fmt.Printf("[server][POST]: request body: %s\n", reqBody)

			}
		})
		mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("[server]: %s /health\n", r.Method)
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
}

func MakeGETRequest() {

	// Make a GET request to the server
	requestURL := fmt.Sprintf("http://localhost:%d?id=0056", serverPort)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	req.Header.Add("REQ-HEADER-MOOD", "HAPPY-AS-ALWAYS")
	req.Header.Add("REQ-HEADER-ATTITUDE", "POSITIVE-AS-ALWAYS")
	if err != nil {
		fmt.Printf("[client][GET]: Could not create request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("[client][GET]: error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("[client][GET]: Status code: %d\n", res.StatusCode)

	fmt.Printf("[client][GET]: Response headers:\n")
	for headerName, headerValue := range res.Header {
		fmt.Printf("\t%s = %s\n", headerName, strings.Join(headerValue, ", "))
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("[client][GET]: Could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("[client][GET]: Response body: %s\n", resBody)
}

func MakePOSTRequest(requestURL string, jsonBody []byte) {
	// Make a POST request to the server
	if requestURL == "" {	
		fmt.Printf("[client][POST]: RequestURL is empty\n")
		os.Exit(1)
	}
	if len(jsonBody) == 0 {
		fmt.Printf("[client][POST]: jsonBody is empty\n")
		os.Exit(1)
	}
	bodyReader := bytes.NewReader(jsonBody)
   
	
	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		fmt.Printf("[client][POST]: Could not create POST request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("[client][POST]: Error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("[client][POST]: Status code: %d\n", res.StatusCode)

	fmt.Printf("[client][POST]: Response headers:\n")
	for headerName, headerValue := range res.Header {
		fmt.Printf("\t%s = %s\n", headerName, strings.Join(headerValue, ", "))
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: Could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("[client][POST]: Response body: %s\n", resBody)
	// print the custom header
	fmt.Printf("[client][POST]: Name: X-Custom-HEADER: Value: %s\n", res.Header.Get("X-Custom-HEADER"))

	fmt.Println("[client][POST]: Request created")
	fmt.Println("POST REQUEST : ", req.URL)
}
func main() {

	requestURL := fmt.Sprintf("http://localhost:%d?id=1234", serverPort)
	jsonBody := []byte(`{"client_message": "hello, server!"}`)

	MakeServer()

	time.Sleep(3000 * time.Millisecond)

	MakeGETRequest()

	MakePOSTRequest(requestURL, jsonBody)

}