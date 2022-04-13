package main

import (
	"fmt" ///importing this package, handles input and output text
	"log"

	"rsc.io/quote" //importing this package

	"example.com/greetings" //importing this package
)

//declare a main package
//code executed as an application must be in a main package
func main () {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	//Get a greeting message and print it out
	message := greetings.Hello("Gladys")
	fmt.Println(message)

	// Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

	message, err := greetings.HelloImproved("danny")
	if err != nil {
		log.Fatal(err)
	}

	//if no error was returned , print the returned message to console
	fmt.Println(message)

	messageRandom, error := greetings.HelloRandom("Marcelo")
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(messageRandom)

	//a slice of names
	names := []string{"Mary", "Kim", "Martha", "Dan"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

	
}

//Go code is grouped into pacakges and packages are grouped into modules
//go mod init command creates a go.mod file to track code dependencies