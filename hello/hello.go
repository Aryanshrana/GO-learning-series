package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//requesting a greeting message
	message, err := greetings.Hello("Aryansh")
	//if error was return print to console and exit the program
	if err != nil {
		log.Fatal(err)
	}
	//if no error print the return message

	fmt.Println(message)
}
