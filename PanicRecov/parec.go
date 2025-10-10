package main

import (
	"fmt"
)

func mayPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Simulate a panic
	panic("something went wrong")
}

func main() {
	fmt.Println("Starting program")
	mayPanic()
	fmt.Println("Program continues after recovery")
}
