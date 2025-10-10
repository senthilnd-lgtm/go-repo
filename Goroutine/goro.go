package main

import (
	"fmt"
	"time"
)

func worker(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Worker %d: %d\n", id, i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go worker(1)
	go worker(2)

	// Wait for goroutines to finish
	time.Sleep(time.Second * 3)
	fmt.Println("Main finished")
}
