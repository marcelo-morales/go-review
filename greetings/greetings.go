package greetings //declare a greetings package

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//implement a Hello function to return the greeting
//Hello returns a greeting for the named person
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message

	// ^ this is the same as
	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)
}

//In Go, a function whose name starts with a capital letter can be called by a 
//function not in the same package -> known as an exported name

//:= declaring and intializing a variable in one line

//this improved function returns two values, a string and an error

func HelloImproved (name string) (string, error) {
	//if no name is given, return an error message
	if name ==""{
		return "", errors.New("empty name")
	}

	//if a name is receieved, return a value that embeds the name
	//in a greeting message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
	//adding nil, meaning no error, as a second value to mean a successful return
}

func HelloRandom (name string) (string, error) {
	if name=="" {
		return "", errors.New("empty name given")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil

}

//sets initial values for variables used in the function
func init () {
	rand.Seed(time.Now().UnixNano())
}

//randomFormat returns one of a set of greeting messages, returned message 
//is selected at random

//this function starts with a lowercase letter so it is only accessible from in its own package
func randomFormat() string{

	//size of the slize is omitted so the size of the array underlying slice can be dynamically changed
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met",
	}

	return formats[rand.Intn(len(formats))]
}

func Hellos (names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := HelloImproved(name)
		if err != nil {
			return nil, err
		}

		//in the map, associate the retrieved messages with the name
		messages[name] = message
	}
	return messages, nil
}

