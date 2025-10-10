package main

import "fmt"

// Add returns the sum of a and b.
func Add(a, b int) int {
	res := a + b
	return res
}

func main() {
	fmt.Println("Sum:", Add(3, 4))
}
