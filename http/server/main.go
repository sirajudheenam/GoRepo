package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

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

// Download handler (dynamic)
func downloadHandler(w http.ResponseWriter, r *http.Request) {
	reqStr := utils.RequestToString(r)
	fmt.Println(reqStr)

	// Parse query ?file=filename
	queryParams, _ := url.ParseQuery(r.URL.RawQuery)
	fileName := queryParams.Get("file")

	if fileName == "" {
		http.Error(w, "Missing 'file' query parameter", http.StatusBadRequest)
		return
	}

	// ✅ Security: Prevent directory traversal (e.g., "../../etc/passwd")
	fileName = filepath.Base(fileName)

	// Define safe base folder
	baseDir := "./files"

	// Construct full path safely
	filePath := filepath.Join(baseDir, fileName)

	// Check if file exists
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	if info.IsDir() {
		http.Error(w, "Cannot download a directory", http.StatusBadRequest)
		return
	}

	// Detect MIME type (optional improvement)
	contentType := utils.DetectContentType(filePath)
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", fileName))

	// Serve file
	http.ServeFile(w, r, filePath)
}

// wget --debug -O my_downloaded_file.pdf "http://localhost:8080/download?file=example.pdf"
// curl -v -o my_downloaded_file.pdf "http://localhost:8080/download?file=example.pdf"

// Simple
// // Download handler
// func downloadHandler(w http.ResponseWriter, r *http.Request) {
// 	reqStr := requestToString(r)
// 	fmt.Println(reqStr) // log to console

// 	// Path to file you want to send
// 	filePath := "./files/example.pdf"

// 	// Set headers for download
// 	w.Header().Set("Content-Disposition", "attachment; filename=\"example.pdf\"")
// 	w.Header().Set("Content-Type", "application/pdf")

// 	// Option 1: Using http.ServeFile (simple and efficient)
// 	http.ServeFile(w, r, filePath)
// }

func main() {
	// --- Register routes ---
	http.HandleFunc("/api/v1/users", apiHandler)
	http.HandleFunc("/download", downloadHandler)
	http.HandleFunc("/", handler)

	// --- Create servers ---
	httpSrv := &http.Server{Addr: ":8080", Handler: nil}

	certsDir := "certs/"
	caCert, err := os.ReadFile(certsDir + "ca.crt")
	if err != nil {
		log.Fatalf("Failed to read CA certificate: %v", err)
	}
	caPool := x509.NewCertPool()
	caPool.AppendCertsFromPEM(caCert)

	tlsCfg := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  caPool,
	}
	httpsSrv := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsCfg,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.TLS != nil {
				for _, cert := range r.TLS.PeerCertificates {
					log.Printf("Client Cert Subject: %s", cert.Subject)
				}
			}
			fmt.Fprintln(w, "Hello, mutual TLS!")
		}),
	}

	// --- Start HTTP server in background ---
	go func() {
		fmt.Println("HTTP server listening on :8080")
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	// --- Optional: Redirect HTTP to HTTPS ---
	// Uncomment the following block to enable HTTP to HTTPS redirection
	// Note: Ensure that the HTTPS server is running on port 8443 before enabling this.
	// go func() {
	// 	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		target := "https://" + r.Host + r.URL.String()
	// 		http.Redirect(w, r, target, http.StatusMovedPermanently)
	// 	})))
	// }()

	// --- Start HTTPS server in background ---
	go func() {
		fmt.Println("HTTPS server listening on :8443")
		if err := httpsSrv.ListenAndServeTLS("certs/server.crt", "certs/server.key"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTPS server error: %v", err)
		}
	}()

	// --- Wait for interrupt signal ---
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	fmt.Println("\nShutting down servers...")

	// --- Graceful shutdown with timeout ---
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpSrv.Shutdown(ctx); err != nil {
		log.Printf("HTTP shutdown error: %v", err)
	}
	if err := httpsSrv.Shutdown(ctx); err != nil {
		log.Printf("HTTPS shutdown error: %v", err)
	}

	fmt.Println("Servers stopped gracefully.")
}
