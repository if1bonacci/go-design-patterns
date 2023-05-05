package main

import (
	"fmt"

	"github.com/if1bonacci/go-design-patterns/internal/services/decorator"
	"github.com/if1bonacci/go-design-patterns/internal/services/dependencyinjection"
)

func main() {

	//dercorator
	book := decorator.NewBook(100, "Book name")
	bookWithTax := decorator.NewBookWithTax(book)

	fmt.Println(bookWithTax.GetName())
	fmt.Println(bookWithTax.GetPrice())

	//dependency injection
	cheeze := dependencyinjection.NewCheeze()
	pizza := dependencyinjection.NewPizza(cheeze)

	fmt.Println(pizza.GetCheeze())
}
