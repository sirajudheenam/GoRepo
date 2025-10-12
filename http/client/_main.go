// package main

// import (
// 	"context"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"net/url"
// 	"strings"
// 	"time"
// )

// type myHttpClient struct {
// 	Client *http.Client
// }

// type myTransport struct {
// 	RoundTripper http.RoundTripper
// }

// func NewMyTransport() *myTransport {
// 	return &myTransport{
// 		RoundTripper: &http.Transport{
// 			MaxIdleConns:        10,
// 			MaxIdleConnsPerHost: 10,
// 			IdleConnTimeout:     30,
// 			DisableCompression:  true,
// 		},
// 	}
// }

// func NewMyHttpClient() *myHttpClient {
// 	return &myHttpClient{
// 		Client: &http.Client{
// 			Timeout:   15 * time.Second,
// 			Transport: NewMyTransport(),
// 		},
// 	}
// }

// func (c *myHttpClient) Do(req *http.Request) (*http.Response, error) {
// 	fmt.Println("Custom Do method called")
// 	return c.Client.Do(req)
// }

// func (c *myHttpClient) Get(url string) (*http.Response, error) {
// 	fmt.Println("Custom Get method called")
// 	return c.Client.Get(url)
// }

// func (c *myHttpClient) Post(url, contentType string, body io.Reader) (*http.Response, error) {
// 	fmt.Println("Custom Post method called")
// 	return c.Client.Post(url, contentType, body)
// }

// func (c *myHttpClient) Head(url string) (*http.Response, error) {
// 	fmt.Println("Custom Head method called")
// 	return c.Client.Head(url)
// }

// func (c *myHttpClient) Put(url string, contentType string, body io.Reader) (*http.Response, error) {
// 	fmt.Println("Custom Put method called")
// 	return c.Client.Put(url, contentType, body)
// }

// func (c *myHttpClient) Delete(url string) (*http.Response, error) {
// 	fmt.Println("Custom Delete method called")
// 	return c.Client.Delete(url)
// }

// func (c *myHttpClient) Patch(url string, contentType string, body io.Reader) (*http.Response, error) {
// 	fmt.Println("Custom Patch method called")
// 	req, err := http.NewRequest(http.MethodPatch, url, body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	req.Header.Set("Content-Type", contentType)
// 	return c.Client.Do(req)
// }

// func main() {

// 	const PORT = "8080"

// 	c := &NewmyHttpClient{
// 		Client: &http.Client{
// 			Timeout:   15 * time.Second,
// 			Transport: NewMyTransport(),
// 		},
// 	}

// 	resp, err := c.Get("http://localhost:" + PORT + "/")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	fmt.Println("Response Status:", resp.Status)
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Response Body [Get]:", string(body))

// 	// Create a context with a timeout
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	// Full URL: http://alice:secret@localhost:8000/api/v1/users?id=42&active=true#section%202
// 	myURL := &url.URL{
// 		Scheme:      "http",
// 		User:        url.UserPassword("alice", "secret"),
// 		Host:        "localhost:" + PORT,
// 		Path:        "/api/v1/users",
// 		RawPath:     "/api/v1/%75sers", // %75 = 'u'
// 		ForceQuery:  false,
// 		RawQuery:    "id=42&active=true", // The query string, without the ?.
// 		Fragment:    "section2",          // The fragment (the part after #), decoded.
// 		RawFragment: "section%202",
// 	}
// 	myRequest := &http.Request{
// 		Method:     http.MethodPost,
// 		URL:        myURL,
// 		Proto:      "HTTP/1.1",
// 		ProtoMajor: 1,
// 		ProtoMinor: 1,
// 		Header: http.Header{
// 			"User-Agent":      {"MyCustomClient/1.0"},
// 			"Accept":          {"application/json"},
// 			"X-Custom-Header": {"CustomValue"},
// 		},
// 		// Host: "localhost:" + PORT, // URL Host is used if Host is empty
// 		ContentLength:    int64(len("Hello, Server!")), // -1 if unknown
// 		TransferEncoding: []string{"chunked"},
// 		Close:            false,
// 		Trailer:          http.Header{},
// 		RemoteAddr:       "localhost:" + PORT,
// 		Body:             io.NopCloser(strings.NewReader("Hello, Server! FROM CLIENT")),
// 	}

// 	myRequest = myRequest.WithContext(ctx)

// 	resp, err = myHttpClient.Do(myRequest)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	fmt.Println("Response Status [Post]:", resp.Status)
// 	postBody, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Response Body [Post]:", string(postBody))
// }
