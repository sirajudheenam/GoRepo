package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/sirajudheenam/GoRepo/http/server/pkg/utils"
)

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
