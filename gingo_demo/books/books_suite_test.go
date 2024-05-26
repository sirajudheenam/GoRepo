package books_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sirajudheenam/GoRepo/gingo_demo/books"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Books Suite")
}

var _ = Describe("Books", func() {
	var foxInSocks, lesMis *books.Book
  
	// Assigning the books to the variables before each test
	BeforeEach(func() {
	  lesMis = &books.Book{
		Title:  "Les Miserables",
		Author: "Victor Hugo",
		Pages:  2783,
	  }
  
	  foxInSocks = &books.Book{
		Title:  "Fox In Socks",
		Author: "Dr. Seuss",
		Pages:  24,
	  }

	})


	Describe("Categorizing books", func() {
	  Context("with more than 300 pages", func() {
		It("should be a novel", func() {
		  Expect(lesMis.Category()).To(Equal(books.CategoryNovel))
		})
	  })
  
	  Context("with fewer than 300 pages", func() {
		It("should be a short story", func() {
		  Expect(foxInSocks.Category()).To(Equal(books.CategoryShortStory))
		})
	  })
	})
  })
  
  var _ = Describe("Books 2", func() {
	var oliverTwist *books.Book
  
	BeforeEach(func() {
	  oliverTwist = &books.Book{
		Title: "Oliver Twist",
		Author: "Charles Dickens",
		Pages: 534,
	  }
	  Expect(oliverTwist.IsValid()).To(BeTrue())
	})
  
	It("can extract the author's last name", func() {
		Expect(oliverTwist.AuthorLastName()).To(Equal("Dickens"))
	  })
	
	  It("interprets a single author name as a last name", func() {
		oliverTwist.Author = "Dickens"
		Expect(oliverTwist.AuthorLastName()).To(Equal("Dickens"))
	  })
	
	  It("can extract the author's first name", func() {
		Expect(oliverTwist.AuthorFirstName()).To(Equal("Charles"))
	  })
	
	  It("returns no first name when there is a single author name", func() {
		oliverTwist.Author = "Dickens"
		Expect(oliverTwist.AuthorFirstName()).To(BeZero()) //BeZero asserts the value is the zero-value for its type.  In this case: ""
	  })
  })
  