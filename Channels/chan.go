package main

import (
	"fmt"
	"time"
)

func worker(ch <-chan int) {
	// how to get channel size here
	fmt.Println("lenth of channel:", len(ch))
	for val := range ch {
		fmt.Println("Received:", val)
	}
}

func main() {
	bufferedChan := make(chan int, 3)
	bufferedChan2 := make(chan int, 3)

	go worker(bufferedChan)
	fmt.Println("Starting to send values to the channel...")
	time.Sleep(1500 * time.Millisecond)

	for i := 1; i <= 5; i++ {
		fmt.Println("Sending:", i)
		bufferedChan <- i
		time.Sleep(1500 * time.Millisecond)
	}

	go worker(bufferedChan2)
	bufferedChan2 <- 10
	close(bufferedChan)

	close(bufferedChan2)
	fmt.Println("All values sent and channels closed.")
}
