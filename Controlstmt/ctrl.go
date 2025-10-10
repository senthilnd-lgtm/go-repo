package main

import (
	"fmt"
	"strings"
)

func main() {

	defer fmt.Println("Deferred statement executed")
	// If statement

	if x := 10; x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is 5 or less")
	}

	// Switch statement
	y := 2
	switch y {
	case 1:
		fmt.Println("y is 1")
	case 2:
		fmt.Println("y is 2")
	default:
		fmt.Println("y is something else")
	}

	// For loop
	for i := 0; i < 3; i++ {
		fmt.Println("i =", i)
	}

	j := 0
	for j < 3 {
		fmt.Println("j =", j)
		j++
	}

	k := 0
outterlable:
	for k < 5 {
		if k == 3 {
			k++
			goto outterlable

		}
		fmt.Println("K =", k)
		k++
	}

	// Infinite loop with break
	n := 0

	for {
		if n >= 3 {
			break
		}
		n++
		fmt.Println("n =", n)
	}

	// looping over a collection
	numbers := []int{1, 2, 3}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// looping over a map
	person := map[string]string{"name": "Alice", "city": "Wonderland"}
	for key, value := range person {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	for key := range person {
		fmt.Printf("Key: %s\n", key)
	}

	// strings functions
	s := "Hello, World!"
	fmt.Println(strings.Contains(s, "World"))
	fmt.Println(strings.HasPrefix(s, "Hello"))
	fmt.Println(strings.Index(s, "o"))
	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.ReplaceAll(s, "o", "0"))

	// string functions with replace example
	original := "foo bar foo"
	replaced := strings.Replace(original, "foo", "baz", 1)
	fmt.Println(replaced) // Output: baz bar foo
	// fmt.Println(strings.Title(original)) depricated
	fmt.Println(strings.ToTitle(original))
	fmt.Println(strings.TrimSpace("  Hello World  "))
	fmt.Println(strings.Split(original, " "))
	fmt.Println(len(original))
	fmt.Println(strings.Repeat("Go! ", 3))
}
