package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "C:\\Senthil\\GoLang\\WS\\Files\\example.txt"

	// Create a file
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write to file
	_, err = file.WriteString("Hello, Go file operations!\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Read from file
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File contents:", string(data))

	// Delete file
	file.Close()
	err = os.Remove(filename)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
	fmt.Println("File deleted successfully.")
}
