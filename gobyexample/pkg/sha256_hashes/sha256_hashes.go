package sha256hashes

// SHA256 hashes are frequently used to compute short identities for binary or text blobs.
// For example, TLS/SSL certificates use SHA256 to compute a certificate’s signature.
// Here’s how to compute SHA256 hashes in Go.

// Go implements several hash functions in various crypto/* packages.

import (
	"crypto/sha256"
	"fmt"
)

func Run() {

	fmt.Println("\nSHA256 Hashes")
	s := "sha256 this string"

	// Here we start with a new hash.

	h := sha256.New()

	// Write expects bytes. If you have a string s, use []byte(s) to coerce it to bytes.

	h.Write([]byte(s))

	// This gets the finalized hash result as a byte slice. The argument to Sum can be used to append to an existing byte slice: it usually isn’t needed.

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}

// Running the program computes the hash and prints it in a human-readable hex format.

// go run main.go
// SHA256 Hashes
// sha256 this string
// 1af1dfa857bf1d8814fe1af8983c18080019922e557f15a8a0d3db739d77aacb

// You can compute other hashes using a similar pattern to the one shown above. For example, to compute SHA512 hashes import crypto/sha512 and use sha512.New().

// Note that if you need cryptographically secure hashes, you should carefully research hash strength!
