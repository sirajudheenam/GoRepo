package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"
)

// detectContentType guesses MIME based on extension
func DetectContentType(path string) string {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".pdf":
		return "application/pdf"
	case ".png":
		return "image/png"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".csv":
		return "text/csv"
	case ".txt":
		return "text/plain"
	default:
		return "application/octet-stream"
	}
}

func RequestToString(r *http.Request) string {
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
