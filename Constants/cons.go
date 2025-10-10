package main

import (
	"fmt"
)

func main() {
	const (
		a = iota + 2 // 0
		b
		c
	)

	fmt.Println(a, b, c) // 0 1 2
}
