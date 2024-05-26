package main

import (
	"fmt"

	"github.com/sirajudheenam/GoRepo/gingo_demo/books"
)

func main() {
	book := &books.Book{
		Title: "Doomsday Conspiracy",
		Author: "Sidney Sheldon",
		Pages: 400,
	  }
	fmt.Println("book.IsValid() = ", book.IsValid())
	fmt.Println("book.Pages = ", book.Pages)
	fmt.Println("book.Title = ", book.Title)
	fmt.Println("book.Author = ", book.Author)
	
}