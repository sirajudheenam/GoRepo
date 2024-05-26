package books

import (
	"strings"
)

type Category string

const (
	CategoryNovel      Category = "Novel"
	CategoryShortStory Category = "Short Story"
)

type Book struct {
	Title  string
	Author string
	Pages  int
}

// func (interface receiver) method_name() return_type {
func (b *Book) Category() Category {
	if b.Pages > 300 {
		return CategoryNovel
	}
	return CategoryShortStory
}

func (b *Book) IsValid() bool {
	return b.Title != "" && b.Author != "" && b.Pages > 0
}

func (b *Book) AuthorLastName() string {
	// split the author name by space
	// if there are no elements, return an empty string
	if b.Author == "" {
		return ""
	} else if strings.Count(b.Author, " ") == 0 {
		// if there is only one element and there are no spaces, return that element
		return b.Author
	} else {
		name := strings.Split(b.Author, " ")
		// if there are more than one element, return the last element
		return name[1];
	}
}

func (b *Book) AuthorFirstName() string {
		// split the author name by space
	// if there are no elements, return an empty string
	if b.Author == "" || strings.Count(b.Author, " ") == 0  {
		return ""
	} else {
		name := strings.Split(b.Author, " ")
		// if there are more than one element, return the last element
		return name[0];
	}
}
