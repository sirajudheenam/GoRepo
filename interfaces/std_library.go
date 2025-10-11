package main

import (
	"fmt"
	"io"
	"os"
)

// // io.Reader and io.Writer are interfaces
// func readFile(r io.Reader) {
// 	buf := make([]byte, 8)
// 	n, _ := r.Read(buf)
// 	fmt.Println("Read bytes:", string(buf[:n]))
// }

// os.File implements both io.Reader and io.Writer
func writeFile(w io.Writer, data string) {
	n, _ := w.Write([]byte(data))
	fmt.Println("Wrote bytes:", n)
}

func main() {
	// readFile(os.Stdin)                    // os.Stdin implements io.Reader
	writeFile(os.Stdout, "Hello, World!\n") // os.Stdout implements io.Writer

	os.Stdout.WriteString("Writing directly to stdout \n")

	myFile := "example.txt"
	
	// delete file after program ends
	defer os.Remove(myFile)

	// check if file exists, if not create it
	if _, err := os.Stat(myFile); os.IsNotExist(err) {
		_, err := os.Create(myFile)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
	}
	f, err := os.Create(myFile) // os.Create returns *os.File which implements io.Writer
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close()

	writeFile(f, "Data written to file.\n")

	// os.File also implements io.Reader

	fileInfo, err := os.Stat(myFile)
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("File size:", fileInfo.Size())
	fmt.Println("Is directory:", fileInfo.IsDir())

	read, err := os.ReadFile(myFile) // os.ReadFile uses io.Reader internally
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content:", string(read))
}
