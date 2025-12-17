package main

import (
	"fmt"

	"rsc.io/quote"

	"example.com/greetings"
)

func main() {
	// Print a quote from rsc.io/quote module.
	fmt.Println(quote.Go())

	// Get a greeting message and print it from local example.com/greetings module.
	message := greetings.Hello("Bruno")
	fmt.Println(message)
}
