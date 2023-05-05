package main

import (
	"fmt"

	"github.com/if1bonacci/go-design-patterns/internal/services/decorator"
)

func main() {
	book := decorator.NewBook(100, "Book name")
	bookWithTax := decorator.NewBookWithTax(book)

	fmt.Println(bookWithTax.GetName())
	fmt.Println(bookWithTax.GetPrice())
}
