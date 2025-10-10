package main

import "fmt"

// Function with no parameters and no return value
func sayHello() {
	fmt.Println("Hello, World!")
}

// Function with parameters and no return value
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// Function with parameters and a return value
func add(a int, b int) int {
	return a + b
}

// Function with multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

//named return values
func namedReturn(a, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

func main() {
	sayHello()
	greet("Senthil")
	sum := add(3, 5)
	fmt.Println("Sum:", sum)
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division Result:", result)
	}
	namedSum, namedDiff := namedReturn(10, 5)
	fmt.Println("Named Return - Sum:", namedSum, "Diff:", namedDiff)
	fmt.Println(namedReturn(20, 10))
}
