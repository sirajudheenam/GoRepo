package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	utils "github.com/sirajudheenam/GoRepo/utils"
)
type Page struct {
    Title string
    Body  []byte
}
func (p *Page) save() error {
    filename := p.Title + ".txt"
    return os.WriteFile(filename, p.Body, 0600)
}
func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}
func savePage(/* title string, body []byte */ ) error {
	// filename := title + ".txt"
	// return os.WriteFile(filename, body, 0600)
	p1 := &Page{Title: "wiki", Body: []byte("This is a sample View Page.")}
    p1.save()
    p2, err := loadPage("wiki")
    fmt.Println(string(p2.Body))
	if err != nil {
		return err
	} else {
		return nil
	}
}

// func removeTrailSlash(s string) string {
// 	if s[len(s)-1:] == "/" {
// 		fmt.Println("\nremoving trail slash")
// 		return s[:len(s)-1]
// 	}
// 	return s
// }

func viewHandler(w http.ResponseWriter, r *http.Request) {
	pathLength := len("/view/")
	fmt.Println("pathLength: ", pathLength)
	// assign the title to the path after /view/
	title := r.URL.Path[len("/view/"):]
	fmt.Printf("\n r.URL: %s \n", r.URL)

	fmt.Printf("\n r.URL.Path: %s\n", r.URL.Path)
	fmt.Printf("\n r.URL.Path[1:]: %s\n", r.URL.Path[1:])

	// title := r.URL.Path[1:]
	if title == "" {
		fmt.Println("title is empty")
		os.Exit(1)
	} else {
		fmt.Println("title : ", title)
	}
	fmt.Printf("title: %s", title)
	cleanTitle := utils.RemoveTrailSlash(title)
	// the following will never work as it is private function of the package utils
	// cleanTitle := utils.removeTrailSlash(title)
	p, _ := loadPage(cleanTitle)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)

}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	_ = savePage()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}