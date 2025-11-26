package main

import (
	"fmt"
	"strings"
)

func isupper(x *string) bool {
	fmt.Println(" Address of x", x)
	fmt.Println("Value of x", *x)
	if strings.ToUpper(*x) == *x {
		return true
	}
	return false
}

func main() {
	// Declare a variable
	var x int = 10

	// Declare a pointer to x
	var p *int = &x

	var c = &x
	// Print the value and the pointer
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", p, c)
	fmt.Println("Value pointed by p:", *p)

	// Change value using pointer
	*p = 20
	fmt.Println("New value of x:", x)

	usenew := new(int) // allocates memory for an int and returns a pointer to it
	*usenew = 30
	fmt.Println("Value of usenew:", usenew, *usenew)
	tt := usenew
	*tt = 77
	fmt.Println("Value of usenew:", usenew, *usenew)

	msg := "HELLO"
	var px *string = &msg

	fmt.Println(isupper(px))

}
