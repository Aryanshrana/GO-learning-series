package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	//in Go , functions whose name starts with capital letters can be called from functions who are not in package
	//error handling
	//if no name is given then returns a error with a message
	if name == "" {
		return name, errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

//init sets initial values for variables used in function
func init() {
	rand.Seed(time.Now().UnixNano())
}

//randomFormat function select the already selected greetings in random manner
func randomFormat() string {
	formats := []string {
		"hi %v, glad to see you",
		"Hi %v, welcome",
		"%v, hail to you",
	}

	//return a randomly select message format by specifying a random index for slice of formats
	return formats[rand.Intn(len(formats))]
}