package main

import (
	"fmt"
)

func main() {
	x := 10
	y := 3
	fmt.Println("x + y =", x+y, x*y, x-y, x/y, x%y)
	fmt.Println(x == y, x != y, x > y, x < y, x >= y, x <= y)
	fmt.Println(2 << 1) // 2 * 2^1 = 4
	fmt.Println(3 >> 1) // 3 / 2^1 = 1

	// 0011 (3)

	xx := 3 // 0011

	yy := &xx // 0011
	fmt.Println(xx, *yy)
	*yy = 5
	fmt.Println(xx, *yy)

}
