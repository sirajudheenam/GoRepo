package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {

	const PORT = "8080"
	myClient := &http.Client{}
	myClient.Transport = &http.Transport{
		MaxIdleConns:        10,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     30,
		DisableCompression:  true,
	}

	resp, err := myClient.Get("http://localhost:" + PORT + "/")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response Body [Get]:", string(body))

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Full URL: http://alice:secret@localhost:8000/api/v1/users?id=42&active=true#section%202
	myURL := &url.URL{
		Scheme:      "http",
		User:        url.UserPassword("alice", "secret"),
		Host:        "localhost:" + PORT,
		Path:        "/api/v1/users",
		RawPath:     "/api/v1/%75sers", // %75 = 'u'
		ForceQuery:  false,
		RawQuery:    "id=42&active=true", // The query string, without the ?.
		Fragment:    "section2",          // The fragment (the part after #), decoded.
		RawFragment: "section%202",
	}
	myRequest := &http.Request{
		Method:     http.MethodPost,
		URL:        myURL,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header: http.Header{
			"User-Agent":      {"MyCustomClient/1.0"},
			"Accept":          {"application/json"},
			"X-Custom-Header": {"CustomValue"},
		},
		// Host: "localhost:" + PORT, // URL Host is used if Host is empty
		ContentLength:    int64(len("Hello, Server!")), // -1 if unknown
		TransferEncoding: []string{"chunked"},
		Close:            false,
		Trailer:          http.Header{},
		RemoteAddr:       "localhost:" + PORT,
		Body:             io.NopCloser(strings.NewReader("Hello, Server! FROM CLIENT")),
	}

	myRequest = myRequest.WithContext(ctx)

	resp, err = myClient.Do(myRequest)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status [Post]:", resp.Status)
	postBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response Body [Post]:", string(postBody))
}
