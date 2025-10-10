package main

import (
	"errors"
	"fmt"
	"os"
)

func doSomething() error {
	// Simulate an error
	return errors.New("something went wrong")
}

func main() {
	f, errr := os.Open("C:\\Senthil\\GoLang\\WS\\HelloWorld.go")
	if errr != nil {

		if errors.Is(errr, os.ErrNotExist) {
			fmt.Println("File does not exist")
		} else {
			fmt.Println("Some other error occurred:", errr)
		}
		// Handle error appropriately
		return
	}
	fmt.Println("File opened successfully", f.Name())

	err := doSomething()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		// Handle error appropriately
		return
	}
	fmt.Println("Operation successful")
}
