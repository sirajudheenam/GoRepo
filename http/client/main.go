package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

const PORT = "8080"

// Custom HTTP Client Wrapper
type MyHTTPClient struct {
	Client *http.Client
}

// Custom Transport with sane defaults
type MyTransport struct {
	RoundTripper http.RoundTripper
}

func NewMyTransport() http.RoundTripper {
	return &http.Transport{
		MaxIdleConns:        10,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     30 * time.Second,
		DisableCompression:  true,
	}
}

func NewMyHTTPClient() *MyHTTPClient {
	return &MyHTTPClient{
		Client: &http.Client{
			Timeout:   15 * time.Second,
			Transport: NewMyTransport(),
		},
	}
}

// Custom wrappers around HTTP methods
func (c *MyHTTPClient) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
	fmt.Println("Request ID:", ctx.Value(requestIDKey), "GET:", req.URL.String())
	req.Header.Set("User-Agent", "MyCustomClient/1.0")
	req = req.WithContext(ctx)
	return c.Client.Do(req)
}

func (c *MyHTTPClient) Get(ctx context.Context, url string) (*http.Response, error) {
	fmt.Println("Request ID:", ctx.Value(requestIDKey), "GET:", url)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "MyCustomClient/1.0")
	return c.Do(ctx, req)
}

func (c *MyHTTPClient) Post(ctx context.Context, url, contentType string, body io.Reader) (*http.Response, error) {
	fmt.Println("Request ID:", ctx.Value(requestIDKey), "POST:", url)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("User-Agent", "MyCustomClient/1.0")
	return c.Do(ctx, req)
}

func (c *MyHTTPClient) Head(ctx context.Context, url string) (*http.Response, error) {
	fmt.Println("Request ID:", ctx.Value(requestIDKey), "HEAD:", url)
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "MyCustomClient/1.0")
	req = req.WithContext(ctx)
	return c.Do(ctx, req)
}

func (c *MyHTTPClient) Put(ctx context.Context, url, contentType string, body io.Reader) (*http.Response, error) {
	fmt.Println("➡️  Custom PUT called:", url)
	fmt.Println("Request ID:", ctx.Value(requestIDKey), "PUT:", url)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("User-Agent", "MyCustomClient/1.0")
	req = req.WithContext(ctx)
	return c.Do(ctx, req)
}

func (c *MyHTTPClient) Delete(ctx context.Context, url string) (*http.Response, error) {
	fmt.Println("➡️  Custom DELETE called:", url)
	fmt.Println("Request ID:", ctx.Value(requestIDKey), "DELETE:", url)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "MyCustomClient/1.0")
	return c.Do(ctx, req)
}

func (c *MyHTTPClient) Patch(ctx context.Context, url, contentType string, body io.Reader) (*http.Response, error) {
	fmt.Println("➡️  Custom PATCH called:", url)
	fmt.Println("Request ID:", ctx.Value(requestIDKey), "PATCH:", url)
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("User-Agent", "MyCustomClient/1.0")
	return c.Do(ctx, req)
}

// Download a file from server and save to disk
func (c *MyHTTPClient) DownloadFile(ctx context.Context, url, dest string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "MyCustomClient/1.0")
	fmt.Println("Request ID:", ctx.Value(requestIDKey), "Downloading:", url)

	resp, err := c.Do(ctx, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

// // 🧩 Simple test HTTP server
// func startTestServer() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Received %s request at %s\n", r.Method, r.URL.Path)
// 	})

// 	http.HandleFunc("/api/v1/users", func(w http.ResponseWriter, r *http.Request) {
// 		body, _ := io.ReadAll(r.Body)
// 		fmt.Printf("📥 Server received [%s]: %s\n", r.Method, string(body))
// 		fmt.Fprintf(w, "Response from /api/v1/users for %s\n", r.Method)
// 	})

// 	go func() {
// 		fmt.Println("🚀 Server running at http://localhost:" + PORT)
// 		if err := http.ListenAndServe(":"+PORT, nil); err != nil {
// 			panic(err)
// 		}
// 	}()
// }

type contextKey int

const (
	requestIDKey contextKey = iota
)

// 🧪 Main: test all methods
func main() {
	// startTestServer()
	time.Sleep(1 * time.Second) // Give server a moment to start

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx = context.WithValue(ctx, requestIDKey, uuid.New().String())
	defer cancel()

	client := NewMyHTTPClient()

	// === 1. Test GET ===
	resp, err := client.Get(ctx, "http://localhost:"+PORT+"/")
	if err != nil {
		fmt.Println("❌ GET Error:", err)
		return
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println("✅ GET Response:", string(body))

	// === 2. Test POST ===
	resp, err = client.Post(ctx, "http://localhost:"+PORT+"/api/v1/users", "text/plain", strings.NewReader("Hello from POST"))
	if err != nil {
		fmt.Println("❌ POST Error:", err)
		return
	}
	postBody, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println("✅ POST Response:", string(postBody))

	// === 3. Test PUT ===
	resp, err = client.Put(ctx, "http://localhost:"+PORT+"/api/v1/users", "application/json", strings.NewReader(`{"update":"true"}`))
	if err != nil {
		fmt.Println("❌ PUT Error:", err)
		return
	}
	putBody, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println("✅ PUT Response:", string(putBody))

	// === 4. Test PATCH ===
	resp, err = client.Patch(ctx, "http://localhost:"+PORT+"/api/v1/users", "application/json", strings.NewReader(`{"patch":"yes"}`))
	if err != nil {
		fmt.Println("❌ PATCH Error:", err)
		return
	}
	patchBody, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println("✅ PATCH Response:", string(patchBody))

	// === 5. Test DELETE ===
	resp, err = client.Delete(ctx, "http://localhost:"+PORT+"/api/v1/users")
	if err != nil {
		fmt.Println("❌ DELETE Error:", err)
		return
	}
	deleteBody, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println("✅ DELETE Response:", string(deleteBody))

	// === 6. Test HEAD ===
	resp, err = client.Head(ctx, "http://localhost:"+PORT+"/api/v1/users")
	if err != nil {
		fmt.Println("❌ HEAD Error:", err)
		return
	}
	resp.Body.Close()
	fmt.Println("✅ HEAD Response Headers:", resp.Header)

	// === 7. Custom Do() with context ===
	myURL := &url.URL{
		Scheme: "http",
		Host:   "localhost:" + PORT,
		Path:   "/api/v1/users",
	}
	req := &http.Request{
		Method: http.MethodPost,
		URL:    myURL,
		Header: http.Header{
			"User-Agent": {"MyCustomClient/1.0"},
		},
		Body: io.NopCloser(strings.NewReader("Custom Do() test")),
	}

	resp, err = client.Do(ctx, req)
	if err != nil {
		fmt.Println("❌ Do() Error:", err)
		return
	}
	doBody, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println("✅ Custom Do() Response:", string(doBody))

	// === 8. Test DownloadFile ===
	err = client.DownloadFile(ctx, "http://localhost:"+PORT+"/download?file=example.pdf", "example.pdf")
	if err != nil {
		fmt.Println("❌ DownloadFile Error:", err)
		return
	}
	fmt.Println("✅ DownloadFile Success: example.pdf downloaded")
}
