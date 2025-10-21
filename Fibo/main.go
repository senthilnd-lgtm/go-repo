package main

import (
	"fmt"
	"sync"
)

var a = 0
var b = 1

func fibo(x, y int, c chan int, wg *sync.WaitGroup) {
	res := x + y
	a = b
	b = res
	c <- res
	wg.Done()
}

func main() {
	var wg = sync.WaitGroup{}
	ch := make(chan int)
	for range 5 {
		wg.Add(1)
		go fibo(a, b, ch, &wg)
		fmt.Println(<-ch)
		fmt.Println(" ")
	}

	//fmt.Println(<-ch)
	wg.Wait()
	close(ch)
}
