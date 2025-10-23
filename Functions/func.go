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

func vara(st string, ar ...int) {
	fmt.Println(st)
	for id, val := range ar {
		fmt.Println(id, val)
	}
}

func argfunc(x func()) {
	x()
}

func ret() func() {
	return func() { fmt.Println("Function return") }
}

func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
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

	a := 10
	func() {
		fmt.Println("Func calling")
	}()

	k := func(x int) {
		fmt.Println("Func var ass", a, x)
	}

	k(20)

	vara("Senthil", 1, 2, 3, 5)

	retfunc := ret()
	retfunc()

	argfunc(retfunc)

	// closure and state

	c1 := makeCounter()
	c2 := makeCounter()

	fmt.Println(c1()) // 1
	fmt.Println(c1()) // 2
	fmt.Println(c2()) // 1
}
