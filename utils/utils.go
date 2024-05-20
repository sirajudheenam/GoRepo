package utils

import "fmt"

// RemoveTrailSlash removes the trailing slash from a string
func RemoveTrailSlash(s string) (string) {
	if s[len(s)-1:] == "/" {
		fmt.Println("\nremoving trail slash")
		return s[:len(s)-1]
	} 
	return s
}

// private function which removes the trailing slash from a string
// it will never be called as it is private function of the package utils
func removeTrailSlash(s string) (string) {
	if s[len(s)-1:] == "/" {
		fmt.Println("\nremoving trail slash")
		return s[:len(s)-1]
	} 
	return s
}
