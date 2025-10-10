package main

import "fmt"

func main() {
	// Create a map with string keys and int values
	myMap := make(map[string]int)

	// Add key-value pairs to the map
	myMap["apple"] = 5
	myMap["banana"] = 3

	// Access a value by key
	fmt.Println("apple:", myMap["apple"])

	// Check if a key exists
	value, exists := myMap["orange"]
	if exists {
		fmt.Println("orange:", value)
	} else {
		fmt.Println("orange not found")
	}

	// Iterate over the map
	for key, value := range myMap {
		fmt.Printf("%s: %d\n", key, value)
	}

	// Delete a key from the map
	delete(myMap, "banana")

	mapliteral := map[string]int{"one": 1}
	fmt.Println("Map Literal:", mapliteral)
}
