package main

import (
	"fmt"

	"github.com/if1bonacci/go-design-patterns/internal/services/decorator"
	"github.com/if1bonacci/go-design-patterns/internal/services/dependencyinjection"
	"github.com/if1bonacci/go-design-patterns/internal/services/factorymethod"
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

	//factory
	barsa, _ := factorymethod.GetTeam("barsa")
	kalmar, _ := factorymethod.GetTeam("kalmar")

	fmt.Printf("Team: %s", barsa.GetName())
	fmt.Println()
	fmt.Printf("Team: %s", kalmar.GetName())
}
