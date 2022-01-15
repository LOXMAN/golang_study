package main

import (
	"fmt"
	"log"

	"rsc.io/quote"

	"golang_study/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	fmt.Println("Hello World!")
	fmt.Println(quote.Go())
	message, err := greetings.Hello("Gladys")

	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err2 := greetings.Hellos(names)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(messages)
}
