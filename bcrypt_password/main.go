package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

func main() {
	// Define command-line flags
	encryptFlag := flag.Bool("encrypt", false, "Encrypt (hash) the given password")
	decryptFlag := flag.Bool("decrypt", false, "Decrypt (compare) the password with a hash")
	hashFlag := flag.String("hash", "", "Provide an existing bcrypt hash for comparison")
	costFlag := flag.Int("cost", 12, "Bcrypt cost factor (default 12)")

	flag.Parse()

	if !*encryptFlag && !*decryptFlag {
		log.Fatal("Please specify either -encrypt or -decrypt flag.")
	}

	if *encryptFlag && *decryptFlag {
		log.Fatal("Please use only one flag: either -encrypt or -decrypt.")
	}

	// --- Encryption mode ---
	if *encryptFlag {
		fmt.Print("Enter password to hash: ")
		password, err := readPassword()
		if err != nil {
			log.Fatalf("Error reading password: %v", err)
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(password), *costFlag)
		if err != nil {
			log.Fatalf("Error generating hash: %v", err)
		}

		fmt.Println("\nGenerated bcrypt hash:")
		fmt.Println(string(hash))
		return
	}

	// --- Decryption (verification) mode ---
	if *decryptFlag {
		if *hashFlag == "" {
			fmt.Print("Enter bcrypt hash to verify against: ")
			reader := bufio.NewReader(os.Stdin)
			h, _ := reader.ReadString('\n')
			*hashFlag = strings.TrimSpace(h)
		}

		fmt.Print("Enter password to verify: ")
		password, err := readPassword()
		if err != nil {
			log.Fatalf("Error reading password: %v", err)
		}

		err = bcrypt.CompareHashAndPassword([]byte(*hashFlag), []byte(password))
		fmt.Println() // new line after hidden input
		if err != nil {
			fmt.Println("❌ Password does NOT match the hash.")
			os.Exit(1)
		}
		fmt.Println("✅ Password matches the hash!")
	}
}

// readPassword hides input when typing
func readPassword() (string, error) {
	bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(bytePassword)), nil
}


// package main

// import (
//     "fmt"

//     "golang.org/x/crypto/bcrypt"
// )

// func main() {
//     password := []byte("concourse")
//     hash, _ := bcrypt.GenerateFromPassword(password, 12)
//     fmt.Println(string(hash))
// }