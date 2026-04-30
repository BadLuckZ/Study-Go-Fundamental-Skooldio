package main

import (
	"fmt"

	"github.com/BadLuckZ/Study-Go-Fundamental-Skooldio/package/book"
)

func main() {
	b := book.Book{
		Title:  "Book1",
		Author: "Author1",
	}

	fmt.Println(b)
}
