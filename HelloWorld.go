package main

import (
	"fmt"
	"math"
)

func add(a, b int) (x, y int) {
	x = a
	y = b
	return
}

func main() {
	num1 := 5
	num2 := 7
	//m := add(num1, num2)
	//t.Printf("Sum of %d and %d is %d\n", num1, num2, sum)
	fmt.Println(add(num1, num2))
	fmt.Println(math.Pi)

	const myconst = 1
	fmt.Printf("type : %T, value :%v", myconst+1.1, myconst+1.1)

	const (
		a = iota + 2 // 0
		b
		c
	)

	fmt.Println(a, b, c) // 0 1 2

}
